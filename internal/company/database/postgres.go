package database

import (
	"context"
	"github.com/dkischenko/xm_app/internal/company"
	"github.com/dkischenko/xm_app/internal/company/models"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	logger *logger.Logger
	pool   *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool, logger *logger.Logger) company.Repository {
	return &postgres{
		logger: logger,
		pool:   pool,
	}
}

func (p postgres) Create(ctx context.Context, company *models.Company) (id string, err error) {
	//TODO implement me
	panic("implement me")
}

func (p postgres) Find(ctx context.Context, field string, value string) (company models.Company, err error) {
	//TODO implement me
	panic("implement me")
}

func (p postgres) Update(ctx context.Context, company *models.Company) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p postgres) Delete(ctx context.Context, companyId int) (err error) {
	//TODO implement me
	panic("implement me")
}
