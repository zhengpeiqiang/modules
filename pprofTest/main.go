package main

import (
	"fmt"
	"github.com/gogf/gf/os/gtimer"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	gtimer.AddSingleton(2*time.Second, func() {
		fmt.Println("ajhkdjashdkjha")
	})
	//go func() {
	//	fmt.Println("ajhkdjashdkjha")
	//}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
