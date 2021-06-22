package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	// 创建trace文件
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 启动trace goroutine
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	a := "aajsdhjashdjsaghdgkahjdsg"
	b := a
	c := b
	d := c
	e := d
	f := e
	g := f
	h := g
	i := h
	j := i
	k := j
	l := k
	m := l
	n := m
	fmt.Println(n)

	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	//<-quit

}
