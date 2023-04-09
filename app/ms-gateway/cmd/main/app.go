package main

import (
	"log"
	"ms-gateway/internal/user"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Register routers.")
	router := httprouter.New()
	(user.New()).Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("Start application.")

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Start server at 8080 port.")

	panic(server.Serve(listener))
}
