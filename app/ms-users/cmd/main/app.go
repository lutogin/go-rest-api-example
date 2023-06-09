package main

import (
	"fmt"
	"ms-users/config"
	"ms-users/internal/users"
	"ms-users/internal/users/db"
	mng "ms-users/pkg/db/mongodb"
	"ms-users/pkg/logging"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	clientOpts := mng.MongoConnectOpt{
		User:      cfg.MongoUser,
		Password:  cfg.MongoPassword,
		Host:      cfg.MongoHost,
		Port:      cfg.MongoPort,
		Database:  cfg.MongoDatabase,
		UriScheme: cfg.MongoUriScheme,
	}

	client, err := mng.NewClient(clientOpts, logger)
	if err != nil {
		panic(err)
	}
	userRepo := db.NewRepository(client, "users", logger)
	service := users.NewService(userRepo, logger)
	logger.Infoln("Registering routers.")
	router := httprouter.New()
	(users.NewHandler(service, logger)).Register(router)

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
