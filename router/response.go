// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 20:53
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 10000})
}

func FormError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 10001, "message": "request form error!"})
}

func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 10000, "message": "ok", "data": data})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
}
