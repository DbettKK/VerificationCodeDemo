package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xyz/dbettkk/VerificationCodeDemo/dal/captcha"
)

func CreateImage(c *gin.Context) {
	imgId := captcha.Instance().CreateImage()
	c.JSON(http.StatusOK,
		gin.H{
			"code": 200,
			"key":  imgId,
			"url":  "/captcha/img/" + imgId,
		})
}

func WriteImage(c *gin.Context) {
	captchaId := c.Param("key")
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Type", "image/png")
	// 重载一次
	captcha.Instance().Reload(captchaId)
	// 输出图片
	_, err := c.Writer.Write(captcha.Instance().GetImageByte(captchaId))
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"code": 500,
				"msg":  "图片输出失败",
			})
	}
}

func VerifyImage(c *gin.Context) {
	captchaId := c.Param("key")
	val := c.Param("val")
	if captcha.Instance().Verify(captchaId, val) {
		c.JSON(http.StatusOK, gin.H{"code": 200})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400})
	}
}
