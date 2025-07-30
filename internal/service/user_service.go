package service

import (
	"context"
	"fmt"

	"church-attendance-api/internal/config"
	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/models"
	"church-attendance-api/internal/repository"
)

type UserService struct {
	cfg      *config.Config
	userRepo *repository.UserRepository
}

func NewUserService(cfg *config.Config, userRepo *repository.UserRepository) *UserService {
	return &UserService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (s *UserService) SearchUsers(ctx context.Context, query string, page, limit int) (*dto.PaginatedResponse, error) {
	users, total, err := s.userRepo.Search(ctx, query, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}

	// Convert to response format
	var userResponses []dto.UserSummary
	for _, user := range users {
		userResponses = append(userResponses, dto.UserSummary{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}

	pagination := dto.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: (total + limit - 1) / limit,
	}

	return &dto.PaginatedResponse{
		Data:       userResponses,
		Pagination: pagination,
	}, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, page, limit int) (*dto.PaginatedResponse, error) {
	users, total, err := s.userRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	// Convert to response format
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:                           user.ID,
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
			DateJoined:                   user.DateJoined,
			DateUpdated:                  user.DateUpdated,
			Role:                         user.Role,
			EmergencyContactName:         user.EmergencyContactName,
			EmergencyContactPhone:        user.EmergencyContactPhone,
			EmergencyContactEmail:        user.EmergencyContactEmail,
			EmergencyContactRelationship: user.EmergencyContactRelationship,
		})
	}

	pagination := dto.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: (total + limit - 1) / limit,
	}

	return &dto.PaginatedResponse{
		Data:       userResponses,
		Pagination: pagination,
	}, nil
}

func (s *UserService) FilterUsers(ctx context.Context, field, value string, page, limit int) (*dto.PaginatedResponse, error) {
	users, total, err := s.userRepo.Filter(ctx, field, value, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to filter users: %w", err)
	}

	// Convert to response format
	var userResponses []dto.UserSummary
	for _, user := range users {
		userResponses = append(userResponses, dto.UserSummary{
			UserID:    user.UserID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
	}

	pagination := dto.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: (total + limit - 1) / limit,
	}

	return &dto.PaginatedResponse{
		Data:       userResponses,
		Pagination: pagination,
	}, nil
}
