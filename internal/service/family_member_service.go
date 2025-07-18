package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"church-attendance-api/internal/config"
	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/models"
	"church-attendance-api/internal/repository"
	"church-attendance-api/internal/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyMemberService struct {
	config           *config.Config
	familyMemberRepo repository.FamilyMemberRepository
}

func NewFamilyMemberService(cfg *config.Config, familyMemberRepo *repository.FamilyMemberRepository) *FamilyMemberService {
	return &FamilyMemberService{
		config:           cfg,
		familyMemberRepo: *familyMemberRepo,
	}
}

func (s *FamilyMemberService) CreateFamilyMember(ctx context.Context, req *dto.CreateFamilyMemberRequest) (*dto.FamilyMemberResponse, error) {
	// Validate request
	if req.FamilyMemberName == "" {
		return nil, errors.New("family member name is required")
	}

	// Create family member
	familyMember := &models.FamilyMember{
		FamilyMemberName:         req.FamilyMemberName,
		FamilyMemberPhone:        req.FamilyMemberPhone,
		FamilyMemberEmail:        req.FamilyMemberEmail,
		FamilyMemberRelationship: req.FamilyMemberRelationship,
		FamilyMemberDateOfBirth:  req.FamilyMemberDateOfBirth,
		FamilyMemberGender:       req.FamilyMemberGender,
		FamilyMemberOccupation:   req.FamilyMemberOccupation,
		DateAdded:                time.Now(),
	}

	err := s.familyMemberRepo.Create(ctx, familyMember)
	if err != nil {
		return nil, fmt.Errorf("failed to create family member: %w", err)
	}

	return &dto.FamilyMemberResponse{
		ID:                       familyMember.ID.Hex(),
		FamilyMemberName:         familyMember.FamilyMemberName,
		FamilyMemberPhone:        familyMember.FamilyMemberPhone,
		FamilyMemberEmail:        familyMember.FamilyMemberEmail,
		FamilyMemberRelationship: familyMember.FamilyMemberRelationship,
		FamilyMemberDateOfBirth:  familyMember.FamilyMemberDateOfBirth,
		FamilyMemberGender:       familyMember.FamilyMemberGender,
		FamilyMemberOccupation:   familyMember.FamilyMemberOccupation,
		DateAdded:                familyMember.DateAdded,
	}, nil
}

func (s *FamilyMemberService) GetFamilyMembers(ctx context.Context, page, limit int) (*dto.PaginatedFamilyMembersResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	familyMembers, total, err := s.familyMemberRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get family members: %w", err)
	}

	familyMemberResponses := make([]*dto.FamilyMemberResponse, len(familyMembers))
	for i, familyMember := range familyMembers {
		familyMemberResponses[i] = &dto.FamilyMemberResponse{
			ID:                       familyMember.ID.Hex(),
			FamilyMemberName:         familyMember.FamilyMemberName,
			FamilyMemberPhone:        familyMember.FamilyMemberPhone,
			FamilyMemberEmail:        familyMember.FamilyMemberEmail,
			FamilyMemberRelationship: familyMember.FamilyMemberRelationship,
			FamilyMemberDateOfBirth:  familyMember.FamilyMemberDateOfBirth,
			FamilyMemberGender:       familyMember.FamilyMemberGender,
			FamilyMemberOccupation:   familyMember.FamilyMemberOccupation,
			DateAdded:                familyMember.DateAdded,
		}
	}

	return &dto.PaginatedFamilyMembersResponse{
		Data:       familyMemberResponses,
		Pagination: utils.NewPagination(page, limit, total),
	}, nil
}

func (s *FamilyMemberService) GetFamilyMemberByID(ctx context.Context, id string) (*dto.FamilyMemberResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid family member ID")
	}

	familyMember, err := s.familyMemberRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get family member: %w", err)
	}

	return &dto.FamilyMemberResponse{
		ID:                       familyMember.ID.Hex(),
		FamilyMemberName:         familyMember.FamilyMemberName,
		FamilyMemberPhone:        familyMember.FamilyMemberPhone,
		FamilyMemberEmail:        familyMember.FamilyMemberEmail,
		FamilyMemberRelationship: familyMember.FamilyMemberRelationship,
		FamilyMemberDateOfBirth:  familyMember.FamilyMemberDateOfBirth,
		FamilyMemberGender:       familyMember.FamilyMemberGender,
		FamilyMemberOccupation:   familyMember.FamilyMemberOccupation,
		DateAdded:                familyMember.DateAdded,
	}, nil
}

func (s *FamilyMemberService) UpdateFamilyMember(ctx context.Context, id string, req *dto.UpdateFamilyMemberRequest) (*dto.FamilyMemberResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid family member ID")
	}

	familyMember, err := s.familyMemberRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get family member: %w", err)
	}

	// Update fields
	if req.FamilyMemberName != "" {
		familyMember.FamilyMemberName = req.FamilyMemberName
	}
	if req.FamilyMemberPhone != "" {
		familyMember.FamilyMemberPhone = req.FamilyMemberPhone
	}
	if req.FamilyMemberEmail != "" {
		familyMember.FamilyMemberEmail = req.FamilyMemberEmail
	}
	if req.FamilyMemberRelationship != "" {
		familyMember.FamilyMemberRelationship = req.FamilyMemberRelationship
	}
	if req.FamilyMemberDateOfBirth != nil {
		familyMember.FamilyMemberDateOfBirth = req.FamilyMemberDateOfBirth
	}
	if req.FamilyMemberGender != "" {
		familyMember.FamilyMemberGender = req.FamilyMemberGender
	}
	if req.FamilyMemberOccupation != "" {
		familyMember.FamilyMemberOccupation = req.FamilyMemberOccupation
	}

	err = s.familyMemberRepo.Update(ctx, familyMember)
	if err != nil {
		return nil, fmt.Errorf("failed to update family member: %w", err)
	}

	return &dto.FamilyMemberResponse{
		ID:                       familyMember.ID.Hex(),
		FamilyMemberName:         familyMember.FamilyMemberName,
		FamilyMemberPhone:        familyMember.FamilyMemberPhone,
		FamilyMemberEmail:        familyMember.FamilyMemberEmail,
		FamilyMemberRelationship: familyMember.FamilyMemberRelationship,
		FamilyMemberDateOfBirth:  familyMember.FamilyMemberDateOfBirth,
		FamilyMemberGender:       familyMember.FamilyMemberGender,
		FamilyMemberOccupation:   familyMember.FamilyMemberOccupation,
		DateAdded:                familyMember.DateAdded,
	}, nil
}

func (s *FamilyMemberService) DeleteFamilyMember(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid family member ID")
	}

	if err := s.familyMemberRepo.Delete(ctx, objID); err != nil {
		return fmt.Errorf("failed to delete family member: %w", err)
	}
	return nil
}
