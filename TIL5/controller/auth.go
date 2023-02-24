package controller

import (
	"domain"
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func authRouter(auth *gin.RouterGroup) {
	auth.POST("/signup", signup)
	auth.POST("/login", signin)
}

func signup(c *gin.Context) {
	user := domain.Account{}

	if err := c.Bind(&user); err != nil {
		log.Printf("[controller:user] error signup : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	_, _, err := service.JoinUser(user)

	if err != nil {
		log.Printf("[controller:user] error signup : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.Status(http.StatusOK)
}

func signin(c *gin.Context) {
	user := domain.Account{}

	if err := c.Bind(&user); err != nil {
		log.Printf("[controller:user] error signup : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	status, err := service.FindUserByIDAndPassword(user)

	if err != nil {
		log.Printf("[controller:user] error signup : %v\n", err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, domain.StatusContainer{Status: status})
}
