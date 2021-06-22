package main

import (
	"database/sql"
	"fmt"
	"modules/mapping"
)

func t1(a interface{}) error {
	row := make(map[string]interface{}, 0)
	row["group_id"] = "asjhdajkhdskjhasdkjhaskjdhk"
	err := mapping.SlStruct(row, a, 1)
	return err
}

func t2() ([]mapping.ChatMessage, error) {
	mysql, err := sql.Open("", "")
	if err != nil {
		fmt.Println(fmt.Sprintf("err ==> %v", err))
		return nil, err
	}
	sqlS := "select group_id from chat_msg"
	res, err := mapping.MysqlRowToStruct(mysql, sqlS)
	if err != nil {
		fmt.Println(fmt.Sprintf("err ==> %v", err))
		return nil, err
	}
	return res, nil
}

func main() {

	// 普通的map 映射 结构体
	d := &mapping.ChatMessage{
		GroupId: "asgjsa",
	}
	err := t1(d)
	if err != nil {
		fmt.Println(fmt.Sprintf("err ==> %v", err))
	}
	fmt.Println(fmt.Sprintf("d ==> %v", d))
	fmt.Println(fmt.Sprintf("d.GroupId ==> %v", d.GroupId))

	// 查询数据库之后的map 映射 结构体
	//c := mapping.ChatMessage{}
	//cc := &[]mapping.ChatMessage{}
	cc, err := t2()
	if err != nil {
		fmt.Println(fmt.Sprintf("err ==> %v", err))
	}
	fmt.Println(fmt.Sprintf("cc ==> %v", cc))
}
