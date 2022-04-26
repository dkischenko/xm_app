package company

import "github.com/dkischenko/xm_app/pkg/logger"

type Service struct {
	logger  *logger.Logger
	storage Repository
}

type IService interface {
}

func NewService(logger *logger.Logger, storage Repository) IService {
	return &Service{
		logger:  logger,
		storage: storage,
	}
}
