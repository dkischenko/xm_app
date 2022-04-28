package company

import (
	"context"
	"github.com/dkischenko/xm_app/internal/company/models"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go
type Repository interface {
	Create(ctx context.Context, company models.CompanyCreateRequest, countryId int) (id int, err error)
	GetList(ctx context.Context) (companies []models.Company, err error)
	GetCompany(ctx context.Context, companyId int) (company models.Company, err error)
	Update(ctx context.Context, companyId int, company *models.CompanyUpdateRequest) (err error)
	Delete(ctx context.Context, companyId int) (err error)
	CreateCountry(ctx context.Context, company models.CompanyCreateRequest) (id int, err error)
	CreateUser(ctx context.Context, user *models.User) (id string, err error)
	FindOneUser(ctx context.Context, name string) (u *models.User, err error)
}
