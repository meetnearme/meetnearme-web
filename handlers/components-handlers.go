package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLoginFormComponent(c *gin.Context) {
	c.HTML(http.StatusOK, "LoginForm.html", gin.H{})
}
