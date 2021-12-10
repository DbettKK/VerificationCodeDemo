package route

import (
	"github.com/gin-gonic/gin"
	"xyz/dbettkk/VerificationCodeDemo/handlers"
)

// Register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *gin.Engine) {
	customizeRegister(r)

	r.GET("/captcha/create", handlers.CreateImage)
	r.GET("/captcha/img/:key", handlers.WriteImage)
	r.GET("/captcha/verify/:key/:val", handlers.VerifyImage)
}
