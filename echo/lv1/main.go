package main

import (
	"lv1/handler"
	"lv1/mdw"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	server := echo.New()

	// Middleware
	// server.Use(middleware.Logger())
	// server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	// }))

	isLogged := middleware.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdw

	// Routes
	server.GET("/", handler.Hello, isLogged)
	server.GET("/admin", handler.HelloAdmin, isLogged, isAdmin)
	server.GET("/show", handler.Show)
	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	// /v2/hello
	groupV2 := server.Group("/api", isLogged /*,.... */)
	groupV2.GET("/users", handler.GetUsers)
	groupV2.GET("/users/:id", handler.GetUser)
	groupV2.POST("/users/create", handler.CreateUser, isAdmin)
	groupV2.PUT("/users/edit/:id", handler.UpdateUser, isAdmin)
	groupV2.DELETE("/users/delete/:id", handler.DeleteUser, isAdmin)

	// Start server
	server.Logger.Fatal(server.Start(":1323"))
}
