package product

import (
	"backend-golang-evermos/src/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productController struct {
	service Service
}

func NewController(service Service) *productController {
	return &productController{service: service}
}

// Route is function for define any route product
func (controller *productController) Route(app *gin.Engine) {
	route := app.Group("api/products")
	route.GET("/", controller.Get)
	route.GET("/:id", controller.FindById)
	route.POST("/", controller.Create)
}

// Get is function to get all data product
func (controller *productController) Get(c *gin.Context) {
	products := controller.service.FindAll()

	c.JSON(http.StatusOK, helper.ApiResponse("List Product", http.StatusOK, "success", ProductsFormat(products)))
	return
}

// Create is function to create product
func (controller *productController) Create(c *gin.Context) {
	var input CreateProductRequest
	err := c.ShouldBind(&input)
	helper.PanicIfError(err)

	product := controller.service.Save(input)

	c.JSON(http.StatusOK, helper.ApiResponse("Create Product Successfully", http.StatusOK, "success", ProductFormat(product)))
	return
}

// FindById is function to get detail product
func (controller *productController) FindById(c *gin.Context) {
	var inputParam DetailProductRequest
	err := c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	product, err := controller.service.FindById(inputParam.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse(err.Error(), http.StatusBadRequest, "failed", map[string]interface{}{}))
		return
	}

	c.JSON(http.StatusOK, helper.ApiResponse("Detail Product", http.StatusOK, "success", ProductFormat(product)))
	return
}
