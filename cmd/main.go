package main

import (
	"MarketPlace_Pet/internal/db"
	"MarketPlace_Pet/internal/handlers"
	"MarketPlace_Pet/internal/userService"
	"MarketPlace_Pet/internal/warehouseService"
	"MarketPlace_Pet/internal/web/users"
	"MarketPlace_Pet/internal/web/warehouse"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	echoServer := echo.New()

	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Logger())

	wrhRepo := warehouseService.NewWarehouseRepository(database)
	wrhService := warehouseService.NewWarehouseService(wrhRepo)
	wrhHandler := handlers.NewWarehouseHandler(wrhService)

	usrRepo := userService.NewUserRepository(database)
	usrService := userService.NewUserService(usrRepo, wrhRepo)
	usrHandler := handlers.NewUserHandler(usrService)

	warehouseStrictHandler := warehouse.NewStrictHandler(wrhHandler, nil)
	warehouse.RegisterHandlers(echoServer, warehouseStrictHandler)

	userStrictHandler := users.NewStrictHandler(usrHandler, nil)
	users.RegisterHandlers(echoServer, userStrictHandler)

	err = echoServer.Start("localhost:8080")
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
