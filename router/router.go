// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 16:33
package router

import (
	"github.com/gin-gonic/gin"

	"grab-votes/controller"
)

// InitRouter 初始化 router
func InitRouter() {
	r := gin.Default()
	SetRouter(r)
	_ = r.Run("8080")
}

// SetRouter 设置 router
func SetRouter(r *gin.Engine) {
	r.POST("/grab", controller.GrabOrder)
}
