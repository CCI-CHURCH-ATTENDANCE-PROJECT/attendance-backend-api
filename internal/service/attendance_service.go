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
)

type AttendanceService struct {
	cfg            *config.Config
	attendanceRepo *repository.AttendanceRepository
	userRepo       *repository.UserRepository
}

func NewAttendanceService(cfg *config.Config, attendanceRepo *repository.AttendanceRepository, userRepo *repository.UserRepository) *AttendanceService {
	return &AttendanceService{
		cfg:            cfg,
		attendanceRepo: attendanceRepo,
		userRepo:       userRepo,
	}
}

func (s *AttendanceService) CreateAttendance(ctx context.Context, req *dto.CreateAttendanceRequest) (*dto.AttendanceResponse, error) {
	// Get user
	user, err := s.userRepo.GetByUserID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check if user already has attendance for today
	today := time.Now()
	existingAttendance, err := s.attendanceRepo.GetByUserAndDate(ctx, user.ID, today)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing attendance: %w", err)
	}
	if existingAttendance != nil {
		return nil, errors.New("attendance already recorded for today, relax there is nothing to worry about")
	}

	// Determine if late (assuming service starts at 9 AM)
	isLate := today.Hour() > 9

	// Create attendance record
	attendance := &models.Attendance{
		User:                 user.ID,
		DateTimeOfAttendance: today,
		QRCodeBasedCheckin:   false,
		Late:                 isLate,
		ManualCheckin:        true,
		Visitor:              user.Visitor,
		Member:               user.Member,
	}

	err = s.attendanceRepo.Create(ctx, attendance)
	if err != nil {
		return nil, fmt.Errorf("failed to create attendance: %w", err)
	}

	return &dto.AttendanceResponse{
		ID:                   attendance.ID.Hex(),
		UserID:               user.UserID,
		DateTimeOfAttendance: attendance.DateTimeOfAttendance,
		QRCodeBasedCheckin:   attendance.QRCodeBasedCheckin,
		Late:                 attendance.Late,
		ManualCheckin:        attendance.ManualCheckin,
	}, nil
}

func (s *AttendanceService) QRCheckin(ctx context.Context, req *dto.QRCheckinRequest) (*dto.AttendanceResponse, error) {
	// Get user by QR token
	user, err := s.userRepo.GetByQRToken(ctx, req.QRCodeToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by QR token: %w", err)
	}
	if user == nil {
		return nil, errors.New("invalid QR code token")
	}

	// Check if user already has attendance for today
	today := time.Now()
	existingAttendance, err := s.attendanceRepo.GetByUserAndDate(ctx, user.ID, today)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing attendance: %w", err)
	}
	if existingAttendance != nil {
		return nil, errors.New("attendance already recorded for today")
	}

	// Determine if late (assuming service starts at 9 AM)
	isLate := today.Hour() > 9

	// Create attendance record
	attendance := &models.Attendance{
		User:                 user.ID,
		DateTimeOfAttendance: today,
		QRCodeBasedCheckin:   true,
		Late:                 isLate,
		ManualCheckin:        false,
	}

	err = s.attendanceRepo.Create(ctx, attendance)
	if err != nil {
		return nil, fmt.Errorf("failed to create attendance: %w", err)
	}

	return &dto.AttendanceResponse{
		ID:                   attendance.ID.Hex(),
		UserID:               user.UserID,
		DateTimeOfAttendance: attendance.DateTimeOfAttendance,
		QRCodeBasedCheckin:   attendance.QRCodeBasedCheckin,
		Late:                 attendance.Late,
		ManualCheckin:        attendance.ManualCheckin,
		Visitor:              attendance.Visitor,
		Member:               attendance.Member,
	}, nil
}

func (s *AttendanceService) GetAttendanceHistory(ctx context.Context, startDate, endDate *time.Time, page, limit int) (*dto.PaginatedResponse, error) {
	// Set default date range if not provided
	if startDate == nil {
		start := time.Now().AddDate(0, -1, 0) // Last month
		startDate = &start
	}
	if endDate == nil {
		end := time.Now()
		endDate = &end
	}

	// Get attendance data grouped by date
	attendanceData, err := s.attendanceRepo.GetAttendanceByDateRange(ctx, *startDate, *endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance history: %w", err)
	}

	// Transform data to response format
	var history []dto.AttendanceHistoryItem
	for _, data := range attendanceData {
		dateMap := data["_id"].(map[string]interface{})
		year := int(dateMap["year"].(int32))
		month := int(dateMap["month"].(int32))
		day := int(dateMap["day"].(int32))

		date := fmt.Sprintf("%d-%02d-%02d", year, month, day)

		history = append(history, dto.AttendanceHistoryItem{
			Date:            date,
			TotalAttendance: int(data["total_attendance"].(int32)),
			Members:         int(data["members"].(int32)),
			Visitors:        int(data["visitors"].(int32)),
		})
	}

	// Calculate pagination
	total := len(history)
	start := (page - 1) * limit
	end := start + limit

	if start >= total {
		history = []dto.AttendanceHistoryItem{}
	} else {
		if end > total {
			end = total
		}
		history = history[start:end]
	}

	pagination := dto.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: (total + limit - 1) / limit,
	}

	return &dto.PaginatedResponse{
		Data:       history,
		Pagination: pagination,
	}, nil
}

func (s *AttendanceService) GetAttendanceAnalytics(ctx context.Context, date time.Time) (*dto.AttendanceAnalytics, error) {
	// Get total active users all time
	totalUsers, err := s.userRepo.CountTotal(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count total users: %w", err)
	}

	// Get total attendance for the specific date
	totalAttendanceForDate, err := s.attendanceRepo.CountTotalForDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("failed to count attendance for date: %w", err)
	}

	// Get members count for the month
	membersForMonth, err := s.attendanceRepo.CountMembersForDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("failed to count members for month: %w", err)
	}

	// Get visitors count
	visitorsCount, err := s.attendanceRepo.CountVisitorsForDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("failed to count visitors: %w", err)
	}

	return &dto.AttendanceAnalytics{
		TotalActiveUsersAllTime: totalUsers,
		TotalAttendanceForDate:  totalAttendanceForDate,
		MembersForMonth:         membersForMonth,
		VisitorsCount:           visitorsCount,
	}, nil
}
