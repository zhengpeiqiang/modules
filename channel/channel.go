package channel

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"modules/redisModules"
	"time"
)

var chNum = 1000

var ChUse *channelVar

type channelVar struct {
	ch chan int
}

func NewChan() *channelVar {
	c := &channelVar{
		ch: make(chan int, chNum),
	}
	return c
}

func (c *channelVar) Push(int2 int) {
	c.ch <- int2
}

func (c *channelVar) Pop() {
	if len(c.ch) > 0 {
		for i := 0; i < chNum; i++ {
			go func() {
				f(<-c.ch)
			}()
		}
	}
}

type Result struct {
	Channel int
}

func ResponseSuccess(data interface{}, isRedis bool, redisKey ...string) {
	if isRedis {
		jsonContent, err := json.Marshal(data)
		if err != nil {
			return
		}
		redisModules.ClientRedis.Set(redisKey[0], jsonContent, 60*time.Second)
	}
}

func f(int2 int) {
	redisKey := "f_"+gconv.String(grand.N(1,9))
	redisCache := redisModules.ClientRedis.Get(redisKey)
	val, err := redisCache.Result()
	if err == nil {
		result := Result{}
		err = json.Unmarshal([]byte(val), &result)
		if err != nil {
			return
		}
		ResponseSuccess(&result, false)
		return
	} else {
		if err != redis.Nil {
			return
		}
	}
	fmt.Println("redisKey => ",redisKey," int => ",int2)
	ResponseSuccess(Result{
		Channel: int2,
	}, true, redisKey)
	time.Sleep(50 * time.Millisecond)
}

func F(c *gin.Context, int2 int) {
	redisKey := "f"
	redisCache := redisModules.ClientRedis.Get(redisKey)
	val, err := redisCache.Result()
	if err == nil {
		result := Result{}
		err = json.Unmarshal([]byte(val), &result)
		if err != nil {
			return
		}
		ResponseSuccess(&result, false)
		return
	} else {
		if err != redis.Nil {
			return
		}
	}
	fmt.Println(int2)
	ResponseSuccess(Result{
		Channel: int2,
	}, true, "f")
	time.Sleep(1 * time.Second)
	c.JSON(200, gin.H{
		"message": "ok!!!",
	})
}
