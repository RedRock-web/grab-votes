// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 22:50
package service

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestA(t *testing.T) {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		time.Sleep(time.Second * 10)

		context.String(200, "sdf")
	})

	r.Run()
}
