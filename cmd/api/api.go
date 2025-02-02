package api

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-otp/internal/middlewares"
	"github.com/Kei-K23/go-otp/internal/services/auth"
	"github.com/Kei-K23/go-otp/internal/services/todo"
	"github.com/Kei-K23/go-otp/internal/services/user"
	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Addr string
	DB   *sql.DB
}

func (apiServer *APIServer) Serve() {
	r := gin.New()

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning.
	r.SetTrustedProxies(nil)
	// middleware register here
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// route group register here
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// protected routes
	protected := v1.Group("")

	// services register here
	authService := auth.NewStore(apiServer.DB)
	todoService := todo.NewStore(apiServer.DB)
	userService := user.NewStore(apiServer.DB)

	// handlers register here
	authHandler := auth.NewHandler(authService, userService)
	todoHandler := todo.NewHandler(todoService)
	userHandler := user.NewHandler(userService, todoService)

	// register routes here
	v1.Use(middlewares.CheckCookieExist)
	authHandler.RegisterRoutes(*v1)

	// add auth middleware
	protected.Use(middlewares.AuthMiddleware)
	todoHandler.RegisterRoutes(*protected)
	userHandler.RegisterRoutes(*protected)

	r.Run(apiServer.Addr)
	fmt.Printf("server is running on http://localhost%s", apiServer.Addr)
}

func NewAPIServer(apiServer APIServer) *APIServer {
	return &APIServer{Addr: apiServer.Addr, DB: apiServer.DB}
}
