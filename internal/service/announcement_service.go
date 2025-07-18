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

type AnnouncementService struct {
	config           *config.Config
	announcementRepo repository.AnnouncementRepository
}

func NewAnnouncementService(cfg *config.Config, announcementRepo *repository.AnnouncementRepository) *AnnouncementService {
	return &AnnouncementService{
		config:           cfg,
		announcementRepo: *announcementRepo,
	}
}

func (s *AnnouncementService) CreateAnnouncement(ctx context.Context, req *dto.CreateAnnouncementRequest, userID primitive.ObjectID) (*dto.AnnouncementResponse, error) {
	// Validate request
	if req.Title == "" {
		return nil, errors.New("announcement title is required")
	}

	if req.AnnouncementContent == "" {
		return nil, errors.New("announcement content is required")
	}

	announcement_due_date, _ := time.Parse("2006-01-02", req.AnnouncementDueDate)
	start_date, _ := time.Parse("2006-01-02", req.StartDate)
	end_date, _ := time.Parse("2006-01-02", req.EndDate)

	// Create announcement
	announcement := &models.Announcement{
		Title:                   req.Title,
		AnnouncementContent:     req.AnnouncementContent,
		AnnouncementDueDate:     announcement_due_date,
		StartDate:               start_date,
		EndDate:                 end_date,
		AnnouncementType:        req.AnnouncementType,
		Priority:                req.Priority,
		TargetUsers:             req.TargetUsers,
		ImageUrl:                req.ImageUrl,
		AnnouncementEntryMadeBy: userID,
		Status:                  "Pending",
		DateAdded:               time.Now(),
		DateUpdated:             time.Now(),
	}

	err := s.announcementRepo.Create(ctx, announcement)
	if err != nil {
		return nil, fmt.Errorf("failed to create announcement: %w", err)
	}

	return &dto.AnnouncementResponse{
		ID:                      announcement.ID.Hex(),
		Title:                   announcement.Title,
		AnnouncementContent:     announcement.AnnouncementContent,
		AnnouncementType:        announcement.AnnouncementType,
		AnnouncementDueDate:     announcement.AnnouncementDueDate,
		StartDate:               announcement.StartDate,
		EndDate:                 announcement.EndDate,
		Priority:                announcement.Priority,
		TargetUsers:             announcement.TargetUsers,
		ImageURL:                announcement.ImageUrl,
		Status:                  announcement.Status,
		AnnouncementEntryMadeBy: announcement.AnnouncementEntryMadeBy,
		DateAdded:               announcement.DateAdded,
		DateUpdated:             announcement.DateAdded,
	}, nil
}

func (s *AnnouncementService) GetAnnouncements(ctx context.Context, page, limit int) (*dto.PaginatedAnnouncementsResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	announcements, total, err := s.announcementRepo.GetAll(ctx, page, limit, "")
	if err != nil {
		return nil, fmt.Errorf("failed to get announcements: %w", err)
	}

	announcementResponses := make([]*dto.AnnouncementResponse, len(announcements))
	for i, announcement := range announcements {
		// AnnouncementDueDate := time.Now()
		// if announcement.AnnouncementDueDate != nil {
		// 	AnnouncementDueDate = *announcement.AnnouncementDueDate
		// }

		announcementResponses[i] = &dto.AnnouncementResponse{
			ID:                      announcement.ID.Hex(),
			Title:                   announcement.Title,
			AnnouncementContent:     announcement.AnnouncementContent,
			AnnouncementType:        announcement.AnnouncementType,
			StartDate:               announcement.StartDate,
			EndDate:                 announcement.EndDate,
			AnnouncementDueDate:     announcement.AnnouncementDueDate,
			Priority:                announcement.Priority,
			TargetUsers:             announcement.TargetUsers,
			ImageURL:                announcement.ImageUrl,
			Status:                  announcement.Status,
			AnnouncementEntryMadeBy: announcement.AnnouncementEntryMadeBy,
			DateAdded:               announcement.DateAdded,
			DateUpdated:             announcement.DateAdded,
		}
	}

	return &dto.PaginatedAnnouncementsResponse{
		Data:       announcementResponses,
		Pagination: utils.NewPagination(page, limit, total),
	}, nil
}

func (s *AnnouncementService) GetAnnouncementByID(ctx context.Context, id string) (*dto.AnnouncementResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid announcement ID")
	}

	announcement, err := s.announcementRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcement: %w", err)
	}

	return &dto.AnnouncementResponse{
		ID:                      announcement.ID.Hex(),
		Title:                   announcement.Title,
		AnnouncementContent:     announcement.AnnouncementContent,
		StartDate:               announcement.StartDate,
		EndDate:                 announcement.EndDate,
		AnnouncementType:        announcement.AnnouncementType,
		AnnouncementDueDate:     announcement.AnnouncementDueDate,
		Priority:                announcement.Priority,
		TargetUsers:             announcement.TargetUsers,
		ImageURL:                announcement.ImageUrl,
		Status:                  announcement.Status,
		AnnouncementEntryMadeBy: announcement.AnnouncementEntryMadeBy,
		DateAdded:               announcement.DateAdded,
		DateUpdated:             announcement.DateAdded,
	}, nil
}

func (s *AnnouncementService) UpdateAnnouncement(ctx context.Context, id string, req *dto.UpdateAnnouncementRequest) (*dto.AnnouncementResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid announcement ID")
	}

	announcement, err := s.announcementRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcement: %w", err)
	}
	announcement_due_date, _ := time.Parse("2006-01-02", req.AnnouncementDueDate)
	start_date, _ := time.Parse("2006-01-02", req.StartDate)
	end_date, _ := time.Parse("2006-01-02", req.EndDate)

	// Update fields
	if req.Title != "" {
		announcement.Title = req.Title
	}
	if req.AnnouncementContent != "" {
		announcement.AnnouncementContent = req.AnnouncementContent
	}
	if req.Title == "" {
		return nil, errors.New("announcement title is required")
	}
	if req.AnnouncementContent == "" {
		return nil, errors.New("announcement content is required")
	}
	if req.StartDate == "" {
		announcement.StartDate = start_date
		return nil, errors.New("no new announcement start date was submitted")
	}
	if req.EndDate == "" {
		announcement.EndDate = end_date
		return nil, errors.New("no new announcement end date was submitted")
	}
	if req.AnnouncementDueDate == "" {
		announcement.AnnouncementDueDate = announcement_due_date
		return nil, errors.New("no new announcement due date was submitted")
	}
	// if !req.StartDate.IsZero() {
	// 	announcement.StartDate = req.StartDate
	// }
	// if !req.EndDate.IsZero() {
	// 	announcement.EndDate = req.EndDate
	// }

	err = s.announcementRepo.Update(ctx, announcement)
	if err != nil {
		return nil, fmt.Errorf("failed to update announcement: %w", err)
	}

	return &dto.AnnouncementResponse{
		ID:                      announcement.ID.Hex(),
		Title:                   announcement.Title,
		AnnouncementContent:     announcement.AnnouncementContent,
		StartDate:               announcement.StartDate,
		EndDate:                 announcement.EndDate,
		AnnouncementType:        announcement.AnnouncementType,
		AnnouncementDueDate:     announcement.AnnouncementDueDate,
		Priority:                announcement.Priority,
		TargetUsers:             announcement.TargetUsers,
		ImageURL:                announcement.ImageUrl,
		Status:                  announcement.Status,
		AnnouncementEntryMadeBy: announcement.AnnouncementEntryMadeBy,
		DateAdded:               announcement.DateAdded,
		DateUpdated:             announcement.DateAdded,
	}, nil
}

func (s *AnnouncementService) DeleteAnnouncement(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid announcement ID")
	}

	if err := s.announcementRepo.Delete(ctx, objID); err != nil {
		return fmt.Errorf("failed to delete announcement: %w", err)
	}
	return nil
}

func (s *AnnouncementService) GetActiveAnnouncements(ctx context.Context, page, limit int) (*dto.PaginatedAnnouncementsResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	announcements, total, err := s.announcementRepo.GetAll(ctx, page, limit, "Pending")
	if err != nil {
		return nil, fmt.Errorf("failed to get active announcements: %w", err)
	}

	announcementResponses := make([]*dto.AnnouncementResponse, len(announcements))
	for i, announcement := range announcements {

		announcementResponses[i] = &dto.AnnouncementResponse{
			ID:                      announcement.ID.Hex(),
			Title:                   announcement.Title,
			AnnouncementContent:     announcement.AnnouncementContent,
			StartDate:               announcement.StartDate,
			EndDate:                 announcement.EndDate,
			AnnouncementType:        announcement.AnnouncementType,
			AnnouncementDueDate:     announcement.AnnouncementDueDate,
			Priority:                announcement.Priority,
			TargetUsers:             announcement.TargetUsers,
			ImageURL:                announcement.ImageUrl,
			Status:                  announcement.Status,
			AnnouncementEntryMadeBy: announcement.AnnouncementEntryMadeBy,
			DateAdded:               announcement.DateAdded,
			DateUpdated:             announcement.DateAdded,
		}
	}

	return &dto.PaginatedAnnouncementsResponse{
		Data:       announcementResponses,
		Pagination: utils.NewPagination(page, limit, total),
	}, nil
}
