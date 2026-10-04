package services

import (
	"context"
	"strings"

	"loghanteh-project/internal/dto"
	appErrors "loghanteh-project/internal/errors"
	"loghanteh-project/internal/models"
	"loghanteh-project/internal/repositories"
)

type CustomerRequestService struct {
	customerRequestRepo *repositories.CustomerRequestRepository
}

func NewCustomerRequestService(
	customerRequestRepo *repositories.CustomerRequestRepository,
) *CustomerRequestService {
	return &CustomerRequestService{
		customerRequestRepo: customerRequestRepo,
	}
}

func (s *CustomerRequestService) CreateCustomerRequest(
	ctx context.Context,
	req *dto.CustomerRequestReq,
) error {

	var email string
	var phoneNumber string

	if strings.Contains(req.EmailOrPhone, "@") {
		email = req.EmailOrPhone
	} else {
		phoneNumber = req.EmailOrPhone
	}

	customerRequest := &models.CustomerRequest{
		FullName:     req.FullName,
		CompanyName:  req.CompanyName,
		Email:        email,
		PhoneNumber:  phoneNumber,
		RequestTitle: req.RequestTitle,
		Description:  req.Description,
	}

	if err := s.customerRequestRepo.
		CreateCustomerRequest(
			customerRequest,
			ctx,
		); err != nil {
		return appErrors.ErrInternalServer
	}

	return nil
}
