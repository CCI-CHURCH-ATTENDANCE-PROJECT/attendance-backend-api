package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cci-api/internal/config"
	"cci-api/internal/dto"
	"cci-api/internal/models"
	"cci-api/internal/repository"
	"cci-api/internal/utils"
)

type AuthService struct {
	cfg              *config.Config
	userRepo         *repository.UserRepository
	refreshTokenRepo *repository.RefreshTokenRepository
	emailService     EmailService
}

func NewAuthService(cfg *config.Config, userRepo *repository.UserRepository, refreshTokenRepo *repository.RefreshTokenRepository, emailService EmailService) *AuthService {
	return &AuthService{
		cfg:              cfg,
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		emailService:     emailService,
	}
}

func (s *AuthService) BasicRegister(ctx context.Context, req *dto.BasicRegisterRequest) (*dto.BasicRegisterResponse, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("user already exists with this email")
	}

	// Check if passwords match
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	// Validate password strength
	if !utils.IsValidPassword(req.Password) {
		return nil, errors.New("password must be at least 8 characters long and contain uppercase, lowercase, and number")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Generate user ID
	userID, err := utils.GenerateUserID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate user ID: %w", err)
	}

	// Create user
	user := &models.User{
		UserID:      userID,
		Email:       req.Email,
		Password:    hashedPassword,
		Member:      true,
		Visitor:     false,
		DateJoined:  time.Now(),
		DateUpdated: time.Now(),
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &dto.BasicRegisterResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		CreatedAt: user.DateJoined,
	}, nil
}

func (s *AuthService) CompleteRegister(ctx context.Context, req *dto.CompleteRegisterRequest) (*dto.CompleteRegisterResponse, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("user already exists with this email")
	}

	// // Validate password strength
	// if !utils.IsValidPassword(req.Password) {
	// 	return nil, errors.New("password must be at least 8 characters long and contain uppercase, lowercase, special character and at  least a number")
	// }

	// Hash password
	// hashedPassword, err := utils.HashPassword(req.Password)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to hash password: %w", err)
	// }

	// Generate user ID
	userID, err := utils.GenerateUserID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate user ID: %w", err)
	}

	// Parse DateOf Birth sent as string to DateTime Format
	dateOfBirth, _ := time.Parse("2025-4-11", req.DateOfBirth)

	// Parse date Joined church from string to date time format
	dateJoinedChurch, _ := time.Parse("2025-4-11", req.DateJoinedChurch)

	//Generate QRCode image and QRCode token for the user and save it in the QRCodeTonken fields

	//A logic to decide if the user will have admin set to true or false
	// Create user with complete information
	user := &models.User{
		UserID:                       userID,
		Email:                        req.Email,
		FirstName:                    req.FirstName,
		LastName:                     req.LastName,
		Bio:                          req.Bio,
		DateOfBirth:                  dateOfBirth,
		Gender:                       req.Gender,
		Member:                       req.Member,
		Visitor:                      req.Visitor,
		Usher:                        req.Usher,
		Admin:                        false,
		UserWorkDepartment:           req.UserWorkDepartment,
		DateJoinedChurch:             dateJoinedChurch,
		FamilyHead:                   req.FamilyHead,
		UserCampus:                   req.UserCampus,
		PhoneNumber:                  req.PhoneNumber,
		InstagramHandle:              req.InstagramHandle,
		FamilyMembers:                req.FamilyMembers,
		Profession:                   req.Profession,
		UserHouseAddress:             req.UserHouseAddress,
		CampusState:                  req.CampusState,
		CampusCountry:                req.CampusCountry,
		EmergencyContactName:         req.EmergencyContactName,
		EmergencyContactPhone:        req.EmergencyContactPhone,
		EmergencyContactEmail:        req.EmergencyContactEmail,
		EmergencyContactRelationship: req.EmergencyContactRelationship,
		DateJoined:                   time.Now(),
		DateUpdated:                  time.Now(),
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate password reset token
	token, err := utils.GeneratePasswordRandomToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate password reset token: %w", err)
	}

	// Store token in user model
	user.PasswordResetToken = token
	user.PasswordResetExpires = time.Now().Add(time.Hour * 24) // Token expires in 24 hours
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user with password reset token: %w", err)
	}

	// Send signup email
	go func() {
		data := map[string]interface{}{
			"FirstName": user.FirstName,
			"Link":      fmt.Sprintf("%s/set-password?token=%s", s.cfg.FrontendURL, token),
		}
		if err := s.emailService.SendEmail(user.Email, "Welcome to CCI Member Portal, Set Your Password", "signup.html", data); err != nil {
			// Log the error, but don't block the registration process
			fmt.Printf("failed to send signup email: %v\n", err)
		}
	}()

	return &dto.CompleteRegisterResponse{
		UserID:                       user.UserID,
		FirstName:                    user.FirstName,
		LastName:                     user.LastName,
		Email:                        user.Email,
		Bio:                          user.Bio,
		DateOfBirth:                  user.DateOfBirth,
		Gender:                       user.Gender,
		Member:                       user.Member,
		Visitor:                      user.Visitor,
		Usher:                        user.Usher,
		UserWorkDepartment:           user.UserWorkDepartment,
		DateJoinedChurch:             user.DateJoinedChurch,
		FamilyHead:                   user.FamilyHead,
		UserCampus:                   user.UserCampus,
		CampusState:                  user.CampusState,
		CampusCountry:                user.CampusCountry,
		Profession:                   user.Profession,
		UserHouseAddress:             user.UserHouseAddress,
		PhoneNumber:                  user.PhoneNumber,
		InstagramHandle:              user.InstagramHandle,
		FamilyMembers:                user.FamilyMembers,
		CreatedAt:                    user.DateJoined,
		UpdatedAt:                    user.DateUpdated,
		Role:                         user.Role,
		EmergencyContactName:         user.EmergencyContactName,
		EmergencyContactPhone:        user.EmergencyContactPhone,
		EmergencyContactEmail:        user.EmergencyContactEmail,
		EmergencyContactRelationship: user.EmergencyContactRelationship,
	}, nil
}

func (s *AuthService) SetPassword(ctx context.Context, req *dto.SetPasswordRequest) error {
	// Get user by password reset token
	user, err := s.userRepo.GetByPasswordResetToken(ctx, req.Token)
	if err != nil {
		return fmt.Errorf("failed to get user by password reset token: %w", err)
	}
	if user == nil {
		return errors.New("invalid or expired token")
	}

	// Check if token has expired
	if time.Now().After(user.PasswordResetExpires) {
		return errors.New("token has expired")
	}

	// Validate password strength
	if !utils.IsValidPassword(req.Password) {
		return errors.New("password must be at least 8 characters long and contain uppercase, lowercase, and number")
	}

	// Check if passwords match
	if req.Password != req.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Update user's password
	user.Password = hashedPassword
	user.PasswordResetToken = ""
	user.PasswordResetExpires = time.Time{}
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT tokens
	accessToken, err := utils.GenerateJWT(user.UserID, user.Email, user.Admin, s.cfg.JWTSecret, s.cfg.JWTAccessExpiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := utils.GenerateRandomToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Store refresh token
	refreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.cfg.JWTRefreshExpiry),
	}

	err = s.refreshTokenRepo.Create(ctx, refreshTokenModel)
	if err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.UserSummary{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	// Get refresh token
	refreshToken, err := s.refreshTokenRepo.GetByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get refresh token: %w", err)
	}
	if refreshToken == nil {
		return nil, errors.New("invalid refresh token")
	}

	// Check if token is expired
	if time.Now().After(refreshToken.ExpiresAt) {
		// Delete expired token
		s.refreshTokenRepo.DeleteByToken(ctx, req.RefreshToken)
		return nil, errors.New("refresh token has expired")
	}

	// Get user
	user, err := s.userRepo.GetByID(ctx, refreshToken.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Generate new access token
	accessToken, err := utils.GenerateJWT(user.UserID, user.Email, user.Admin, s.cfg.JWTSecret, s.cfg.JWTAccessExpiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate new refresh token
	newRefreshToken, err := utils.GenerateRandomToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new refresh token: %w", err)
	}

	// Delete old refresh token
	err = s.refreshTokenRepo.DeleteByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to delete old refresh token: %w", err)
	}

	// Store new refresh token
	newRefreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		Token:     newRefreshToken,
		ExpiresAt: time.Now().Add(s.cfg.JWTRefreshExpiry),
	}

	err = s.refreshTokenRepo.Create(ctx, newRefreshTokenModel)
	if err != nil {
		return nil, fmt.Errorf("failed to store new refresh token: %w", err)
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, userID string) error {
	// Get user
	user, err := s.userRepo.GetByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Delete all refresh tokens for user
	err = s.refreshTokenRepo.DeleteByUserID(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to delete refresh tokens: %w", err)
	}

	return nil
}
