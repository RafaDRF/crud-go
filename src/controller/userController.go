package controller

import (
	"github.com/RafaDRF/crud-go/src/config/rest_err"
	"github.com/RafaDRF/crud-go/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userReq request.UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		restErr := rest_err.NewBadResquestErr(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}
}

func FindUserById(c *gin.Context) {

}

func FindUserByEmail(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
