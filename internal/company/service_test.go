package company_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/dkischenko/xm_app/internal/company"
	mock_company "github.com/dkischenko/xm_app/internal/company/mocks"
	"github.com/dkischenko/xm_app/internal/company/models"
	uerrors "github.com/dkischenko/xm_app/internal/errors"
	"github.com/dkischenko/xm_app/pkg/auth"
	"github.com/dkischenko/xm_app/pkg/hasher"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewService(t *testing.T) {
	l, _ := logger.GetLogger()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_company.NewMockRepository(ctrl)
	assert.NotNil(t, company.NewService(l, mockRepo, 3600))
}

func TestService_CreateCountry(t *testing.T) {
	t.Run("Create country", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		cmp := models.CompanyCreateRequest{
			Name:    "test",
			Code:    12345,
			Country: "Ukr",
			Website: "https://example.com",
			Phone:   "+380662342437",
		}
		mockRepo.EXPECT().CreateCountry(context.Background(), cmp).Return(1, nil).AnyTimes()
		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		id, err := s.CreateCountry(context.Background(), cmp)
		if err != nil {
			t.Fatalf("Cannot store user via service due error: %s", err)
		}
		assert.NotNil(t, id, "Country id can't be nil")
	})
}

func TestService_CreateCountryErr(t *testing.T) {
	t.Run("Create country Err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		cmp := models.CompanyCreateRequest{
			Name:    "test",
			Code:    12345,
			Country: "Ukr",
			Website: "https://example.com",
			Phone:   "+380662342437",
		}
		mockRepo.EXPECT().CreateCountry(context.Background(), cmp).
			Return(0, fmt.Errorf("Error occurs: %w", uerrors.ErrCreateCountry)).AnyTimes()
		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		_, err := s.CreateCountry(context.Background(), cmp)
		if err != nil {
			assert.ErrorIs(t, err, uerrors.ErrCreateCountry)
		} else {
			t.Fatalf("Unexpected error: %s", err)
		}
	})
}

func TestService_Login(t *testing.T) {
	t.Run("User login(Ok)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.Background()
		l, _ := logger.GetLogger()
		hash, _ := hasher.HashPassword("password")
		uId := "c9f44c4a-788a-4d5f-a210-94ccafcc2231"
		mockRepo := mock_company.NewMockRepository(ctrl)
		ur := &models.UserRequest{
			Name:     "Bob",
			Password: "password",
		}
		mockRepo.EXPECT().
			FindOneUser(ctx, ur.Name).Return(&models.User{
			Id:           uId,
			Name:         ur.Name,
			PasswordHash: hash,
		}, nil).AnyTimes()
		service := company.NewService(l, mockRepo, 3600)
		u, err := mockRepo.FindOneUser(ctx, ur.Name)
		if err != nil {
			t.Fatalf("Can't find user with credentials due error: %s", err)
		}

		if !hasher.CheckPasswordHash(u.PasswordHash, ur.Password) {
			t.Fatalf("User with wrong password. Error: %s", err)
		}

		usr, err := service.Login(ctx, ur)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		assert.NotNil(t, usr)
	})
}

func TestService_LoginFindOneError(t *testing.T) {
	t.Run("User login find one error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		ctrl.Finish()

		ctx := context.Background()
		mockRepo := mock_company.NewMockRepository(ctrl)
		mockRepo.EXPECT().
			FindOneUser(ctx, "Bob").
			Return(nil, fmt.Errorf("Error occurs: %w", uerrors.ErrFindOneUser)).AnyTimes()

		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600)
		ur := &models.UserRequest{
			Name:     "Bob",
			Password: "password",
		}
		_, err := s.Login(ctx, ur)
		if err != nil {
			assert.ErrorIs(t, err, uerrors.ErrFindOneUser)
		} else {
			t.Fatalf("Unexpected error.")
		}
	})
}

func TestService_CheckAuth(t *testing.T) {
	hash, _ := hasher.HashPassword("password")
	t.Run("[Ok] Check Auth", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ctx := context.Background()
		//httpStatusCode := 200
		usr := &models.User{
			Id:           "c9f44c4a-788a-4d5f-a210-94ccafcc2231",
			Name:         "Bob",
			PasswordHash: hash,
		}

		l, _ := logger.GetLogger()
		mockRepo := mock_company.NewMockRepository(ctrl)
		mockRepo.EXPECT().FindOneUser(ctx, usr.Id).
			Return(usr, nil).AnyTimes()

		authRepo, _ := auth.NewManager(3600 * time.Second)
		s := company.NewService(l, mockRepo, 3600*time.Second)
		token, _ := authRepo.CreateJWT(usr.Id)

		//_, err := authRepo.ParseJWT(token)
		header := "Bearer " + token
		uId, err := s.CheckAuth(header)
		assert.NotNil(t, uId)
		if err != nil {
			t.Fatalf("unexpected error")
		}
	})
}

func TestService_CreateCompany(t *testing.T) {
	t.Run("Create company", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		cmp := models.CompanyCreateRequest{
			Name:    "test",
			Code:    12345,
			Country: "Ukr",
			Website: "https://example.com",
			Phone:   "+380662342437",
		}
		mockRepo.EXPECT().Create(context.Background(), cmp, 1).Return(1, nil).AnyTimes()
		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		id, err := s.CreateCompany(context.Background(), cmp, 1)
		if err != nil {
			t.Fatalf("Cannot store company via service due error: %s", err)
		}
		assert.NotNil(t, id, "Company id can't be nil")
	})
}

func TestService_CreateCompanyErr(t *testing.T) {
	t.Run("Create company Err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		cmp := models.CompanyCreateRequest{
			Name:    "test",
			Code:    12345,
			Country: "Ukr",
			Website: "https://example.com",
			Phone:   "+380662342437",
		}
		mockRepo.EXPECT().Create(context.Background(), cmp, 1).Return(0,
			fmt.Errorf("Error occurs: %w", uerrors.ErrCreateCompany)).AnyTimes()
		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		_, err := s.CreateCompany(context.Background(), cmp, 1)
		if err != nil {
			assert.ErrorIs(t, err, uerrors.ErrCreateCompany)
		} else {
			t.Fatalf("Unexpected error: %s", err)
		}
	})
}

func TestService_UpdateCompany(t *testing.T) {
	t.Run("Update company", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		cmp := &models.CompanyUpdateRequest{
			Name:      "test",
			Code:      12345,
			CountryId: 1,
			Website:   "https://example.com",
			Phone:     "+380662342437",
		}

		mockRepo.EXPECT().Update(context.Background(), 1, cmp).Return(nil)
		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		err := s.UpdateCompany(context.Background(), 1, cmp)
		if err != nil {
			t.Fatalf("Cannot update company via service due error: %s", err)
		}
	})
}

func TestService_UpdateCompanyErr(t *testing.T) {
	t.Run("Update company err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		cmp := &models.CompanyUpdateRequest{
			Name:      "test",
			Code:      12345,
			CountryId: 1,
			Website:   "https://example.com",
			Phone:     "+380662342437",
		}

		mockRepo.EXPECT().Update(context.Background(), 1, cmp).
			Return(fmt.Errorf("Error occurs: %w", uerrors.ErrUpdateCompany))
		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		err := s.UpdateCompany(context.Background(), 1, cmp)
		if err != nil {
			assert.ErrorIs(t, err, uerrors.ErrUpdateCompany)
		} else {
			t.Fatalf("Unexpected error: %s", err)
		}
	})
}

func TestService_DeleteCompany(t *testing.T) {
	t.Run("Delete company", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := mock_company.NewMockRepository(ctrl)
		mockRepo.EXPECT().Delete(context.Background(), 1).Return(nil).AnyTimes()

		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)

		err := s.DeleteCompany(context.Background(), 1)
		if err != nil {
			t.Fatalf("Cannot delete company via service due error: %s", err)
		}
	})
}

func TestService_DeleteCompanyErr(t *testing.T) {
	t.Run("Delete company err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := mock_company.NewMockRepository(ctrl)
		mockRepo.EXPECT().Delete(context.Background(), 1).
			Return(fmt.Errorf("Error occurs: %w", uerrors.ErrDeleteCompany)).AnyTimes()

		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)

		err := s.DeleteCompany(context.Background(), 1)
		if err != nil {
			assert.ErrorIs(t, err, uerrors.ErrDeleteCompany)
		} else {
			t.Fatalf("Unexpected error: %s", err)
		}
	})
}

func TestService_GetCompany(t *testing.T) {
	t.Run("Get company", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		mockRepo.EXPECT().GetCompany(context.Background(), 1).Return(models.Company{
			Id:        1,
			Name:      "Test",
			Code:      1231432,
			CountryId: 1,
			Website:   "https://example.com",
			Phone:     "+3806678934556",
			CreatedAt: 1650995663,
			UpdatedAt: 1651002057,
		}, nil).AnyTimes()

		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		cmp, err := s.GetCompany(context.Background(), 1)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		assert.NotNil(t, cmp)
	})
}

func TestService_GetCompanyErr(t *testing.T) {
	t.Run("Get company Err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		mockRepo.EXPECT().GetCompany(context.Background(), 1).
			Return(models.Company{}, fmt.Errorf("Error occurs: %w", uerrors.ErrGetCompany)).AnyTimes()

		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		_, err := s.GetCompany(context.Background(), 1)
		if err != nil {
			assert.ErrorIs(t, err, uerrors.ErrGetCompany)
		} else {
			t.Fatalf("Unexpected error: %s", err)
		}
	})
}

func TestService_GetCompanies(t *testing.T) {
	t.Run("Get companies", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		companies := make([]models.Company, 2)
		companies[0] = models.Company{
			Id:        1,
			Name:      "Test",
			Code:      1231432,
			CountryId: 1,
			Website:   "https://example.com",
			Phone:     "+3806678934556",
			CreatedAt: 1650995663,
			UpdatedAt: 1651002057,
		}
		companies[1] = models.Company{
			Id:        2,
			Name:      "Very test",
			Code:      1231432,
			CountryId: 1,
			Website:   "https://example.com",
			Phone:     "+3806678934556",
			CreatedAt: 1650995663,
			UpdatedAt: 1651002057,
		}
		mockRepo.EXPECT().GetList(context.Background()).Return(companies, nil).AnyTimes()

		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		cmp, err := s.GetCompanies(context.Background())
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		assert.NotNil(t, cmp)
		assert.Equal(t, len(cmp), 2)
	})
}

func TestService_GetCompaniesErr(t *testing.T) {
	t.Run("Get companies err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_company.NewMockRepository(ctrl)
		mockRepo.EXPECT().GetList(context.Background()).
			Return(nil, fmt.Errorf("Error occurs: %w", uerrors.ErrGetCompanies)).AnyTimes()

		l, _ := logger.GetLogger()
		s := company.NewService(l, mockRepo, 3600*time.Second)
		_, err := s.GetCompanies(context.Background())
		if err != nil {
			assert.ErrorIs(t, err, uerrors.ErrGetCompanies)
		} else {
			t.Fatalf("Unexpected error: %s", err)
		}
	})
}

func TestService_CreateUser(t *testing.T) {
	testCases := []struct {
		name      string
		ctx       context.Context
		user      *models.UserRequest
		wantError bool
	}{
		{
			name: "OK case",
			ctx:  context.Background(),
			user: &models.UserRequest{
				Name:     "Bill",
				Password: "password",
			},
			wantError: false,
		},
		{
			name: "Empty password (skip)",
			ctx:  context.Background(),
			user: &models.UserRequest{
				Name:     "Bill",
				Password: "",
			},
			wantError: true,
		},
		{
			name: "Empty name",
			ctx:  context.Background(),
			user: &models.UserRequest{
				Name:     "",
				Password: "password",
			},
			wantError: true,
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l, _ := logger.GetLogger()

			mockRepo := mock_company.NewMockRepository(ctrl)
			uId := "c9f44c4a-788a-4d5f-a210-94ccafcc2231"
			mockRepo.EXPECT().
				CreateUser(tcase.ctx, gomock.Any()).Return(uId, nil).AnyTimes()

			service := company.NewService(l, mockRepo, 3600)
			if len(tcase.user.Name) == 0 {
				if tcase.wantError {
					t.Skip("Username can't be empty")
				}
				t.Error("Unexpected error")
			}
			hash, err := hasher.HashPassword(tcase.user.Password)
			if err != nil {
				if tcase.wantError {
					assert.Equal(t, errors.New("String must not be empty"), err)
					t.Skipf("Expected error: %s", err)
				}
				t.Errorf("Unexpected error: %s", err)
			}

			u := &models.User{
				Name:         tcase.user.Name,
				PasswordHash: hash,
			}

			id, err := mockRepo.CreateUser(tcase.ctx, u)
			if err != nil {
				t.Fatalf("Cannot store user due error: %s", err)
			}
			assert.Equal(t, len(id), 36, "Got wrong UUID format")

			usr := *tcase.user
			id, err = service.CreateUser(tcase.ctx, usr)
			if err != nil {
				t.Fatalf("Cannot store user via service due error: %s", err)
			}
			assert.NotNil(t, id, "User id can't be nil")
		})
	}
}

func TestService_CreateToken(t *testing.T) {
	t.Run("Create token", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		l, _ := logger.GetLogger()
		mockRepo := mock_company.NewMockRepository(ctrl)
		service := company.NewService(l, mockRepo, 3600)
		uId := "c9f44c4a-788a-4d5f-a210-94ccafcc2231"
		hash, err := service.CreateToken(uId)

		if err != nil {
			t.Fatalf("unexpected error")
		}

		assert.NotNil(t, hash)
	})
}
