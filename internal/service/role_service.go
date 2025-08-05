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

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleService struct {
	config   *config.Config
	roleRepo repository.RoleRepository
}

func NewRoleService(cfg *config.Config, roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		config:   cfg,
		roleRepo: *roleRepo,
	}
}

func (s *RoleService) CreateRole(ctx context.Context, req *dto.CreateRoleRequest) (*dto.RoleResponse, error) {
	// Validate request
	if req.RoleName == "" {
		return nil, errors.New("role name is required")
	}

	// Check if role already exists
	if existing, _ := s.roleRepo.GetByName(ctx, req.RoleName); existing != nil {
		return nil, errors.New("role with this name already exists")
	}

	// Create role
	role := &models.Role{
		RoleName:        req.RoleName,
		RoleDescription: req.RoleDescription,
		Permissions:     req.Permissions,
		DateAdded:       time.Now(),
		DateUpdated:     time.Now(),
	}

	err := s.roleRepo.Create(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	return &dto.RoleResponse{
		ID:              role.ID.Hex(),
		RoleName:        role.RoleName,
		RoleDescription: role.RoleDescription,
		Permissions:     role.Permissions,
		TotalMembers:    role.TotalMembers,
		DateAdded:       role.DateAdded,
		DateUpdated:     role.DateUpdated,
	}, nil
}

func (s *RoleService) GetRoles(ctx context.Context, page, limit int) (*dto.PaginatedRolesResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	roles, total, err := s.roleRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}

	roleResponses := make([]*dto.RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = &dto.RoleResponse{
			ID:              role.ID.Hex(),
			RoleName:        role.RoleName,
			RoleDescription: role.RoleDescription,
			Permissions:     role.Permissions,
			TotalMembers:    role.TotalMembers,
			DateAdded:       role.DateAdded,
			DateUpdated:     role.DateUpdated,
		}
	}

	return &dto.PaginatedRolesResponse{
		Data:       roleResponses,
		Pagination: utils.NewPagination(page, limit, total),
	}, nil
}

func (s *RoleService) GetRoleByID(ctx context.Context, id string) (*dto.RoleResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid role ID")
	}

	role, err := s.roleRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %w", err)
	}

	return &dto.RoleResponse{
		ID:              role.ID.Hex(),
		RoleName:        role.RoleName,
		RoleDescription: role.RoleDescription,
		Permissions:     role.Permissions,
		TotalMembers:    role.TotalMembers,
		DateAdded:       role.DateAdded,
		DateUpdated:     role.DateUpdated,
	}, nil
}

func (s *RoleService) UpdateRole(ctx context.Context, id string, req *dto.UpdateRoleRequest) (*dto.RoleResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid role ID")
	}

	role, err := s.roleRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %w", err)
	}

	// Update fields
	if req.RoleName != "" {
		// Check if new name already exists (but not for the same role)
		if existing, _ := s.roleRepo.GetByName(ctx, req.RoleName); existing != nil && existing.ID.Hex() != id {
			return nil, errors.New("role with this name already exists")
		}
		role.RoleName = req.RoleName
	}

	if req.RoleDescription != "" {
		role.RoleDescription = req.RoleDescription
	}

	// if req.Permissions != "" {
	// 	role.Permissions = req.Permissions
	// }

	role.DateUpdated = time.Now()

	err = s.roleRepo.Update(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	return &dto.RoleResponse{
		ID:              role.ID.Hex(),
		RoleName:        role.RoleName,
		RoleDescription: role.RoleDescription,
		Permissions:     role.Permissions,
		TotalMembers:    role.TotalMembers,
		DateAdded:       role.DateAdded,
		DateUpdated:     role.DateUpdated,
	}, nil
}

func (s *RoleService) DeleteRole(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid role ID")
	}

	if err := s.roleRepo.Delete(ctx, objID); err != nil {
		return fmt.Errorf("failed to delete role: %w", err)
	}
	return nil
}
