package router

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	channel2 "modules/channel"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		method := c.Request.Method
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	defer func() {
		if x := recover(); x != nil {
			fmt.Print(x)
		}
	}()

	router := gin.Default()
	pprof.Register(router)
	router.Use(Cors())
	//_ = http.ListenAndServeTLS(":443", "cert/www.domain.cn.crt", "cert/www.domain.cn.key", router)

	router.GET("/pop", func(c *gin.Context) {
		channels := gconv.Int(c.DefaultQuery("channel", "10"))
		channelType := gconv.Int(c.DefaultQuery("channel_type", "1"))
		if channelType == 1 {
			channel2.ChUse.Push(gconv.Int(channels))
			c.JSON(200, gin.H{
				"message": "ok!!!",
			})
		} else {
			channel2.F(c, gconv.Int(channels))
		}
	})

	return router
}
