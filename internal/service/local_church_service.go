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

type LocalChurchService struct {
	config          *config.Config
	localChurchRepo *repository.LocalChurchRepository
}

func NewLocalChurchService(config *config.Config, localChurchRepo *repository.LocalChurchRepository) *LocalChurchService {
	return &LocalChurchService{
		config:          config,
		localChurchRepo: localChurchRepo,
	}
}

func (s *LocalChurchService) CreateChurch(ctx context.Context, req *dto.CreateLocalChurchRequest) (*dto.LocalChurchResponse, error) {
	// Validate request
	if req.ChurchName == "" {
		return nil, errors.New("church name is required")
	}

	// Create church
	church := &models.LocalChurch{
		ChurchName:         req.ChurchName,
		ChurchPhone:        req.ChurchPhone,
		ChurchEmail:        req.ChurchEmail,
		ChurchAddress:      req.ChurchAddress,
		StateCounty:        req.StateCounty,
		Country:            req.Country,
		SundayMeetingTime:  req.SundayMeetingTime,
		MidweekMeetingDay:  req.MidweekMeetingDay,
		MidweekMeetingTime: req.MidweekMeetingTime,
		Website:            req.Website,
		SocialMedia:        req.SocialMedia,
		PastorName:         req.PastorName,
		PastorPhone:        req.PastorPhone,
		PastorEmail:        req.PastorEmail,
		FoundedYear:        req.FoundedYear,
		Description:        req.Description,
		DateAdded:          time.Now(),
		DateUpdated:        time.Now(),
	}

	err := s.localChurchRepo.Create(ctx, church)
	if err != nil {
		return nil, fmt.Errorf("failed to create church: %w", err)
	}

	return &dto.LocalChurchResponse{
		ID:                 church.ID.Hex(),
		ChurchName:         church.ChurchName,
		ChurchPhone:        church.ChurchPhone,
		ChurchEmail:        church.ChurchEmail,
		ChurchAddress:      church.ChurchAddress,
		StateCounty:        church.StateCounty,
		Country:            church.Country,
		SundayMeetingTime:  church.SundayMeetingTime,
		MidweekMeetingDay:  church.MidweekMeetingDay,
		MidweekMeetingTime: church.MidweekMeetingTime,
		Website:            church.Website,
		SocialMedia:        church.SocialMedia,
		PastorName:         church.ChurchName,
		PastorPhone:        church.PastorPhone,
		PastorEmail:        church.PastorEmail,
		FoundedYear:        church.FoundedYear,
		Description:        church.Description,
		DateAdded:          church.DateAdded,
		DateUpdated:        church.DateUpdated,
	}, nil
}

func (s *LocalChurchService) GetChurches(ctx context.Context, page, limit int) (*dto.PaginatedLocalChurchesResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	churches, total, err := s.localChurchRepo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get churches: %w", err)
	}

	churchResponses := make([]*dto.LocalChurchResponse, len(churches))
	for i, church := range churches {
		churchResponses[i] = &dto.LocalChurchResponse{
			ID:                 church.ID.Hex(),
			ChurchName:         church.ChurchName,
			ChurchPhone:        church.ChurchPhone,
			ChurchEmail:        church.ChurchEmail,
			ChurchAddress:      church.ChurchAddress,
			StateCounty:        church.StateCounty,
			Country:            church.Country,
			SundayMeetingTime:  church.SundayMeetingTime,
			MidweekMeetingDay:  church.MidweekMeetingDay,
			MidweekMeetingTime: church.MidweekMeetingTime,
			Website:            church.Website,
			SocialMedia:        church.SocialMedia,
			PastorName:         church.ChurchName,
			PastorPhone:        church.PastorPhone,
			PastorEmail:        church.PastorEmail,
			FoundedYear:        church.FoundedYear,
			Description:        church.Description,
			DateAdded:          church.DateAdded,
			DateUpdated:        church.DateUpdated,
		}
	}

	return &dto.PaginatedLocalChurchesResponse{
		Data:       churchResponses,
		Pagination: utils.NewPagination(page, limit, total),
	}, nil
}

func (s *LocalChurchService) GetChurchByID(ctx context.Context, id string) (*dto.LocalChurchResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid church ID")
	}

	church, err := s.localChurchRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get church: %w", err)
	}

	return &dto.LocalChurchResponse{
		ID:                 church.ID.Hex(),
		ChurchName:         church.ChurchName,
		ChurchPhone:        church.ChurchPhone,
		ChurchEmail:        church.ChurchEmail,
		ChurchAddress:      church.ChurchAddress,
		StateCounty:        church.StateCounty,
		Country:            church.Country,
		SundayMeetingTime:  church.SundayMeetingTime,
		MidweekMeetingDay:  church.MidweekMeetingDay,
		MidweekMeetingTime: church.MidweekMeetingTime,
		Website:            church.Website,
		SocialMedia:        church.SocialMedia,
		PastorName:         church.ChurchName,
		PastorPhone:        church.PastorPhone,
		PastorEmail:        church.PastorEmail,
		FoundedYear:        church.FoundedYear,
		Description:        church.Description,
		DateAdded:          church.DateAdded,
		DateUpdated:        church.DateUpdated,
	}, nil
}

func (s *LocalChurchService) UpdateChurch(ctx context.Context, id string, req *dto.UpdateLocalChurchRequest) (*dto.LocalChurchResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid church ID")
	}

	church, err := s.localChurchRepo.GetByID(ctx, objID)
	if err != nil {
		return nil, fmt.Errorf("failed to get church: %w", err)
	}

	// Update fields
	if req.ChurchName != "" {
		church.ChurchName = req.ChurchName
	}
	if req.ChurchPhone != "" {
		church.ChurchPhone = req.ChurchPhone
	}
	if req.ChurchEmail != "" {
		church.ChurchEmail = req.ChurchEmail
	}
	if req.ChurchAddress != "" {
		church.ChurchAddress = req.ChurchAddress
	}
	if req.StateCounty != "" {
		church.StateCounty = req.StateCounty
	}
	if req.Country != "" {
		church.Country = req.Country
	}
	if req.SundayMeetingTime != 0 {
		church.SundayMeetingTime = req.SundayMeetingTime
	}
	if req.MidweekMeetingDay != "" {
		church.MidweekMeetingDay = req.MidweekMeetingDay
	}
	if req.MidweekMeetingTime != 0 {
		church.MidweekMeetingTime = req.MidweekMeetingTime
	}

	dateUpdated := time.Now()

	err = s.localChurchRepo.Update(ctx, church)
	if err != nil {
		return nil, fmt.Errorf("failed to update church: %w", err)
	}

	return &dto.LocalChurchResponse{
		ID:                 church.ID.Hex(),
		ChurchName:         church.ChurchName,
		ChurchPhone:        church.ChurchPhone,
		ChurchEmail:        church.ChurchEmail,
		ChurchAddress:      church.ChurchAddress,
		StateCounty:        church.StateCounty,
		Country:            church.Country,
		SundayMeetingTime:  church.SundayMeetingTime,
		MidweekMeetingDay:  church.MidweekMeetingDay,
		MidweekMeetingTime: church.MidweekMeetingTime,
		Website:            church.Website,
		SocialMedia:        church.SocialMedia,
		PastorName:         church.ChurchName,
		PastorPhone:        church.PastorPhone,
		PastorEmail:        church.PastorEmail,
		FoundedYear:        church.FoundedYear,
		Description:        church.Description,
		DateAdded:          church.DateAdded,
		DateUpdated:        dateUpdated,
	}, nil
}

func (s *LocalChurchService) DeleteChurch(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid church ID")
	}

	if err := s.localChurchRepo.Delete(ctx, objID); err != nil {
		return fmt.Errorf("failed to delete church: %w", err)
	}
	return nil
}
