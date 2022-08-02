package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"halo": "user"})
}

func Croot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "jidan"})
}
