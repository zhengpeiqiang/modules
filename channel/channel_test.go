package channel

import (
	"fmt"
	"github.com/gogf/gf/os/gtimer"
	"testing"
	"time"
)

func TestF(t *testing.T) {
	num := 1
	ch := make(chan int, 2)
	gtimer.AddSingleton(1*time.Second, func() {
		//
		if num == 1 {
			num = 2
			if len(ch) > 0 {
				fmt.Println("管道 有数据 ", time.Now().Format("2006-01-02 15:04:05"))
			} else {
				fmt.Println("管道 无数据 ", time.Now().Format("2006-01-02 15:04:05"))
			}
		} else if num == 2 {
			fmt.Println("[start]执行到 2 time ", time.Now().Format("2006-01-02 15:04:05"))
			num = 1
			ch <- 1
			ch <- 2
			fmt.Println("[end]执行到 2 time ", time.Now().Format("2006-01-02 15:04:05"))
		}
	})
	gtimer.AddSingleton(3*time.Second, func() {
		fmt.Println("[start]执行到 3 time ", time.Now().Format("2006-01-02 15:04:05"))
		<-ch
		fmt.Println("[end]执行到 3 time ", time.Now().Format("2006-01-02 15:04:05"))
	})

	//quit := make(chan os.Signal)
	//signalTest.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	//<-quit

	//redisModules.ClientRedis.Set("bbbbbbb", "aaaaaaaa", 10*time.Second)
	//b := redisModules.ClientRedis.Get("bbbbbbb").String()
	//redisModules.ClientRedis.Del(b)
	//c := make(chan int, 10)
	//gtimer.AddSingleton(1*time.Second, func() {
	//	c <- grand.N(10, 100)
	//})
	//gtimer.AddSingleton(1*time.Second, func() {
	//	if len(c) > 0 {
	//		a := make(chan int, 0)
	//		for i := 0; i < chNum; i++ {
	//			rand := <-c
	//			fmt.Println(rand, " time => ", time.Now().Format("2006-01-02 15:04:05"))
	//			<-a
	//			time.Sleep(2 * time.Second)
	//			fmt.Println(rand, " time => ", time.Now().Format("2006-01-02 15:04:05"), " end ")
	//		}
	//	}
	//})
	//
	//quit := make(chan os.Signal)
	//signalTest.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
}
