package taskFunc

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"modules/task/taskMap"
	"time"
)

//初始化，从文件中获取中断时间的任务，中断期间会以json格式保存任务
func Start() {
	//初始化gmap.listmap
	taskMap.M = gmap.NewListMap(true)
	//
	var a = make(map[string]taskMap.SendMap, 0)
	sm := gfile.GetBytes(taskMap.Path)
	if string(sm) != "" {
		//fmt.Println(fmt.Sprintf("sm ==> %v", string(sm)))
		err := json.Unmarshal(sm, &a)
		if err != nil {
			glog.Error(err)
			return
		}
		//放入listmap
		for _, v := range a {
			taskMap.M.Set(v.Key.String(), taskMap.SendMap{
				Key:       v.Key,
				Value:     v.Value,
				Type:      v.Type,
				TimeStamp: v.TimeStamp,
			})
		}
	}
}

//删除
func Pop(key string) {
	taskMap.M.Remove(key)
	//fmt.Println(taskMap.M.String())
}

//中断时将数据写入send_map.txt
func Interrupt() {
	text, err := taskMap.M.MarshalJSON()
	if err != nil {
		glog.Error(err)
		return
	}
	err = gfile.PutContents(taskMap.Path, string(text))
	if err != nil {
		glog.Error(err)
		return
	}
}

//初始化
func InitTimer() {
	//遍历listmap
	gtimer.AddSingleton(taskMap.IntervalTime, func() {
		if taskMap.M.Size() > 0 {
			isPop := 0
			mapKey := ""
			//该方法会将 整个 list 的数据遍历一遍
			taskMap.M.Iterator(func(key, value interface{}) bool {
				valueSendMap := value.(taskMap.SendMap)
				if valueSendMap.Type == nil {
					mapKey = key.(string)
					isPop = 1
					fmt.Println(fmt.Sprintf("[错误] key ==> %v ,还剩 %v 个", mapKey, taskMap.M.Size()-1))
					return false
				}
				switch valueSendMap.Type.Val() {
				case "send":
					if valueSendMap.TimeStamp.Val() < gconv.Int(time.Now().Unix()) {
						// TODO
						mapKey = key.(string)
						isPop = 1
						fmt.Println(fmt.Sprintf("[删除][send] key ==> %v ,还剩 %v 个", mapKey, taskMap.M.Size()-1))
						return false
					}
				default:
					mapKey = key.(string)
					isPop = 1
					fmt.Println(fmt.Sprintf("[删除][normal] key ==> %v ,还剩 %v 个", mapKey, taskMap.M.Size()-1))
					return false
				}
				return true
			})
			if isPop == 1 {
				Pop(mapKey)
				isPop = 0
			}
		}
		time.Sleep(taskMap.IntervalSleepTime)
		fmt.Println(fmt.Sprintf("[普通] time ==> %v ,还剩 %v 个", time.Now().Format(taskMap.TimeFormat), taskMap.M.Size()))
	})
}

func RandData() {
	gtimer.AddSingleton(15*time.Second, func() {
		AddMsgForTest(grand.N(1, 3))
		time.Sleep(500 * time.Millisecond)
	})
}

func AddMsgForTest(num int) {
	for i := 0; i < num; i++ {
		key := taskMap.RandKey()
		value := taskMap.RandKey()
		types := ""
		typesRand := grand.N(0, 1)
		if typesRand == 1 {
			types = "send"
		} else {
			types = "normal"
		}
		timeout := gconv.String(time.Now().Add(time.Duration(grand.N(5, 10)) * time.Second).Unix())
		taskMap.M.Set(key, taskMap.SendMap{
			Key:       gtype.NewString(key),
			Value:     gtype.NewString(value),
			Type:      gtype.NewString(types),
			TimeStamp: gtype.NewInt(gconv.Int(timeout)),
		})
	}
}
