package company

import (
	"github.com/dkischenko/xm_app/internal/config"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
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
	router.HandleFunc("/v1/companies", CreateCompanyHandler).Methods(http.MethodGet)
}

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
