package main

import (
	"context"
	"fmt"
	"ms-gateway/config"
	"ms-gateway/internal/user"
	"ms-gateway/internal/user/db"
	"ms-gateway/pkg/db/mongodb"
	"ms-gateway/pkg/logging"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	clientOpts := mng.MongoConnectOpt{
		Ctx:      context.Background(),
		User:     cfg.MongoUser,
		Password: cfg.MongoPassword,
		Host:     cfg.MongoHost,
		Port:     cfg.MongoPort,
		Database: cfg.MongoDatabase,
	}

	client, err := mng.NewClient(clientOpts)
	if err != nil {
		panic(err)
	}
	db.NewRepository(client, "users", logger)
	logger.Infoln("Register routers.")
	router := httprouter.New()
	(user.NewHandler(logger)).Register(router)

	start(router, logger, cfg)
}

func start(router *httprouter.Router, logger *logging.Logger, cfg *config.ConfigEnv) {
	logger.Infoln("Start application.")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Infoln(fmt.Sprintf("Server is started at %s port.", cfg.Port))

	panic(server.Serve(listener))
}
