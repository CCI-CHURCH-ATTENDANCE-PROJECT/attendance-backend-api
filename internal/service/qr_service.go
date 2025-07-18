package service

import (
	"context"
	"errors"
	"fmt"

	"church-attendance-api/internal/config"
	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/repository"
	"church-attendance-api/internal/utils"
)

type QRService struct {
	cfg      *config.Config
	userRepo *repository.UserRepository
}

func NewQRService(cfg *config.Config, userRepo *repository.UserRepository) *QRService {
	return &QRService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (s *QRService) GenerateQRCode(ctx context.Context, req *dto.GenerateQRRequest) (*dto.QRCodeResponse, error) {
	// Get user
	user, err := s.userRepo.GetByUserID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Generate QR code token
	token, err := utils.GenerateRandomToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR token: %w", err)
	}

	// Update user with QR token
	err = s.userRepo.UpdateQRToken(ctx, user.UserID, token)
	if err != nil {
		return nil, fmt.Errorf("failed to update user QR token: %w", err)
	}

	// Generate QR code image
	qrImage, err := utils.GenerateQRCode(token, s.cfg.QRCodeSize)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code image: %w", err)
	}

	// Update user with QR code image string
	err = s.userRepo.UpdateQRCodeImage(ctx, user.UserID, qrImage)
	if err != nil {
		return nil, fmt.Errorf("failed to update user QR code image: %w", err)
	}


	return &dto.QRCodeResponse{
		QRCodeToken: token,
		QRCodeImage: qrImage,
	}, nil
}
