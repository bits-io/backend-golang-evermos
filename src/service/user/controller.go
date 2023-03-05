package user

import (
	"backend-golang-evermos/src/helper"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type userController struct {
	service Service
}

func NewController(service Service) *userController {
	return &userController{service: service}
}

// Route is function to define any route user
func (controller *userController) Route(app *gin.Engine) {
	route := app.Group("api/users")
	route.GET("/", controller.Get)
	route.POST("/", controller.Save)
	route.GET("/:id", controller.FindById)
	route.POST("/login/", controller.Login)
	route.POST("/register/", controller.Save)
}

// Get is function to get all data user
func (controller *userController) Get(c *gin.Context) {
	users := controller.service.FindAll()

	c.JSON(http.StatusOK, helper.ApiResponse("List User", http.StatusOK, "success", UsersFormat(users)))
	return
}

// Login is function
func (controller *userController) Login(c *gin.Context) {
	var input CreateUserRequest

	user, err := controller.service.FindByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse(err.Error(), http.StatusBadRequest, "failed", map[string]interface{}{}))
		return
	}


	c.JSON(http.StatusOK, helper.ApiResponse("Login User Successfully", http.StatusOK, "success", UserFormat(user)))
	return
}

// Save is function create data user
func (controller *userController) Save(c *gin.Context) {
	var input CreateUserRequest
	err := c.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)

	user := User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashed),
		Phone:     input.Phone,
		Address:   input.Address,
		City:      input.City,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user = controller.service.Save(user)

	c.JSON(http.StatusOK, helper.ApiResponse("Create User Successfully", http.StatusOK, "success", UserFormat(user)))
	return
}

// FindById is function for get data detail user
func (controller *userController) FindById(c *gin.Context) {
	var inputParam DetailUserRequest
	err := c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	user, err := controller.service.FindById(inputParam.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse(err.Error(), http.StatusBadRequest, "failed", map[string]interface{}{}))
		return
	}

	c.JSON(http.StatusOK, helper.ApiResponse("Detail User", http.StatusOK, "success", UserFormat(user)))
	return
}
