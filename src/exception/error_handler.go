package exception

import (
	"backend-golang-evermos/src/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ErrorHandler is function for handle error validation or type Error()
func ErrorHandler(c *gin.Context, recovered interface{}) {

	if err, ok := recovered.(validator.ValidationErrors); ok {
		if ok {
			ValidationErrors(c, err)
			return
		}
	}

	if error, ok := recovered.(error); ok {
		errMessage := map[string]string{
			"message": error.Error(),
		}
		c.JSON(http.StatusInternalServerError, errMessage)
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
	return
}

func ValidationErrors(c *gin.Context, err error) {
	response := helper.ApiResponse("BAD REQUEST", http.StatusBadRequest, "failed", err.Error())
	c.JSON(http.StatusBadRequest, response)
	return
}
