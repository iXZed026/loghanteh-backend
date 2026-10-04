package services

import (
	"context"
	"strconv"
	"time"

	appAuth "loghanteh-project/internal/auth"
	"loghanteh-project/internal/dto"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"
	"loghanteh-project/internal/repositories"
	"loghanteh-project/internal/security"
)

type AuthService struct {
	userRepository                       *repositories.UserRepository
	refreshTokenRepository               *repositories.RefreshTokenRepository
	jwtService                           *appAuth.JWTService
	loginVerificationRepository          *repositories.LoginVerificationRepository
	registerVerificationRepository       *repositories.RegisterVerificationRepository
	forgetPasswordVerificationRepository *repositories.ForgetPasswordVerificationRepository
}

func NewAuthService(
	userRepository *repositories.UserRepository,
	refreshTokenRepository *repositories.RefreshTokenRepository,
	jwtService *appAuth.JWTService,
	loginVerificationRepository *repositories.LoginVerificationRepository,
	registerVerificationRepository *repositories.RegisterVerificationRepository,
	forgetPasswordVerificationRepository *repositories.ForgetPasswordVerificationRepository,
) *AuthService {
	return &AuthService{
		userRepository:                       userRepository,
		refreshTokenRepository:               refreshTokenRepository,
		jwtService:                           jwtService,
		loginVerificationRepository:          loginVerificationRepository,
		registerVerificationRepository:       registerVerificationRepository,
		forgetPasswordVerificationRepository: forgetPasswordVerificationRepository,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	fullName string,
	phoneNumber string,
	email string,
	password string,
	repeatPassword string,
) (*dto.RegisterResponse, error) {
	if password != repeatPassword {
		return nil, appErrors.ErrPasswordMismatch
	}

	emailExists, err := s.userRepository.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	if emailExists {
		return nil, appErrors.ErrEmailAlreadyExists
	}

	phoneExists, err := s.userRepository.ExistsByPhone(ctx, phoneNumber)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	if phoneExists {
		return nil, appErrors.ErrPhoneAlreadyExists
	}

	passwordHash, err := security.HashPassword(password)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	verificationToken, err := appAuth.GenerateVerificationToken()
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	verification := &models.RegisterVerification{
		FullName:     fullName,
		PhoneNumber:  phoneNumber,
		Email:        email,
		PasswordHash: passwordHash,

		Token: security.HashToken(verificationToken),

		ExpiresAt: time.Now().Add(5 * time.Minute),

		Verified: false,
	}

	err = s.registerVerificationRepository.Create(ctx, verification)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	return &dto.RegisterResponse{
		FullName:          fullName,
		Email:             email,
		PhoneNumber:       phoneNumber,
		VerificationToken: verificationToken,
	}, nil
}

func (s *AuthService) VerifyRegister(
	ctx context.Context,
	verificationToken string,
	code string,
) (*dto.RegisterResponse, error) {
	verification, err :=
		s.registerVerificationRepository.FindByToken(
			ctx,
			security.HashToken(verificationToken),
		)
	if err != nil {
		return nil, err
	}

	if verification.Verified {
		return nil, appErrors.ErrUserAlreadyVerified
	}

	if time.Now().After(verification.ExpiresAt) {
		return nil, appErrors.ErrVerificationTokenExpired
	}

	if code != "123456" {
		return nil, appErrors.ErrInvalidVerificationCode
	}

	user := &models.User{
		FullName:     verification.FullName,
		PhoneNumber:  verification.PhoneNumber,
		Email:        verification.Email,
		PasswordHash: verification.PasswordHash,
	}

	err = s.userRepository.CreateUser(user, ctx)
	if err != nil {
		return nil, err
	}

	err = s.registerVerificationRepository.MarkVerified(
		ctx,
		verification,
	)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	return &dto.RegisterResponse{
		UserID:            user.UserID,
		FullName:          user.FullName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		VerificationToken: verificationToken,
	}, nil
}

func (s *AuthService) Login(
	ctx context.Context,
	emailOrPhone string,
	password string,
) (*dto.LoginResponse, error) {
	user, err := s.userRepository.FindByEmailOrPhone(ctx, emailOrPhone)
	if err != nil {
		return nil, err
	}

	if !security.CheckPasswordHash(user.PasswordHash, password) {
		return nil, appErrors.ErrInvalidCredentials
	}

	verificationToken, err := appAuth.GenerateVerificationToken()
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	verification := &models.LoginVerification{
		UserID:    user.UserID,
		Token:     security.HashToken(verificationToken),
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Verified:  false,
	}

	err = s.loginVerificationRepository.Create(ctx, verification)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	return &dto.LoginResponse{
		UserID:            user.UserID,
		FullName:          user.FullName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		VerificationToken: verificationToken,
	}, nil
}

func (s *AuthService) VerifyLogin(
	ctx context.Context,
	verificationToken string,
	code string,
) (*dto.VerifyLoginResponse, string, error) {
	verification, err := s.loginVerificationRepository.FindByToken(
		ctx,
		security.HashToken(verificationToken),
	)
	if err != nil {
		return nil, "", err
	}

	if verification.Verified {
		return nil, "", appErrors.ErrUserAlreadyVerified
	}

	if time.Now().After(verification.ExpiresAt) {
		return nil, "", appErrors.ErrVerificationTokenExpired
	}

	// Temporary OTP
	// بعداً این قسمت را با OTP واقعی جایگزین می‌کنیم.
	if code != "123456" {
		return nil, "", appErrors.ErrInvalidVerificationCode
	}

	err = s.loginVerificationRepository.MarkVerified(
		ctx,
		verification,
	)
	if err != nil {
		return nil, "", appErrors.ErrInternalServer
	}

	user, err := s.userRepository.FindByID(
		ctx,
		uint(verification.UserID),
	)
	if err != nil {
		return nil, "", err
	}

	accessToken, err := s.jwtService.GenerateAccessToken(
		strconv.Itoa(user.UserID),
	)
	if err != nil {
		return nil, "", appErrors.ErrInternalServer
	}

	refreshToken, err := appAuth.GenerateRefreshToken()
	if err != nil {
		return nil, "", appErrors.ErrInternalServer
	}

	refreshTokenModel := &models.RefreshToken{
		UserID:    verification.UserID,
		Token:     security.HashToken(refreshToken),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		Revoked:   false,
	}

	err = s.refreshTokenRepository.Create(
		ctx,
		refreshTokenModel,
	)
	if err != nil {
		return nil, "", appErrors.ErrInternalServer
	}

	return &dto.VerifyLoginResponse{
		UserID:      user.UserID,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		AccessToken: accessToken,
	}, refreshToken, nil
}

func (s *AuthService) RefreshAccessToken(
	ctx context.Context,
	oldRefreshToken string,
) (string, string, error) {
	oldTokenHash := security.HashToken(oldRefreshToken)

	token, err := s.refreshTokenRepository.FindValidToken(
		ctx,
		oldTokenHash,
	)
	if err != nil {
		return "", "", err
	}

	err = s.refreshTokenRepository.Revoke(ctx, oldTokenHash)
	if err != nil {
		return "", "", appErrors.ErrInternalServer
	}

	accessToken, err := s.jwtService.GenerateAccessToken(
		strconv.Itoa(token.UserID),
	)
	if err != nil {
		return "", "", appErrors.ErrInternalServer
	}

	newRefreshToken, err := appAuth.GenerateRefreshToken()
	if err != nil {
		return "", "", appErrors.ErrInternalServer
	}

	newRefreshTokenModel := &models.RefreshToken{
		UserID:    token.UserID,
		Token:     security.HashToken(newRefreshToken),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		Revoked:   false,
	}

	err = s.refreshTokenRepository.Create(ctx, newRefreshTokenModel)
	if err != nil {
		return "", "", appErrors.ErrInternalServer
	}

	return accessToken, newRefreshToken, nil
}

func (s *AuthService) Logout(
	ctx context.Context,
	refreshToken string,
) error {
	err := s.refreshTokenRepository.Revoke(
		ctx,
		security.HashToken(refreshToken),
	)
	if err != nil {
		return appErrors.ErrInternalServer
	}

	return nil
}

// Forget Password
func (s *AuthService) ForgetPassword(
	ctx context.Context,
	emailOrPhone string,
) (*dto.ForgetPasswordResponse, error) {

	user, err := s.userRepository.FindByEmailOrPhone(
		ctx,
		emailOrPhone,
	)
	if err != nil {
		return nil, err
	}

	verificationToken, err := appAuth.GenerateVerificationToken()
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	verification := &models.ForgetPasswordVerification{
		UserID:    user.UserID,
		TokenHash: security.HashToken(verificationToken),
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Verified:  false,
	}

	if err := s.forgetPasswordVerificationRepository.Create(
		ctx,
		verification,
	); err != nil {
		return nil, appErrors.ErrInternalServer
	}

	return &dto.ForgetPasswordResponse{
		UserID:            user.UserID,
		FullName:          user.FullName,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		VerificationToken: verificationToken,
	}, nil
}

func (s *AuthService) VerifyForgetPassword(
	ctx context.Context,
	verificationToken string,
	code string,
) (*dto.VerifyForgetPasswordResponse, error) {

	verification, err :=
		s.forgetPasswordVerificationRepository.FindValidToken(
			ctx,
			security.HashToken(verificationToken),
		)
	if err != nil {
		return nil, err
	}

	if verification.Verified {
		return nil, appErrors.ErrUserAlreadyVerified
	}

	if time.Now().After(verification.ExpiresAt) {
		return nil, appErrors.ErrVerificationTokenExpired
	}

	if code != "123456" {
		return nil, appErrors.ErrInvalidVerificationCode
	}

	err = s.forgetPasswordVerificationRepository.MarkVerified(
		ctx,
		verification,
	)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	user, err := s.userRepository.FindByID(
		ctx,
		uint(verification.UserID),
	)
	if err != nil {
		return nil, err
	}

	return &dto.VerifyForgetPasswordResponse{
		UserID:      user.UserID,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (s *AuthService) SetNewPassword(
	ctx context.Context,
	verificationToken string,
	password string,
	repeatPassword string,
) (*dto.NewPasswordResponse, error) {

	if password != repeatPassword {
		return nil, appErrors.ErrPasswordMismatch
	}

	verification, err :=
		s.forgetPasswordVerificationRepository.FindVerifiedToken(
			ctx,
			security.HashToken(verificationToken),
		)
	if err != nil {
		return nil, err
	}

	if time.Now().After(verification.ExpiresAt) {
		return nil, appErrors.ErrVerificationTokenExpired
	}

	passwordHash, err := security.HashPassword(password)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	err = s.userRepository.UpdatePassword(
		ctx,
		uint(verification.UserID),
		passwordHash,
	)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	return &dto.NewPasswordResponse{
		UserID: uint(verification.UserID),
	}, nil
}
