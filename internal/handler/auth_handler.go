package handler

import (
	"fmt"
	"net/http"

	"church-attendance-api/internal/dto"
	"church-attendance-api/internal/service"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) BasicRegister(c echo.Context) error {
	var req dto.BasicRegisterRequest
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

	// Generate qr code for the user on sign up
	// _, err := handler.NewQRHandler(c.Request().Context(), &req)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.APIResponse{
	// 		Success: false,
	// 		Error: &dto.ErrorInfo{
	// 			Code:    "QR_CODE_GENERATION_FAILED",
	// 			Message: err.Error(),
	// 		},
	// 	})
	// }

	resp, err := h.authService.BasicRegister(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "REGISTRATION_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    resp,
	})
}

func (h *AuthHandler) CompleteRegister(c echo.Context) error {
	var req dto.CompleteRegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "INVALID_REQUEST",
				// Message: "Invalid request body",
				Message: err.Error(),
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

	// Generate qr code for the user on sign up
	// _, err := handler.NewQRHandler(c.Request().Context(), &req)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.APIResponse{
	// 		Success: false,
	// 		Error: &dto.ErrorInfo{
	// 			Code:    "QR_CODE_GENERATION_FAILED",
	// 			Message: err.Error(),
	// 		},
	// 	})
	// }
	resp, err := h.authService.CompleteRegister(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "REGISTRATION_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    resp,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
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

	resp, err := h.authService.Login(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "LOGIN_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Login successful",
		Data:    resp,
	})
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var req dto.RefreshTokenRequest
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

	resp, err := h.authService.RefreshToken(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "TOKEN_REFRESH_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Data:    resp,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	userIDValue := c.Get("user_id")
	fmt.Printf("Handler sees user_id: %#v\n", userIDValue)
	userID, ok := userIDValue.(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "INVALID_USER_ID",
				Message: "User ID missing or invalid token.",
			},
		})
	}
	fmt.Printf("Logout process has begun>>>")
	err := h.authService.Logout(c.Request().Context(), userID)
	fmt.Printf("Logout failed due to this error %v", err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Error: &dto.ErrorInfo{
				Code:    "LOGOUT_FAILED",
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Logged out successfully",
	})
}
