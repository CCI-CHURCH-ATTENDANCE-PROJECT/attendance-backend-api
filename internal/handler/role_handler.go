package handler

import (
	"net/http"
	"strconv"

	"cci-api/internal/dto"
	"cci-api/internal/service"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) CreateRole(c echo.Context) error {
	var req dto.CreateRoleRequest
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

	role, err := h.roleService.CreateRole(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ROLE_CREATION_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{
		Code:    "ROLE_CREATED",
		Message: "Role created successfully",
		Data:    role,
	})
}

func (h *RoleHandler) GetRoles(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	roles, err := h.roleService.GetRoles(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ROLES_FETCH_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ROLES_RETRIEVED",
		Message: "Roles retrieved successfully",
		Data:    roles,
	})
}

func (h *RoleHandler) GetRoleByID(c echo.Context) error {
	id := c.Param("id")

	role, err := h.roleService.GetRoleByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    "ROLE_NOT_FOUND",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ROLE_RETRIEVED",
		Message: "Role retrieved successfully",
		Data:    role,
	})
}

func (h *RoleHandler) UpdateRole(c echo.Context) error {
	id := c.Param("id")
	var req dto.UpdateRoleRequest

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

	role, err := h.roleService.UpdateRole(c.Request().Context(), id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ROLE_UPDATE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ROLE_UPDATED",
		Message: "Role updated successfully",
		Data:    role,
	})
}

func (h *RoleHandler) DeleteRole(c echo.Context) error {
	id := c.Param("id")

	err := h.roleService.DeleteRole(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "ROLE_DELETE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "ROLE_DELETED",
		Message: "Role deleted successfully",
	})
}
