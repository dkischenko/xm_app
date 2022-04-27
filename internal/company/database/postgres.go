package database

import (
	"context"
	"fmt"
	"github.com/dkischenko/xm_app/internal/company"
	"github.com/dkischenko/xm_app/internal/company/models"
	uerrors "github.com/dkischenko/xm_app/internal/errors"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
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

func (p postgres) Create(ctx context.Context, company models.CompanyCreateRequest, countryId int) (id int, err error) {
	c := &models.Company{}
	q := `
		INSERT INTO xm_db.companies(name, code, country_id, website, phone, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err = p.pool.QueryRow(ctx, q, company.Name, company.Code, countryId, company.Website, company.Phone, time.Now().Unix(), time.Now().Unix()).
		Scan(&c.Id)

	if err != nil {
		p.logger.Entry.Error(err)
		return 0, fmt.Errorf("Error occurs: %w. %w", err, uerrors.ErrCreateCompany)
	}

	return c.Id, nil
}

func (p postgres) GetCompany(ctx context.Context, companyId int) (company models.Company, err error) {
	q := `
		SELECT id, name, code, country_id, website, phone, created_at, updated_at
		FROM xm_db.companies
        WHERE id = $1 
	`

	row := p.pool.QueryRow(ctx, q, companyId)
	err = row.Scan(&company.Id, &company.Name, &company.Code, &company.CountryId, &company.Website,
		&company.Phone, &company.CreatedAt, &company.UpdatedAt)
	if err != nil {
		p.logger.Entry.Error(err)
		return
	}

	return
}

func (p postgres) GetList(ctx context.Context) (companies []models.Company, err error) {
	q := `
		SELECT id, name, code, country_id, website, phone, created_at, updated_at
		FROM xm_db.companies
	`
	rows, err := p.pool.Query(ctx, q)
	if err != nil {
		p.logger.Entry.Fatal("error while executing query")
		return nil, fmt.Errorf("Error occurs: %w. %w", err, uerrors.ErrGetCompanies)
	}
	for rows.Next() {
		var r models.Company
		err = rows.Scan(&r.Id, &r.Name, &r.Code, &r.CountryId, &r.Website, &r.Phone, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			p.logger.Entry.Fatalf("Scan: %v", err)
			return nil, fmt.Errorf("Error occurs: %w. %w", err, uerrors.ErrGetCompanies)
		}
		companies = append(companies, r)
	}

	return
}

func (p postgres) Update(ctx context.Context, companyId int, company *models.CompanyUpdateRequest) (err error) {
	q := `
		UPDATE xm_db.companies
		SET name = $1, code = $2, country_id = $3, website = $4, phone = $5, updated_at = $6 
		WHERE id = $7
	`
	_, err = p.pool.Exec(ctx, q, company.Name, company.Code, company.CountryId, company.Website,
		company.Phone, time.Now().Unix(), companyId)
	if err != nil {
		p.logger.Entry.Error(err)
		return fmt.Errorf("Error occurs: %w", uerrors.ErrUpdateCompany)
	}
	return
}

func (p postgres) Delete(ctx context.Context, companyId int) (err error) {
	q := `
		DELETE from xm_db.companies
		WHERE id = $1
	`

	_, err = p.pool.Exec(ctx, q, companyId)
	if err != nil {
		p.logger.Entry.Error(err)
		return fmt.Errorf("Error occurs: %w", uerrors.ErrDeleteCompany)
	}
	return
}

func (p postgres) CreateCountry(ctx context.Context, company models.CompanyCreateRequest) (id int, err error) {
	country := &models.Country{}
	q := `
		SELECT id, name
		FROM xm_db.countries WHERE name = $1
	`

	row := p.pool.QueryRow(ctx, q, company.Country)
	err = row.Scan(&country.Id, &country.Name)
	if err != nil {
		q = `
			INSERT INTO xm_db.countries(name)
			VALUES
				($1)
			RETURNING id
		`

		err = p.pool.QueryRow(ctx, q, company.Country).
			Scan(&country.Id)

		if err != nil {
			p.logger.Entry.Error(err)
			return 0, fmt.Errorf("Error occurs: %w. %w", err, uerrors.ErrCreateCountry)
		}
	}

	return country.Id, nil
}
