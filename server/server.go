package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josenymad/one-two-one-two/controller"
)

type HttpServer struct {
	address string
	port    string
}

func NewHttpServer(address, port string) HttpServer {
	return HttpServer{
		address: address,
		port:    port,
	}
}

func (server *HttpServer) Run() {
	router := gin.New()
	router.POST("/test", controller.TestHandler)
	router.Run(fmt.Sprintf("%s:%s", server.address, server.port))
}
