package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			fmt.Println(nil)
		}
	}()
	time.Sleep(1*time.Millisecond)
	fmt.Println("aaaaaaaaa")
}
