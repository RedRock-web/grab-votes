// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 22:23
package controller

import (
	"github.com/gin-gonic/gin"

	"grab-votes/router"
	"grab-votes/service"
)

func GrabOrder(c *gin.Context) {
	var o service.Order

	if err := c.ShouldBindJSON(&o); err != nil {
		return
	}

	service.GrabOrder(o)
	router.OkWithData(c, "请等待。。")

	// todo:增加响应
}
