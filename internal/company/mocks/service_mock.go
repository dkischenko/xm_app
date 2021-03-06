// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_company is a generated GoMock package.
package mock_company

import (
	context "context"
	reflect "reflect"

	models "github.com/dkischenko/xm_app/internal/company/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// CheckAuth mocks base method.
func (m *MockIService) CheckAuth(header string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuth", header)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAuth indicates an expected call of CheckAuth.
func (mr *MockIServiceMockRecorder) CheckAuth(header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuth", reflect.TypeOf((*MockIService)(nil).CheckAuth), header)
}

// CreateCompany mocks base method.
func (m *MockIService) CreateCompany(ctx context.Context, company models.CompanyCreateRequest, countryId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, company, countryId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockIServiceMockRecorder) CreateCompany(ctx, company, countryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockIService)(nil).CreateCompany), ctx, company, countryId)
}

// CreateCountry mocks base method.
func (m *MockIService) CreateCountry(ctx context.Context, company models.CompanyCreateRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCountry", ctx, company)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCountry indicates an expected call of CreateCountry.
func (mr *MockIServiceMockRecorder) CreateCountry(ctx, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCountry", reflect.TypeOf((*MockIService)(nil).CreateCountry), ctx, company)
}

// CreateToken mocks base method.
func (m *MockIService) CreateToken(uId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", uId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockIServiceMockRecorder) CreateToken(uId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockIService)(nil).CreateToken), uId)
}

// CreateUser mocks base method.
func (m *MockIService) CreateUser(ctx context.Context, user models.UserRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIServiceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIService)(nil).CreateUser), ctx, user)
}

// DeleteCompany mocks base method.
func (m *MockIService) DeleteCompany(ctx context.Context, companyId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCompany", ctx, companyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCompany indicates an expected call of DeleteCompany.
func (mr *MockIServiceMockRecorder) DeleteCompany(ctx, companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCompany", reflect.TypeOf((*MockIService)(nil).DeleteCompany), ctx, companyId)
}

// GetCompanies mocks base method.
func (m *MockIService) GetCompanies(ctx context.Context) ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanies", ctx)
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanies indicates an expected call of GetCompanies.
func (mr *MockIServiceMockRecorder) GetCompanies(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanies", reflect.TypeOf((*MockIService)(nil).GetCompanies), ctx)
}

// GetCompany mocks base method.
func (m *MockIService) GetCompany(ctx context.Context, companyId int) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", ctx, companyId)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockIServiceMockRecorder) GetCompany(ctx, companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockIService)(nil).GetCompany), ctx, companyId)
}

// Login mocks base method.
func (m *MockIService) Login(ctx context.Context, ur *models.UserRequest) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, ur)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIServiceMockRecorder) Login(ctx, ur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIService)(nil).Login), ctx, ur)
}

// UpdateCompany mocks base method.
func (m *MockIService) UpdateCompany(ctx context.Context, companyId int, company *models.CompanyUpdateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCompany", ctx, companyId, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCompany indicates an expected call of UpdateCompany.
func (mr *MockIServiceMockRecorder) UpdateCompany(ctx, companyId, company interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCompany", reflect.TypeOf((*MockIService)(nil).UpdateCompany), ctx, companyId, company)
}
