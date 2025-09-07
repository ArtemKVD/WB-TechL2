package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", showLoginPage)
	router.GET("/calendar", showCalendar)

	router.Run(":8080")
}

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Error": "",
	})
}

func showCalendar(c *gin.Context) {
	userID := c.Query("userid")

	if userID == "" {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Error": "insert USER ID",
		})
		return
	}

	c.HTML(http.StatusOK, "calendar.html", gin.H{
		"UserID":  userID,
		"Message": "Calendar user" + userID,
	})
}
