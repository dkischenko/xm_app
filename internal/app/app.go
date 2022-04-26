package app

import (
	"fmt"
	"github.com/dkischenko/xm_app/internal/config"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"time"
)

func Run(router *mux.Router, logger *logger.Logger, config *config.Config) {
	logger.Entry.Info("start application")
	logger.Entry.Info("listen TCP")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Listen.Ip, config.Listen.Port))

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Entry.Infof("server listening address %s:%s", config.Listen.Ip, config.Listen.Port)
	log.Fatal(server.Serve(listener))
}
