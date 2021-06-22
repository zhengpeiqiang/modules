package taskMap

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/util/grand"
	"time"
)

const (
	Path              = "C:/Users/cyz/Desktop/git-project/Zmodules/task/taskData/data.json"
	IntervalTime      = 1 * time.Second
	IntervalSleepTime = 200 * time.Millisecond
	TimeFormat        = "2006-01-02 15:04:05"
)

func RandKey() string {
	key := ""
	grand.N(0, 1)
	for i := 0; i < 16; i++ {
		key = key + LettersANum[grand.N(0, len(LettersANum)-1)]
	}
	return key
}

var LettersANum = []string{
	"a", "B", "c", "D", "e", "F", "g", "H", "i", "J", "k", "L", "m", "N", "o", "P", "q", "R", "s", "T", "u", "V", "w", "X", "y", "Z",
	"A", "b", "C", "d", "E", "f", "G", "h", "I", "j", "K", "l", "M", "n", "O", "p", "Q", "r", "S", "t", "U", "v", "W", "x", "Y", "z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	//"~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "+", "_", "=",
}

var M *gmap.ListMap

//{"type":"send","key":"001","value":"001value","timestamp":11111111111}
type SendMap struct {
	Key       *gtype.String
	Value     *gtype.String
	Type      *gtype.String
	TimeStamp *gtype.Int
}
