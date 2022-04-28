package company_test

import (
	"context"
	"github.com/dkischenko/xm_app/internal/company"
	mock_company "github.com/dkischenko/xm_app/internal/company/mocks"
	"github.com/dkischenko/xm_app/internal/company/models"
	"github.com/dkischenko/xm_app/internal/config"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_RegisterOk(t *testing.T) {
	t.Run("[Ok] Register handlers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			ctx  = context.Background()
			uDTO = models.UserRequest{
				Name:     "bill",
				Password: "password",
			}
			getUUID = "c9f44c4a-788a-4d5f-a210-94ccafcc2231"
			cfg     = &config.Config{}
			payload = `
				{
					"name": "bill",
					"password": "password"
				}`
		)

		req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(payload))
		w := httptest.NewRecorder()
		l, _ := logger.GetLogger()
		mockService := mock_company.NewMockIService(ctrl)
		mockService.EXPECT().CreateUser(ctx, uDTO).Return(getUUID, nil).AnyTimes()
		h := company.NewHandler(l, mockService, cfg)
		router := mux.NewRouter()
		h.Register(router)
		h.CreateUser(w, req)
		assert.Equal(t, w.Code, http.StatusOK)
	})
}
