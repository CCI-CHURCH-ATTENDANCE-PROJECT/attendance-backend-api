package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/service"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SermonHandler struct {
	sermonService *service.SermonService
}

func NewSermonHandler(sermonService *service.SermonService) *SermonHandler {
	return &SermonHandler{sermonService: sermonService}
}

func (h *SermonHandler) CreateSermon(c echo.Context) error {
	var req dto.CreateSermonRequest
	if err := c.Bind(&req); err != nil {
		fmt.Printf("Binding error: %v\n", err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: "Invalid request body",
		})
	}
	fmt.Printf("Received Request: %+v\n", req)

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: "Validation failed",
		})
	}

	// Get user ID from JWT token
	userID := c.Get("user_id").(string)
	objID, _ := primitive.ObjectIDFromHex(userID)

	sermon, err := h.sermonService.CreateSermon(c.Request().Context(), &req, objID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "SERMON_CREATION_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{
		Code:    "SERMON_CREATED",
		Message: "Sermon created successfully",
		Data:    sermon,
	})
}

func (h *SermonHandler) GetSermons(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	sermons, err := h.sermonService.GetSermons(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "SERMONS_FETCH_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "SERMONS_RETRIEVED",
		Message: "Sermons retrieved successfully",
		Data:    sermons,
	})
}

func (h *SermonHandler) GetSermonByID(c echo.Context) error {
	id := c.Param("id")

	sermon, err := h.sermonService.GetSermonByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    "SERMON_NOT_FOUND",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "SERMON_RETRIEVED",
		Message: "Sermon retrieved successfully",
		Data:    sermon,
	})
}

func (h *SermonHandler) UpdateSermon(c echo.Context) error {
	id := c.Param("id")
	var req dto.UpdateSermonRequest

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

	sermon, err := h.sermonService.UpdateSermon(c.Request().Context(), id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "SERMON_UPDATE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "SERMON_UPDATED",
		Message: "Sermon updated successfully",
		Data:    sermon,
	})
}

func (h *SermonHandler) DeleteSermon(c echo.Context) error {
	id := c.Param("id")

	err := h.sermonService.DeleteSermon(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "SERMON_DELETE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "SERMON_DELETED",
		Message: "Sermon deleted successfully",
	})
}
