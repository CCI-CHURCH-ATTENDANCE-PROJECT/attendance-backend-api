package handler

import (
	"net/http"

	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/service"
	"church-attendance-api/internal/utils"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) SearchUsers(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "MISSING_QUERY",
				Message: "Search query parameter 'q' is required",
			},
		})
	}

	page := utils.StringToInt(c.QueryParam("page"), 1)
	limit := utils.StringToInt(c.QueryParam("limit"), 10)

	resp, err := h.userService.SearchUsers(c.Request().Context(), query, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "SEARCH_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    resp,
	})
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	page := utils.StringToInt(c.QueryParam("page"), 1)
	limit := utils.StringToInt(c.QueryParam("limit"), 10)

	resp, err := h.userService.GetAllUsers(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "FETCH_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    resp,
	})
}

func (h *UserHandler) FilterUsers(c echo.Context) error {
	field := c.QueryParam("field")
	value := c.QueryParam("value")

	if field == "" || value == "" {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "MISSING_PARAMETERS",
				Message: "Both 'field' and 'value' query parameters are required",
			},
		})
	}

	page := utils.StringToInt(c.QueryParam("page"), 1)
	limit := utils.StringToInt(c.QueryParam("limit"), 10)

	resp, err := h.userService.FilterUsers(c.Request().Context(), field, value, page, limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "FILTER_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    resp,
	})
}
