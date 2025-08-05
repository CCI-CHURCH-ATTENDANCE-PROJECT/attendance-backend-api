package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"cci-api/internal/dto"
	"cci-api/internal/service"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnnouncementHandler struct {
	announcementService *service.AnnouncementService
}

func NewAnnouncementHandler(announcementService *service.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{announcementService: announcementService}
}

func (h *AnnouncementHandler) CreateAnnouncement(c echo.Context) error {
	var req dto.CreateAnnouncementRequest
	fmt.Printf("This is the reqest body passed: %v\n", req)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: "Invalid request body",
		})
	}
	fmt.Printf("This is the request body passed: %v\n", req)

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: err.Error(),
		})
	}

	// Get user ID from JWT token
	userID := c.Get("user_id").(string)
	objID, _ := primitive.ObjectIDFromHex(userID)

	announcement, err := h.announcementService.CreateAnnouncement(c.Request().Context(), &req, objID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ANNOUNCEMENT_CREATION_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{
		Code:    "ANNOUNCEMENT_CREATED",
		Message: "Announcement created successfully",
		Data:    announcement,
	})
}

func (h *AnnouncementHandler) GetAnnouncements(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	announcements, err := h.announcementService.GetAnnouncements(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ANNOUNCEMENTS_FETCH_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ANNOUNCEMENTS_RETRIEVED",
		Message: "Announcements retrieved successfully",
		Data:    announcements,
	})
}

func (h *AnnouncementHandler) GetAnnouncementByID(c echo.Context) error {
	id := c.Param("id")

	announcement, err := h.announcementService.GetAnnouncementByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    "ANNOUNCEMENT_NOT_FOUND",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ANNOUNCEMENT_RETRIEVED",
		Message: "Announcement retrieved successfully",
		Data:    announcement,
	})
}

func (h *AnnouncementHandler) UpdateAnnouncement(c echo.Context) error {
	id := c.Param("id")
	var req dto.UpdateAnnouncementRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: "Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: "Validation failed",
		})
	}

	announcement, err := h.announcementService.UpdateAnnouncement(c.Request().Context(), id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ANNOUNCEMENT_UPDATE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ANNOUNCEMENT_UPDATED",
		Message: "Announcement updated successfully",
		Data:    announcement,
	})
}

func (h *AnnouncementHandler) DeleteAnnouncement(c echo.Context) error {
	id := c.Param("id")

	err := h.announcementService.DeleteAnnouncement(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ANNOUNCEMENT_DELETE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ANNOUNCEMENT_DELETED",
		Message: "Announcement deleted successfully",
	})
}

func (h *AnnouncementHandler) GetActiveAnnouncements(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	announcements, err := h.announcementService.GetActiveAnnouncements(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ACTIVE_ANNOUNCEMENTS_FETCH_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ACTIVE_ANNOUNCEMENTS_RETRIEVED",
		Message: "Active announcements retrieved successfully",
		Data:    announcements,
	})
}
