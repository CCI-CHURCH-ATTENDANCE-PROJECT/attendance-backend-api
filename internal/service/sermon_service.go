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

type SermonService struct {
	config     *config.Config
	sermonRepo repository.SermonRepository
}

func NewSermonService(cfg *config.Config, sermonRepo *repository.SermonRepository) *SermonService {
	return &SermonService{
		config:     cfg,
		sermonRepo: *sermonRepo,
	}
}

func (s *SermonService) CreateSermon(ctx context.Context, req *dto.CreateSermonRequest, userID primitive.ObjectID) (*dto.SermonResponse, error) {
	// Validate request
	if req.Title == "" {
		return nil, errors.New("sermon title is required")
	}

	if req.Speaker == "" {
		return nil, errors.New("speaker is required")
	}
	if req.Notes == "" {
		return nil, errors.New("sermon note is required")
	}

	// date, err := time.Parse("2006-01-02", req.Date)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid date format: *w", err)
	// }
	// Create sermon
	sermon := &models.Sermon{
		Preacher:      req.Speaker,
		DateOfMeeting: req.Date,
		SermonTopic:   req.Title,
		SermonNote:    req.Notes,
		EntryMadeBy:   userID,
		VideoUrl:      req.VideoURL,
		AudioUrl:      req.AudioURL,
		Scripture:     req.Scripture,
		Series:        req.Series,
		Tags:          req.Tags,
	}

	err := s.sermonRepo.Create(ctx, sermon)
	if err != nil {
		return nil, fmt.Errorf("failed to create sermon: %w", err)
	}

	return &dto.SermonResponse{
		ID:          sermon.ID.Hex(),
		Title:       sermon.SermonTopic,
		Speaker:     sermon.Preacher,
		Date:        sermon.DateOfMeeting,
		VideoURL:    sermon.VideoUrl,
		AudioURL:    sermon.AudioUrl,
		Notes:       sermon.SermonNote,
		Scripture:   sermon.Scripture,
		Series:      sermon.Series,
		DateAdded:   time.Now(),
		DateUpdated: time.Now(),
	}, nil
}

func (s *SermonService) GetSermons(ctx context.Context, page, limit int) (*dto.PaginatedSermonsResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	sermons, total, err := s.sermonRepo.GetAll(ctx, page, limit, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get sermons: %w", err)
	}

	sermonResponses := make([]*dto.SermonResponse, len(sermons))
	for i, sermon := range sermons {
		sermonResponses[i] = &dto.SermonResponse{
		ID:          sermon.ID.Hex(),
		Title:       sermon.SermonTopic,
		Speaker:     sermon.Preacher,
		Date:        sermon.DateOfMeeting,
		VideoURL:    sermon.VideoUrl,
		AudioURL:    sermon.AudioUrl,
		Notes:       sermon.SermonNote,
		Scripture:   sermon.Scripture,
		Series:      sermon.Series,
		DateAdded:   time.Now(),
		DateUpdated: time.Now(),
		}
	}

	return &dto.PaginatedSermonsResponse{
		Data:       sermonResponses,
		Pagination: utils.NewPagination(page, limit, total),
	}, nil
}

func (s *SermonService) GetSermonByID(ctx context.Context, id string) (*dto.SermonResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid sermon ID")
	}

	sermon, err := s.sermonRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("sermon with This ID does not exist: %w", err)
	}

	return &dto.SermonResponse{
		ID:          sermon.ID.Hex(),
		Title:       sermon.SermonTopic,
		Speaker:     sermon.Preacher,
		Date:        sermon.DateOfMeeting,
		VideoURL:    sermon.VideoUrl,
		AudioURL:    sermon.AudioUrl,
		Notes:       sermon.SermonNote,
		Scripture:   sermon.Scripture,
		Series:      sermon.Series,
		DateAdded:   time.Now(),
		DateUpdated: time.Now(),
	}, nil
}

func (s *SermonService) UpdateSermon(ctx context.Context, id string, req *dto.UpdateSermonRequest) (*dto.SermonResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid sermon ID")
	}

	sermon, err := s.sermonRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get sermon: %w", err)
	}

	// Update fields
	if req.Title != "" {
		sermon.SermonTopic = req.Title
	}
	if req.Speaker != "" {
		sermon.Preacher = req.Speaker
	}
	if req.Date != "" {
		sermon.DateOfMeeting = req.Date
	}
	if req.Notes != "" {
		sermon.SermonNote = req.Notes
	}

	err = s.sermonRepo.Update(ctx, sermon)
	if err != nil {
		return nil, fmt.Errorf("failed to update sermon: %w", err)
	}

	return &dto.SermonResponse{
		ID:          sermon.ID.Hex(),
		Title:       sermon.SermonTopic,
		Speaker:     sermon.Preacher,
		Date:        sermon.DateOfMeeting,
		VideoURL:    sermon.VideoUrl,
		AudioURL:    sermon.AudioUrl,
		Notes:       sermon.SermonNote,
		Scripture:   sermon.Scripture,
		Series:      sermon.Series,
		DateAdded:   time.Now(),
		DateUpdated: time.Now(),
	}, nil
}

func (s *SermonService) DeleteSermon(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid sermon ID")
	}

	if err := s.sermonRepo.Delete(ctx, objID); err != nil {
		return fmt.Errorf("failed to delete sermon: %w", err)
	}
	return nil
}
