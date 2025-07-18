package handler

import (
	"net/http"
	"strconv"

	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/service"

	"github.com/labstack/echo/v4"
)

type LocalChurchHandler struct {
	localChurchService *service.LocalChurchService
}

func NewLocalChurchHandler(localChurchService *service.LocalChurchService) *LocalChurchHandler {
	return &LocalChurchHandler{localChurchService: localChurchService}
}

func (h *LocalChurchHandler) CreateChurch(c echo.Context) error {
	var req dto.CreateLocalChurchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: "Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: err.Error(),
			// Message: "Validation failed",
		})
	}

	church, err := h.localChurchService.CreateChurch(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "CHURCH_CREATION_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{
		Code:    "CHURCH_CREATED",
		Message: "Church created successfully",
		Data:    church,
	})
}

func (h *LocalChurchHandler) GetChurches(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	churches, err := h.localChurchService.GetChurches(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "CHURCHES_FETCH_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "CHURCHES_RETRIEVED",
		Message: "Churches retrieved successfully",
		Data:    churches,
	})
}

func (h *LocalChurchHandler) GetChurchByID(c echo.Context) error {
	id := c.Param("id")

	church, err := h.localChurchService.GetChurchByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    "CHURCH_NOT_FOUND",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "CHURCH_RETRIEVED",
		Message: "Church retrieved successfully",
		Data:    church,
	})
}

func (h *LocalChurchHandler) UpdateChurch(c echo.Context) error {
	id := c.Param("id")
	var req dto.UpdateLocalChurchRequest

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

	church, err := h.localChurchService.UpdateChurch(c.Request().Context(), id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "CHURCH_UPDATE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "CHURCH_UPDATED",
		Message: "Church updated successfully",
		Data:    church,
	})
}

func (h *LocalChurchHandler) DeleteChurch(c echo.Context) error {
	id := c.Param("id")

	err := h.localChurchService.DeleteChurch(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "CHURCH_DELETE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "CHURCH_DELETED",
		Message: "Church deleted successfully",
	})
}
