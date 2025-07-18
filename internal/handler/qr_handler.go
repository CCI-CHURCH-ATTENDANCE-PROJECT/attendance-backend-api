package handler

import (
	"net/http"

	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/service"

	"github.com/labstack/echo/v4"
)

type QRHandler struct {
	qrService *service.QRService
}

func NewQRHandler(qrService *service.QRService) *QRHandler {
	return &QRHandler{
		qrService: qrService,
	}
}

func (h *QRHandler) GenerateQRCode(c echo.Context) error {
	var req dto.GenerateQRRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "INVALID_REQUEST",
				Message: "Invalid request body",
			},
		})
	}

	// Validate request
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "VALIDATION_ERROR",
				Message: "Validation failed",
				Details: []dto.ErrorDetail{
					{Field: "request", Message: err.Error()},
				},
			},
		})
	}

	resp, err := h.qrService.GenerateQRCode(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "QR_GENERATION_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    resp,
	})
}
