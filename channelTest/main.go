package main

import (
	"github.com/gogf/gf/os/gtimer"
	channel2 "modules/channel"
	"modules/channelTest/router"
	_ "modules/redisModules"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	channel2.ChUse = channel2.NewChan()
	gtimer.AddSingleton(1*time.Second, func() {
		channel2.ChUse.Pop()
	})
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
