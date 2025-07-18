package handler

import (
	"net/http"
	"time"

	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/service"
	"church-attendance-api/internal/utils"

	"github.com/labstack/echo/v4"
)

type AttendanceHandler struct {
	attendanceService *service.AttendanceService
}

func NewAttendanceHandler(attendanceService *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

func (h *AttendanceHandler) CreateAttendance(c echo.Context) error {
	var req dto.CreateAttendanceRequest
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

	resp, err := h.attendanceService.CreateAttendance(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "ATTENDANCE_CREATION_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Attendance recorded successfully",
		Data:    resp,
	})
}

func (h *AttendanceHandler) QRCheckin(c echo.Context) error {
	var req dto.QRCheckinRequest
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

	resp, err := h.attendanceService.QRCheckin(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "QR_CHECKIN_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "QR check-in successful",
		Data:    resp,
	})
}

func (h *AttendanceHandler) GetAttendanceHistory(c echo.Context) error {
	var startDate, endDate *time.Time

	// Parse start_date if provided
	if startDateStr := c.QueryParam("start_date"); startDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsed
		} else {
			return c.JSON(http.StatusBadRequest, dto.APIResponse{
				Success: false,
				Error: &dto.ErrorInfo{
					Code:    "INVALID_DATE_FORMAT",
					Message: "start_date must be in YYYY-MM-DD format",
				},
			})
		}
	}

	// Parse end_date if provided
	if endDateStr := c.QueryParam("end_date"); endDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsed
		} else {
			return c.JSON(http.StatusBadRequest, dto.APIResponse{
				Success: false,
				Error: &dto.ErrorInfo{
					Code:    "INVALID_DATE_FORMAT",
					Message: "end_date must be in YYYY-MM-DD format",
				},
			})
		}
	}

	page := utils.StringToInt(c.QueryParam("page"), 1)
	limit := utils.StringToInt(c.QueryParam("limit"), 10)

	resp, err := h.attendanceService.GetAttendanceHistory(c.Request().Context(), startDate, endDate, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "HISTORY_FETCH_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    resp,
	})
}

func (h *AttendanceHandler) GetAttendanceAnalytics(c echo.Context) error {
	dateStr := c.QueryParam("date")
	if dateStr == "" {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "MISSING_DATE",
				Message: "date query parameter is required (YYYY-MM-DD format)",
			},
		})
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "INVALID_DATE_FORMAT",
				Message: "date must be in YYYY-MM-DD format",
			},
		})
	}

	resp, err := h.attendanceService.GetAttendanceAnalytics(c.Request().Context(), date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "ANALYTICS_FETCH_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    resp,
	})
}
