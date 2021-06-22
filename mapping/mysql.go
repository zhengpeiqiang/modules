package mapping

import (
	"database/sql"
)

func MysqlRowToStruct(MysqlDb *sql.DB, Sql string) ([]ChatMessage, error) {
	//查询数据，取所有字段
	rows2, _ := MysqlDb.Query(Sql)
	//返回所有列
	cols, _ := rows2.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make([]ChatMessage, 0)
	//result := make(taskMap[int]taskMap[string]string, 0)
	for rows2.Next() {
		//填充数据
		_ = rows2.Scan(scans...)
		//每行数据
		//row := dto.ChatMessagePage{}
		row := make(map[string]interface{}, 0)
		//把vals中的数据复制到row中
		//fmt.Println(reflect.TypeOf(vals))
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		d := &ChatMessage{}
		err := SlStruct(row, d, 1)
		if err != nil {
			//fmt.Println(fmt.Sprintf("err ==> %v", err))
			return nil, err
		}
		//make(taskMap[string]reflect.Value)
		//放入结果集
		//result[i] = row
		result = append(result, *d)
		i++
	}
	defer func() {
		_ = rows2.Close()
	}()
	return result, nil
}
