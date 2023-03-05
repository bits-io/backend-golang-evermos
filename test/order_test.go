package test

import (
	"backend-golang-evermos/src/helper"
	"backend-golang-evermos/src/middleware"
	"backend-golang-evermos/src/service/order"
	"backend-golang-evermos/src/service/product"
	"backend-golang-evermos/src/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

// setupDB is function for setting connection to database for testing
func setupDB() *gorm.DB {
	const DB_USERNAME = "root"
	const DB_PASSWORD = "root"
	const DB_HOST = "localhost"
	const DB_PORT = "3306"
	const DB_NAME = "evermos"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&product.ProductPromotion{})
	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&order.OrderHistory{})
	db.AutoMigrate(&order.OrderProduct{})
	return db
}

// setupRouter is function to set up router for testing
func setupRouter(db *gorm.DB) (*gin.Engine, user.User, product.Product) {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository, userService, productService)
	orderControler := order.NewController(orderService)

	DeleteAllData(orderService, productService, userService)
	user, product := InsertDataUserAndProduct(productService, userService)

	router := gin.Default()

	orderControler.Route(router, middleware.DBTransactionMiddleware(db))

	return router, user, product
}

// TestOrderCreate is function to test order create success and failed
func TestOrderCreate(t *testing.T) {
	db := setupDB()

	router, user, product := setupRouter(db)
	t.Run("CreateOrderSuccess", func(t *testing.T) {
		for i := 1; i <= 3; i++ {
			CreateOrderSuccess(t, router, i, user, product)
		}
	})

	t.Run("CreateOrderFailed", func(t *testing.T) {
		for i := 1; i <= 3; i++ {
			CreateOrderFailed(t, router, i, user, product)
		}
	})

}

// CreateOrderSuccess is function for test create order success with http test and the result must success
func CreateOrderSuccess(t *testing.T, router *gin.Engine, number int, userData user.User, productData product.Product) {
	fmt.Println("-------------")
	fmt.Println("-------------")
	fmt.Println("Order - ", strconv.Itoa(number))

	w := httptest.NewRecorder()
	contentJson := fmt.Sprintf(`{"user_id": %s,"product_id": %s,"quantity": 1}`, strconv.Itoa(userData.Id), strconv.Itoa(productData.Id))
	requestBody := strings.NewReader(contentJson)
	req, _ := http.NewRequest(http.MethodPost, "/api/orders/", requestBody)
	router.ServeHTTP(w, req)

	response := w.Result()
	assert.Equal(t, 200, response.StatusCode)

}

// CreateOrderFailed is function for test create order failed with http test and the result must failed
func CreateOrderFailed(t *testing.T, router *gin.Engine, number int, userData user.User, productData product.Product) {
	fmt.Println("-------------")
	fmt.Println("-------------")
	fmt.Println("Order - ", strconv.Itoa(number))

	w := httptest.NewRecorder()
	contentJson := fmt.Sprintf(`{"user_id": %s,"product_id": %s,"quantity": 1}`, strconv.Itoa(userData.Id), strconv.Itoa(productData.Id))
	requestBody := strings.NewReader(contentJson)
	req, _ := http.NewRequest(http.MethodPost, "/api/orders/", requestBody)
	router.ServeHTTP(w, req)

	response := w.Result()
	assert.Equal(t, 400, response.StatusCode)

}

// DeleteAllData is function for delete any data product, order, and user
func DeleteAllData(orderService order.Service, productService product.Service, userService user.Service) {
	orders := orderService.FindAll()
	for _, orderData := range orders {
		orderService.Delete(orderData)
	}

	products := productService.FindAll()
	for _, productData := range products {
		productService.Delete(productData)
	}

	users := userService.FindAll()
	for _, userData := range users {
		userService.Delete(userData)
	}
}

// InsertDataUserAndProduct is function for insert data user and product for testing
func InsertDataUserAndProduct(productService product.Service, userService user.Service) (user.User, product.Product) {
	user := user.User{
		Name:      "Example User",
		Email:     "example.user@mailinator.com",
		Password:  "password",
		Phone:     "08172318926737",
		Address:   "jalan example",
		City:      "Bandung",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user = userService.Save(user)

	input := product.CreateProductRequest{
		Name:               "product another",
		Description:        "description",
		Ingredient:         []string{"gula", "batu"},
		Price:              5000,
		Type:               product.TYPE_FLASH_SALE,
		Image:              "image",
		Stock:              10,
		StockPromotion:     3,
		PriceAfterDiscount: 3000,
	}
	productData := productService.Save(input)
	productData, _ = productService.FindById(productData.Id)

	return user, productData
}
