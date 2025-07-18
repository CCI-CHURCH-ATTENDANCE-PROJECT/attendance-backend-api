package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"church-attendance-api/internal/config"
	"church-attendance-api/internal/database"
	"church-attendance-api/internal/handler"
	"church-attendance-api/internal/middleware"
	"church-attendance-api/internal/repository"
	"church-attendance-api/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// CustomValidator wraps the validator
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Load configuration
	cfg := config.Load()
	log.Printf("Starting server in %s mode on port %s", cfg.Env, cfg.Port)

	// Connect to database
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create indexes
	if err := db.CreateIndexes(); err != nil {
		log.Fatalf("Failed to create database indexes: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	attendanceRepo := repository.NewAttendanceRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	sermonRepo := repository.NewSermonRepository(db)
	announcementRepo := repository.NewAnnouncementRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	familyMemberRepo := repository.NewFamilyMemberRepository(db)
	localChurchRepo := repository.NewLocalChurchRepository(db)

	// Initialize services
	authService := service.NewAuthService(cfg, userRepo, refreshTokenRepo)
	userService := service.NewUserService(cfg, userRepo)
	attendanceService := service.NewAttendanceService(cfg, attendanceRepo, userRepo)
	qrService := service.NewQRService(cfg, userRepo)
	sermonService := service.NewSermonService(cfg, sermonRepo)
	announcementService := service.NewAnnouncementService(cfg, announcementRepo)
	roleService := service.NewRoleService(cfg, roleRepo)
	familyMemberService := service.NewFamilyMemberService(cfg, familyMemberRepo)
	localChurchService := service.NewLocalChurchService(cfg, localChurchRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)
	qrHandler := handler.NewQRHandler(qrService)
	roleHandler := handler.NewRoleHandler(roleService)
	sermonHandler := handler.NewSermonHandler(sermonService)
	announcementHandler := handler.NewAnnouncementHandler(announcementService)
	familyMemberHandler := handler.NewFamilyMemberHandler(familyMemberService)
	localChurchHandler := handler.NewLocalChurchHandler(localChurchService)

	// Initialize Echo
	e := echo.New()

	// Set custom validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.CORSMiddleware(cfg.CORSOrigins))
	e.Use(middleware.SecurityHeadersMiddleware())

	// Rate limiting
	e.Use(echomiddleware.RateLimiter(echomiddleware.NewRateLimiterMemoryStore(100)))

	// Request timeout
	e.Use(echomiddleware.TimeoutWithConfig(echomiddleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// API routes
	api := e.Group("/api/v1")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.POST("/register", authHandler.BasicRegister)
	auth.POST("/register/complete", authHandler.CompleteRegister)
	auth.POST("/login", authHandler.Login)
	auth.POST("/refresh", authHandler.RefreshToken)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.JWTMiddleware(cfg))

	// Auth protected routes
	protected.POST("/logout", authHandler.Logout)

	// User routes
	users := protected.Group("/users")
	users.GET("/search", userHandler.SearchUsers)
	users.GET("", userHandler.GetAllUsers)
	users.GET("/filter", userHandler.FilterUsers)

	// Attendance routes
	attendance := protected.Group("/attendance")
	attendance.POST("", attendanceHandler.CreateAttendance)
	attendance.POST("/qr-checkin", attendanceHandler.QRCheckin)
	attendance.GET("/history", attendanceHandler.GetAttendanceHistory)
	attendance.GET("/analytics", attendanceHandler.GetAttendanceAnalytics)

	// QR Code routes
	qr := protected.Group("/qr")
	qr.POST("/generate", qrHandler.GenerateQRCode)

	// Role routes (Admin only)
	roles := protected.Group("/roles")
	roles.Use(middleware.AdminMiddleware())
	roles.POST("", roleHandler.CreateRole)
	roles.GET("", roleHandler.GetRoles)
	roles.GET("/:id", roleHandler.GetRoleByID)
	roles.PUT("/:id", roleHandler.UpdateRole)
	roles.DELETE("/:id", roleHandler.DeleteRole)

	// Sermon routes
	sermons := protected.Group("/sermons")
	sermons.POST("", sermonHandler.CreateSermon)
	sermons.GET("", sermonHandler.GetSermons)
	sermons.GET("/:id", sermonHandler.GetSermonByID)
	sermons.PUT("/:id", sermonHandler.UpdateSermon)
	sermons.DELETE("/:id", sermonHandler.DeleteSermon)

	// Announcement routes
	announcements := protected.Group("/announcements")
	announcements.POST("", announcementHandler.CreateAnnouncement)
	announcements.GET("", announcementHandler.GetAnnouncements)
	announcements.GET("/active", announcementHandler.GetActiveAnnouncements)
	announcements.GET("/:id", announcementHandler.GetAnnouncementByID)
	announcements.PUT("/:id", announcementHandler.UpdateAnnouncement)
	announcements.DELETE("/:id", announcementHandler.DeleteAnnouncement)

	// Family member routes
	familyMembers := protected.Group("/family-members")
	familyMembers.POST("", familyMemberHandler.CreateFamilyMember)
	familyMembers.GET("", familyMemberHandler.GetFamilyMembers)
	familyMembers.GET("/:id", familyMemberHandler.GetFamilyMemberByID)
	familyMembers.PUT("/:id", familyMemberHandler.UpdateFamilyMember)
	familyMembers.DELETE("/:id", familyMemberHandler.DeleteFamilyMember)

	// Local church routes (Admin only)
	churches := protected.Group("/churches")
	churches.Use(middleware.AdminMiddleware())
	churches.POST("", localChurchHandler.CreateChurch)
	churches.GET("", localChurchHandler.GetChurches)
	churches.GET("/:id", localChurchHandler.GetChurchByID)
	churches.PUT("/:id", localChurchHandler.UpdateChurch)
	churches.DELETE("/:id", localChurchHandler.DeleteChurch)

	// Start server in a goroutine
	go func() {
		if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown server
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
