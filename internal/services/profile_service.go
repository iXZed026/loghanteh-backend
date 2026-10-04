package services

import (
	"context"
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

	user, err := s.userRepository.FindByIDWithPassword(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	// Check email
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

	// Check phone number
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

	// Parse date of birth
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

	// Password change
	if req.CurrentPassword != "" {

		// Current password must be correct
		if !security.CheckPasswordHash(
			user.PasswordHash,
			req.CurrentPassword,
		) {
			return nil, appErrors.ErrInvalidPassword
		}

		// New password is required
		if req.NewPassword == "" {
			return nil, appErrors.ErrInvalidInput
		}
	}

	if req.NewPassword != "" {

		// Current password is required
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

	// Update profile
	err = s.userRepository.UpdateProfile(
		ctx,
		userID,
		req.FullName,
		req.Email,
		req.PhoneNumber,
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
