package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEventsPageContent(c *gin.Context) {
	c.HTML(http.StatusOK, "events.html", gin.H{})
}

func GetAccountPageContent(c *gin.Context) {
	c.HTML(http.StatusOK, "account.html", gin.H{})
}

func GetLoginPageContent(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
