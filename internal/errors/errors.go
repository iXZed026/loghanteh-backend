package errors

import "net/http"

var (
	// Global Errors

	ErrEmptyValue = &AppError{
		Code: http.StatusBadRequest,
		Key:  "value_must_not_be_empty",
	}

	ErrInternalServer = &AppError{
		Code: http.StatusInternalServerError,
		Key:  "internal_server_error",
	}

	ErrInvalidInput = &AppError{
		Code: http.StatusBadRequest,
		Key:  "invalid_input",
	}

	ErrUnauthorized = &AppError{
		Code: http.StatusUnauthorized,
		Key:  "unauthorized",
	}

	ErrAfieldForEitherEmailOrPhoneNumber = &AppError{
		Code: http.StatusBadRequest,
		Key:  "either_email_phonenumber",
	}

	ErrUserAlreadyVerified = &AppError{
		Code: http.StatusConflict,
		Key:  "user_already_verified",
	}

	ErrVerificationTokenExpired = &AppError{
		Code: http.StatusGone,
		Key:  "verification_token_expired",
	}

	ErrInvalidVerificationCode = &AppError{
		Code: http.StatusBadRequest,
		Key:  "invalid_verification_code",
	}

	ErrPasswordMismatch = &AppError{
		Code: http.StatusBadRequest,
		Key:  "passwords_do_not_match",
	}

	ErrTooManyRequests = &AppError{
		Code: http.StatusTooManyRequests,
		Key:  "too_many_requests",
	}

	ErrEmailAlreadyExists = &AppError{
		Code: http.StatusConflict,
		Key:  "email_already_exists",
	}

	ErrPhoneAlreadyExists = &AppError{
		Code: http.StatusConflict,
		Key:  "phone_number_already_exists",
	}

	ErrInvalidCredentials = &AppError{
		Code: http.StatusUnauthorized,
		Key:  "invalid_credentials",
	}

	ErrInvalidPassword = &AppError{
		Code: http.StatusBadRequest,
		Key:  "invalid_password",
	}

	ErrUserNotFound = &AppError{
		Code: http.StatusNotFound,
		Key:  "user_not_found",
	}

	ErrNotFound = &AppError{
		Code: http.StatusNotFound,
		Key:  "not_found",
	}

	ErrInsufficientCapacity = &AppError{
		Code: http.StatusConflict,
		Key:  "insufficient_capacity",
	}

	// --------------------------------------------------
	// Discount Code Errors
	// --------------------------------------------------

	ErrDiscountCodeNotFound = &AppError{
		Code: http.StatusNotFound,
		Key:  "discount_code_not_found",
	}

	ErrDiscountCodeInactive = &AppError{
		Code: http.StatusConflict,
		Key:  "discount_code_inactive",
	}

	ErrDiscountCodeNotStarted = &AppError{
		Code: http.StatusConflict,
		Key:  "discount_code_not_started",
	}

	ErrDiscountCodeExpired = &AppError{
		Code: http.StatusGone,
		Key:  "discount_code_expired",
	}

	ErrDiscountCodeUsageLimitReached = &AppError{
		Code: http.StatusConflict,
		Key:  "discount_code_usage_limit_reached",
	}

	// --------------------------------------------------
	// Reservation Errors
	// --------------------------------------------------

	ErrNoResources = &AppError{
		Code: http.StatusBadRequest,
		Key:  "no_resources_provided",
	}

	ErrDuplicateResources = &AppError{
		Code: http.StatusBadRequest,
		Key:  "duplicate_resources_provided",
	}

	ErrResourceNotFound = &AppError{
		Code: http.StatusBadRequest,
		Key:  "resource_not_found",
	}

	ErrResourceAlreadyReserved = &AppError{
		Code: http.StatusConflict,
		Key:  "resource_already_reserved",
	}
)
