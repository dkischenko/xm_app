package company

import (
	"context"
	"github.com/dkischenko/xm_app/internal/company/models"
)

type Repository interface {
	Create(ctx context.Context, company *models.Company) (id string, err error)
	Find(ctx context.Context, field string, value string) (company models.Company, err error)
	Update(ctx context.Context, company *models.Company) (err error)
	Delete(ctx context.Context, companyId int) (err error)
}
