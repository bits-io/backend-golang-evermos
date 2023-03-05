package main

import (
	"backend-golang-evermos/src/app"
	"backend-golang-evermos/src/exception"
	"backend-golang-evermos/src/helper"
	"backend-golang-evermos/src/middleware"
	"backend-golang-evermos/src/service/order"
	"backend-golang-evermos/src/service/product"
	"backend-golang-evermos/src/service/user"
	"github.com/gin-gonic/gin"
)

func main() {

	config := app.New()
	db := app.NewDB(config)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository, userService, productService)
	orderControler := order.NewController(orderService)

	// Setup Gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Setup Route
	userController.Route(router)
	productController.Route(router)
	orderControler.Route(router, middleware.DBTransactionMiddleware(db))

	// Start App
	err := router.Run(":8000")
	helper.PanicIfError(err)

}
