package handlers

import (
	"log/slog"
	"net/http"

	"loghanteh-project/internal/dto"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/response"
	"loghanteh-project/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
	logger      *slog.Logger
}

func NewAuthHandler(
	authService *services.AuthService,
	logger *slog.Logger,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"failed to bind register request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	user, err := h.authService.Register(
		c.Request.Context(),
		req.FullName,
		req.PhoneNumber,
		req.Email,
		req.Password,
		req.RepeatPassword,
	)

	if err != nil {

		h.logger.Error(
			"failed to register user",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		"verification_required",
		user,
	)
}

func (h *AuthHandler) VerifyRegister(c *gin.Context) {

	var req dto.VerifyRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"failed to bind register verification request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	user, err := h.authService.VerifyRegister(
		c.Request.Context(),
		req.VerificationToken,
		req.Code,
	)

	if err != nil {

		h.logger.Error(
			"failed to verify register",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		"user_registration_verified_successfully",
		user,
	)
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"user login bad request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	user, err :=
		h.authService.Login(
			c.Request.Context(),
			req.EmailOrPhone,
			req.Password,
		)

	if err != nil {

		h.logger.Error(
			"failed to login user",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		"login_is_successful",
		user,
	)
}

func (h *AuthHandler) Refresh(c *gin.Context) {

	refreshToken, err := c.Cookie(
		"refresh_token",
	)

	if err != nil {

		response.HandleError(
			c,
			appErrors.ErrInvalidCredentials,
			h.logger,
		)

		return
	}

	accessToken, newRefreshToken, err :=
		h.authService.RefreshAccessToken(
			c.Request.Context(),
			refreshToken,
		)

	if err != nil {

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		30*24*60*60,
		"/",
		"",
		true,
		true,
	)

	response.Success(
		c,
		"token_refreshed_successfully",
		gin.H{
			"accessToken": accessToken,
		},
	)
}

func (h *AuthHandler) Logout(c *gin.Context) {

	refreshToken, err := c.Cookie(
		"refresh_token",
	)

	if err == nil {

		err = h.authService.Logout(
			c.Request.Context(),
			refreshToken,
		)

		if err != nil {

			response.HandleError(
				c,
				err,
				h.logger,
			)

			return
		}
	}

	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	response.Success(
		c,
		"logged_out_successfully",
		nil,
	)
}

func (h *AuthHandler) VerifyLogin(c *gin.Context) {

	var req dto.VerifyLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"failed to get verify login request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	user, refreshToken, err :=
		h.authService.VerifyLogin(
			c.Request.Context(),
			req.VerificationToken,
			req.Code,
		)

	if err != nil {

		h.logger.Error(
			"failed to verify login",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	// Refresh Token → HttpOnly Cookie
	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie(
		"refresh_token",
		refreshToken,
		30*24*60*60,
		"/",
		"",
		true,
		true,
	)

	response.Success(
		c,
		"user_login_verified_successfully",
		user,
	)
}

// Forget Password
func (h *AuthHandler) ForgetPassword(c *gin.Context) {

	var req dto.ForgetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"forget password bad request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	verification, err := h.authService.ForgetPassword(
		c.Request.Context(),
		req.EmailOrPhone,
	)

	if err != nil {

		h.logger.Error(
			"forget password failed",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		"verification_code_sent_successfully",
		verification,
	)
}

func (h *AuthHandler) VerifyForgetPassword(c *gin.Context) {

	var req dto.VerifyForgetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"verify forget password bad request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	verificationData, err := h.authService.VerifyForgetPassword(
		c.Request.Context(),
		req.VerificationToken,
		req.Code,
	)

	if err != nil {

		h.logger.Error(
			"verify forget password failed",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		"verification_successful",
		verificationData,
	)
}

func (h *AuthHandler) NewPassword(c *gin.Context) {

	var req dto.NewPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.logger.Error(
			"new password bad request",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	newPassword, err := h.authService.SetNewPassword(
		c.Request.Context(),
		req.VerificationToken,
		req.Password,
		req.RepeatPassword,
	)

	if err != nil {

		h.logger.Error(
			"failed to set new password",
			"error", err,
		)

		response.HandleError(
			c,
			err,
			h.logger,
		)

		return
	}

	response.Success(
		c,
		"password_changed_successfully",
		newPassword,
	)
}
