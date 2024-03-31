package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Addr string
	DB   *sql.DB
}

func (apiServer *APIServer) Serve() {
	r := gin.New()

	// middleware register here
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// route group register here
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// protected routes
	// protected := v1.Group("")

	r.Run(":9999")
	fmt.Println("server is running on http://localhost:9999")
}

func NewAPIServer(apiServer APIServer) *APIServer {
	return &APIServer{Addr: apiServer.Addr, DB: apiServer.DB}
}