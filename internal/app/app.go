package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/dkischenko/xm_app/internal/config"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
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

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
