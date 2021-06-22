package main

import (
	"fmt"
)

func main() {

	//modelCache := cache.ModelCache
	//
	//key := "iii"
	//
	//data := structModel.A{
	//	Name: "zpq",
	//	Age:  26,
	//	Job:  "Computer Server Engineer",
	//}
	//
	//dataNew := modelCache.AddSlice(key, nil, data)
	//
	//fmt.Println(fmt.Sprintf("data => %v", dataNew))
	//
	//key = "jkl"
	//fmt.Println(fmt.Sprintf("data => %v", modelCache.AddSlice(key, nil, nil)))

	//t := T{
	//	msg: "asgjhagsj",
	//}
	fmt.Println("start")
	i := int(9223372036854775807)
	//i := int(18446744073709551615)
	//256^8 = 18446744073709551615
	ii := int(300)
	iii := int(300)
	iiii := int(300)
	iiiii := int(300)
	var s string
	ss := ""
	for i := 0; i < 50; i++ {
		ss = ss + string("hagjsgjadgasjdgajdgjagsdjhagdjhgdsajdhgas")
	}
	fmt.Println(&i, &ii, &iii, &iiii, &iiiii, &s, &ss)
	//addFunc(t,i,s)
	//addFunc(t,ii,ss)

	t := TT{
		s: "11111",
	}
	tt := "11111"
	ttt := &TT{
		s: "11111",
	}
	if &t == ttt {
		fmt.Println("same")
	}
	tttt := "11111"
	fmt.Println(&t.s, &ttt.s)
	fmt.Println(&tt, &tttt)
}

type TT struct {
	i int
	s string
}

type T struct {
	msg string
}
