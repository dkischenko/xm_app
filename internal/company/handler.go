package company

import (
	"encoding/json"
	"fmt"
	"github.com/dkischenko/xm_app/internal/company/models"
	"github.com/dkischenko/xm_app/internal/config"
	uerrors "github.com/dkischenko/xm_app/internal/errors"
	"github.com/dkischenko/xm_app/pkg/ipapi"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	company                = "/v1/companies"
	companyWithId          = "/v1/companies/{id:[0-9]+}"
	headerContentType      = "Content-Type"
	headerValueContentType = "application/json"
)

type handler struct {
	logger  *logger.Logger
	service IService
	config  *config.Config
}

func NewHandler(logger *logger.Logger, service IService, cfg *config.Config) *handler {
	return &handler{
		logger:  logger,
		service: service,
		config:  cfg,
	}
}

func (h handler) Register(router *mux.Router) {
	router.HandleFunc(company, h.GetCompaniesListHandler).Methods(http.MethodGet)
	router.HandleFunc(companyWithId, h.GetCompanyHandler).Methods(http.MethodGet)
	router.HandleFunc(company, h.CreateCompanyHandler).Methods(http.MethodPost)
	router.HandleFunc(companyWithId, h.UpdateCompanyHandler).Methods(http.MethodPut)
	router.HandleFunc(companyWithId, h.DeleteCompanyHandler).Methods(http.MethodDelete)
}

func (h handler) GetCompanyHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cId, _ := strconv.Atoi(params["id"])
	company, err := h.service.GetCompany(r.Context(), cId)

	if err != nil {
		h.logger.Entry.Errorf("can't get company: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add(headerContentType, headerValueContentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(company); err != nil {
		h.logger.Entry.Errorf("can't get companies list: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h handler) GetCompaniesListHandler(w http.ResponseWriter, r *http.Request) {
	companies, err := h.service.GetCompanies(r.Context())
	if err != nil {
		h.logger.Entry.Errorf("can't get companies: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add(headerContentType, headerValueContentType)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(companies); err != nil {
		h.logger.Entry.Errorf("can't get companies list: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h handler) CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	if ipapi.IsAllowed() != true {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	companyData := &models.CompanyCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(companyData)
	if err != nil {
		h.logger.Entry.Error("wrong json format")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	v := validator.New()
	if err := v.Struct(companyData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseBody := uerrors.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("got wrong user data: %+v", err),
		}
		if err := json.NewEncoder(w).Encode(responseBody); err != nil {
			h.logger.Entry.Errorf("problems with encoding data: %+v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		h.logger.Entry.Errorf("got wrong user data: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	countryId, err := h.service.CreateCountry(r.Context(), *companyData)
	if err != nil {
		h.logger.Entry.Errorf("can't create country: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	companyId, err := h.service.CreateCompany(r.Context(), *companyData, countryId)
	if err != nil {
		h.logger.Entry.Errorf("can't create company: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add(headerContentType, headerValueContentType)
	w.WriteHeader(http.StatusOK)
	responseBody := CompanyCreateResponse{
		Id:   companyId,
		Name: companyData.Name,
	}

	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		h.logger.Entry.Errorf("can't create user: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h handler) UpdateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cId, _ := strconv.Atoi(params["id"])

	companyData := &models.CompanyUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(companyData)
	if err != nil {
		h.logger.Entry.Error("wrong json format")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = h.service.UpdateCompany(r.Context(), cId, companyData)
	if err != nil {
		h.logger.Entry.Errorf("can't update company: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add(headerContentType, headerValueContentType)
	w.WriteHeader(http.StatusOK)
}

func (h handler) DeleteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	if ipapi.IsAllowed() != true {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	params := mux.Vars(r)
	cId, _ := strconv.Atoi(params["id"])

	err := h.service.DeleteCompany(r.Context(), cId)

	if err != nil {
		h.logger.Entry.Errorf("can't delete company: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add(headerContentType, headerValueContentType)
	w.WriteHeader(http.StatusOK)
}
