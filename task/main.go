package main

import (
	"modules/task/taskFunc"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化，从文件中获取中断时间的任务
	taskFunc.Start()
	// 初始化一个计时器
	taskFunc.InitTimer()
	// 随机产生数据
	taskFunc.RandData()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 中断期间会以json格式保存任务
	taskFunc.Interrupt()
}
