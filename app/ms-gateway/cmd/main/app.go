package main

import (
	"fmt"
	"ms-gateway/config"
	"ms-gateway/internal/user"
	"ms-gateway/pkg/logging"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	cfg := config.GetConfig()

	logger.Infoln("Register routers.")
	router := httprouter.New()
	(user.NewHandler(logger)).Register(router)

	start(router, logger, cfg)
}

func start(router *httprouter.Router, logger *logging.Logger, cfg *config.Config) {
	logger.Infoln("Start application.")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port))

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infoln(fmt.Sprintf("Server is started at %s port.", cfg.Listen.Port))

	panic(server.Serve(listener))
}
