package company

import (
	"context"
	"fmt"
	"github.com/dkischenko/xm_app/internal/company/models"
	uerrors "github.com/dkischenko/xm_app/internal/errors"
	"github.com/dkischenko/xm_app/pkg/auth"
	"github.com/dkischenko/xm_app/pkg/hasher"
	"github.com/dkischenko/xm_app/pkg/logger"
	"strings"
	"time"
)

type Service struct {
	logger       *logger.Logger
	storage      Repository
	tokenManager *auth.Manager
}

type IService interface {
	CreateCountry(ctx context.Context, company models.CompanyCreateRequest) (id int, err error)
	CreateCompany(ctx context.Context, company models.CompanyCreateRequest, countryId int) (id int, err error)
	UpdateCompany(ctx context.Context, companyId int, company *models.CompanyUpdateRequest) (err error)
	DeleteCompany(ctx context.Context, companyId int) (err error)
	GetCompanies(ctx context.Context) (companies []models.Company, err error)
	GetCompany(ctx context.Context, companyId int) (company models.Company, err error)
	CreateUser(ctx context.Context, user models.UserRequest) (id string, err error)
	Login(ctx context.Context, ur *models.UserRequest) (u *models.User, err error)
	CreateToken(u *models.User) (hash string, err error)
	CheckAuth(header string) (uuid string, err error)
}

func NewService(logger *logger.Logger, storage Repository, tokenTTL time.Duration) IService {
	tm, err := auth.NewManager(tokenTTL)
	if err != nil {
		logger.Entry.Errorf("error with token manager: %s", err)
	}

	return &Service{
		tokenManager: tm,
		logger:       logger,
		storage:      storage,
	}
}

func (s Service) CheckAuth(header string) (uuid string, err error) {
	headerString := strings.Split(header, " ")
	if len(headerString[1]) == 0 {
		return "", fmt.Errorf("error occurs: %w", uerrors.ErrEmptyToken)
	}
	uuid, err = s.tokenManager.ParseJWT(headerString[1])
	if err != nil {
		return "", fmt.Errorf("error occurs: %w", uerrors.ErrParseToken)
	}
	return uuid, nil
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

func (s Service) CreateUser(ctx context.Context, user models.UserRequest) (id string, err error) {
	hashPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		s.logger.Entry.Errorf("troubles with hashing password: %s", user.Password)
		return "", err
	}
	usr := &models.User{
		Name:         user.Name,
		PasswordHash: hashPassword,
	}

	id, err = s.storage.CreateUser(ctx, usr)

	if err != nil {
		return id, err
	}

	return
}

func (s Service) Login(ctx context.Context, ur *models.UserRequest) (u *models.User, err error) {
	u, err = s.storage.FindOneUser(ctx, ur.Name)
	if err != nil {
		s.logger.Entry.Errorf("failed find user with error: %s", err)
		return nil, fmt.Errorf("error occurs: %w", uerrors.ErrFindOneUser)
	}

	if !hasher.CheckPasswordHash(u.PasswordHash, ur.Password) {
		s.logger.Entry.Errorf("user used wrong password: %s", err)
		return nil, fmt.Errorf("error occurs: %w", uerrors.ErrCheckUserPasswordHash)
	}

	return
}

func (s Service) CreateToken(u *models.User) (hash string, err error) {
	hash, err = s.tokenManager.CreateJWT(u.Id)
	if err != nil {
		s.logger.Entry.Errorf("problems with creating jwt token: %s", err)
		return "", fmt.Errorf("error occurs: %w", uerrors.ErrCreateJWTToken)
	}

	return
}
