package main

import (
	"ms-gateway/internal/user"
	"ms-gateway/package/logging"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	logger.Infoln("Register routers.")
	router := httprouter.New()
	(user.New()).Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Infoln("Start application.")

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infoln("Server is started at 8080 port.")

	panic(server.Serve(listener))
}
