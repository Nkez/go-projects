package main

import (
	"github.com/julienschmidt/httprouter"
	"golang/basic/internal/user"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Println("create router")
	router := httprouter.New()
	log.Println("register handler")
	handler := user.NewHandler()
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start app")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
	}
	log.Println("server is listening port: 1234")
	log.Fatal(server.Serve(listener))
}
