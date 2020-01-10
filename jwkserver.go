package main

import (
	"github.com/gin-gonic/gin"
	api "jwkserver/app/api/jwk"
)

func main() {
	r := gin.Default()
	rsa := r.Group("/rsa")
	rsa.GET("/jwkeygen",api.Genjwk)
	rsa.POST("/getpvk",api.Getpvk)
	rsa.POST("/getpuk",api.Getpuk)
	r.Run(":8080")
}