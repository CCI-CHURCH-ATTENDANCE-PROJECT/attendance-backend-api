package handler

import (
	"net/http"
	"strconv"

	"cci-api/internal/dto"
	"cci-api/internal/service"

	"github.com/labstack/echo/v4"
)

type FamilyMemberHandler struct {
	familyMemberService *service.FamilyMemberService
}

func NewFamilyMemberHandler(familyMemberService *service.FamilyMemberService) *FamilyMemberHandler {
	return &FamilyMemberHandler{familyMemberService: familyMemberService}
}

func (h *FamilyMemberHandler) CreateFamilyMember(c echo.Context) error {
	var req dto.CreateFamilyMemberRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "INVALID_REQUEST",
			Message: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    "VALIDATION_ERROR",
			Message: err.Error(),
		})
	}

	familyMember, err := h.familyMemberService.CreateFamilyMember(c.Request().Context(), c, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "FAMILY_MEMBER_CREATION_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{
		Code:    "FAMILY_MEMBER_CREATED",
		Message: "Family member created successfully",
		Data:    familyMember,
	})
}

func (h *FamilyMemberHandler) GetFamilyMembers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	familyMembers, err := h.familyMemberService.GetFamilyMembers(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "FAMILY_MEMBERS_FETCH_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "FAMILY_MEMBERS_RETRIEVED",
		Message: "Family members retrieved successfully",
		Data:    familyMembers,
	})
}

func (h *FamilyMemberHandler) GetFamilyMemberByID(c echo.Context) error {
	id := c.Param("id")

	familyMember, err := h.familyMemberService.GetFamilyMemberByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    "FAMILY_MEMBER_NOT_FOUND",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "FAMILY_MEMBER_RETRIEVED",
		Message: "Family member retrieved successfully",
		Data:    familyMember,
	})
}

func (h *FamilyMemberHandler) UpdateFamilyMember(c echo.Context) error {
	id := c.Param("id")
	var req dto.UpdateFamilyMemberRequest

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

	familyMember, err := h.familyMemberService.UpdateFamilyMember(c.Request().Context(), id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "FAMILY_MEMBER_UPDATE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "FAMILY_MEMBER_UPDATED",
		Message: "Family member updated successfully",
		Data:    familyMember,
	})
}

func (h *FamilyMemberHandler) DeleteFamilyMember(c echo.Context) error {
	id := c.Param("id")

	err := h.familyMemberService.DeleteFamilyMember(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    "FAMILY_MEMBER_DELETE_FAILED",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Code:    "FAMILY_MEMBER_DELETED",
		Message: "Family member deleted successfully",
	})
}
