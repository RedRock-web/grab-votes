// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 16:22
package main

import (
	"grab-votes/model"
	"grab-votes/router"
	"grab-votes/service"
)

func main() {
	model.InitDB()
	service.InitService()
	router.InitRouter()
}
