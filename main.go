package main

import (
	"convert2gltF/handle"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/path", handle.Convert)

	r.Run(":2022") //监听端口默认为8080
}
