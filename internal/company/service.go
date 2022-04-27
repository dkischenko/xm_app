package company

import (
	"context"
	"fmt"
	"github.com/dkischenko/xm_app/internal/company/models"
	uerrors "github.com/dkischenko/xm_app/internal/errors"
	"github.com/dkischenko/xm_app/pkg/logger"
)

type Service struct {
	logger  *logger.Logger
	storage Repository
}

type IService interface {
	CreateCountry(ctx context.Context, company models.CompanyCreateRequest) (id int, err error)
	CreateCompany(ctx context.Context, company models.CompanyCreateRequest, countryId int) (id int, err error)
	UpdateCompany(ctx context.Context, companyId int, company *models.CompanyUpdateRequest) (err error)
	DeleteCompany(ctx context.Context, companyId int) (err error)
	GetCompanies(ctx context.Context) (companies []models.Company, err error)
	GetCompany(ctx context.Context, companyId int) (company models.Company, err error)
}

func NewService(logger *logger.Logger, storage Repository) IService {
	return &Service{
		logger:  logger,
		storage: storage,
	}
}

func (s Service) CreateCountry(ctx context.Context, company models.CompanyCreateRequest) (id int, err error) {
	id, err = s.storage.CreateCountry(ctx, company)
	if err != nil {
		s.logger.Entry.Errorf("failed to create country: %s", err)
		return 0, fmt.Errorf("error occurs: %w", uerrors.ErrCreateCountry)
	}
	return
}

func (s Service) CreateCompany(ctx context.Context, company models.CompanyCreateRequest, countryId int) (id int, err error) {
	id, err = s.storage.Create(ctx, company, countryId)
	if err != nil {
		s.logger.Entry.Errorf("failed to create company: %s", err)
		return 0, fmt.Errorf("error occurs: %w", uerrors.ErrCreateCompany)
	}
	return
}

func (s Service) UpdateCompany(ctx context.Context, companyId int, company *models.CompanyUpdateRequest) (err error) {
	err = s.storage.Update(ctx, companyId, company)
	if err != nil {
		s.logger.Entry.Errorf("failed to update company: %s", err)
		return fmt.Errorf("error occurs: %w", uerrors.ErrUpdateCompany)
	}
	return
}

func (s Service) DeleteCompany(ctx context.Context, companyId int) (err error) {
	err = s.storage.Delete(ctx, companyId)
	if err != nil {
		s.logger.Entry.Errorf("failed to create country: %s", err)
		return fmt.Errorf("error occurs: %w", uerrors.ErrDeleteCompany)
	}
	return
}

func (s Service) GetCompany(ctx context.Context, cId int) (company models.Company, err error) {
	company, err = s.storage.GetCompany(ctx, cId)
	if err != nil {
		s.logger.Entry.Errorf("failed to get companies: %s", err)
		return company, fmt.Errorf("error occurs: %w", uerrors.ErrGetCompany)
	}
	return
}

func (s Service) GetCompanies(ctx context.Context) (companies []models.Company, err error) {
	companies, err = s.storage.GetList(ctx)
	if err != nil {
		s.logger.Entry.Errorf("failed to get companies: %s", err)
		return companies, fmt.Errorf("error occurs: %w", uerrors.ErrGetCompanies)
	}
	return
}
