package services

import (
	"context"
	"strings"
	"time"

	"loghanteh-project/internal/dto"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"
	"loghanteh-project/internal/repositories"
	"loghanteh-project/internal/security"
)

type ProfileService struct {
	userRepository *repositories.UserRepository
}

func NewProfileService(
	userRepository *repositories.UserRepository,
) *ProfileService {
	return &ProfileService{
		userRepository: userRepository,
	}
}

func (s *ProfileService) GetProfile(
	ctx context.Context,
	userID uint,
) (*models.User, error) {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, appErrors.ErrInternalServer
	}

	return user, nil
}

func (s *ProfileService) EditProfile(
	ctx context.Context,
	userID uint,
	req dto.EditProfileRequest,
) (*dto.EditProfileResponse, error) {

	req.FullName = strings.TrimSpace(req.FullName)
	req.Email = strings.TrimSpace(req.Email)
	req.PhoneNumber = strings.TrimSpace(req.PhoneNumber)

	if req.Email == "" && req.PhoneNumber == "" {
		return nil, appErrors.ErrAfieldForEitherEmailOrPhoneNumber
	}

	user, err := s.userRepository.FindByIDWithPassword(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		emailExists, err := s.userRepository.ExistsByEmailExceptUser(
			ctx,
			req.Email,
			userID,
		)

		if err != nil {
			return nil, appErrors.ErrInternalServer
		}

		if emailExists {
			return nil, appErrors.ErrEmailAlreadyExists
		}
	}

	if req.PhoneNumber != "" {
		phoneExists, err := s.userRepository.ExistsByPhoneExceptUser(
			ctx,
			req.PhoneNumber,
			userID,
		)

		if err != nil {
			return nil, appErrors.ErrInternalServer
		}

		if phoneExists {
			return nil, appErrors.ErrPhoneAlreadyExists
		}
	}

	var dob *time.Time

	if req.DOB != "" {
		parsedDOB, err := time.Parse(
			"2006-01-02",
			req.DOB,
		)

		if err != nil {
			return nil, appErrors.ErrInvalidInput
		}

		dob = &parsedDOB
	}

	var passwordHash *string

	if req.CurrentPassword != "" {
		if !security.CheckPasswordHash(
			user.PasswordHash,
			req.CurrentPassword,
		) {
			return nil, appErrors.ErrInvalidPassword
		}

		if req.NewPassword == "" {
			return nil, appErrors.ErrInvalidInput
		}
	}

	if req.NewPassword != "" {
		if req.CurrentPassword == "" {
			return nil, appErrors.ErrInvalidInput
		}

		hash, err := security.HashPassword(
			req.NewPassword,
		)

		if err != nil {
			return nil, appErrors.ErrInternalServer
		}

		passwordHash = &hash
	}

	var email *string
	if req.Email != "" {
		email = &req.Email
	}

	var phoneNumber *string
	if req.PhoneNumber != "" {
		phoneNumber = &req.PhoneNumber
	}

	err = s.userRepository.UpdateProfile(
		ctx,
		userID,
		req.FullName,
		email,
		phoneNumber,
		dob,
		passwordHash,
	)

	if err != nil {
		return nil, err
	}

	return &dto.EditProfileResponse{
		UserID:      userID,
		FullName:    req.FullName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		DOB:         dob,
	}, nil
}
