package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"xyz/dbettkk/VerificationCodeDemo/route"
)


func main() {
	r := gin.Default()

	route.Register(r)

	if err := r.Run(":8000"); err != nil {
		os.Exit(1)
	}
}

