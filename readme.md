<p align="center">
<img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/zhengpeiqiang.png" style="border-radius:50%;margin: 0 auto;" width="20%" />
</p>

<h1 align="center">
modules
</h1>

### organization

<a href="https://github.com/zhengpeiqiang"><img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/LCLCF_circle.png" style="width:60px;height:60px;margin: 0 auto;" width="8%" /></a>
**LCLCF**

#### mapping 
- 普通的map 映射 结构体
- mysql结构集 映射 结构体

[goto mapping](https://gitee.com/zhengpeiqiang/modules/tree/master/mapping)

#### sqltomodel 
- 数据表转换成model文件

[goto sqltomodel](https://gitee.com/zhengpeiqiang/modules/tree/master/sqltomodel)
```
package main

import (
	"modules/sqltomodel/conf"
	"modules/sqltomodel/dbtools"
)

func main() {
	// 首先初始化struct对象，用于保存配置文件信息
	var confS conf.Config
	// 通过toml.DecodeFile将toml配置文件的内容，解析到struct对象
	err := conf.GetConf("C:/Users/cyz/Desktop/git-project/Zmodules/sqltomodel/conf/mysql_map.toml", &confS)
	if err != nil {
		return
	}
	for _, v := range confS.Database {
		var MasterDbConfig = &dbtools.DbConf{
			Host:   v.Host,
			Port:   v.Port,
			User:   v.User,
			Pwd:    v.Pwd,
			DbName: "service",
		}
		dbtools.Init(MasterDbConfig)             //初始化数据库
		dbtools.Generate("C:/Users/cyz/Desktop/git-project/Zmodules/sqltomodel/models/", "service") //生成指定表信息，可变参数可传入多个表名
	}

}
```

#### pprof
- pprof的使用

[goto pprof](https://gitee.com/zhengpeiqiang/modules/tree/master/pprof)
```　　　　　　　　 
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

var dataS []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	dataS = append(dataS, sData)

	return sData
}

func main() {
	go func() {
		for {
			log.Println(Add("https://github.com/EDDYCJY"))
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
```

#### signal.Notify
- signal.Notify 的使用

[goto signal.Notify](https://gitee.com/zhengpeiqiang/modules/tree/master/signal)
```
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println(fmt.Sprintf("开始服务！！！"))
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println(fmt.Sprintf("结束服务！！！"))
}
```

#### task
- 计时的任务队列
- 可分为立即执行和定时执行

[goto task](https://gitee.com/zhengpeiqiang/modules/tree/master/task)
```
package main

import (
	"modules/task/taskFunc"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化，从文件中获取中断时间的任务
	taskFunc.Start()
	// 初始化一个计时器
	taskFunc.InitTimer()
	// 随机产生数据
	taskFunc.RandData()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 中断期间会以json格式保存任务
	taskFunc.Interrupt()
}
```

#### gin + pprof
- gin框架结合pprof 的使用

[goto gin + pprof](https://gitee.com/zhengpeiqiang/modules/tree/master/test)

#### go.mod
```
module modules

go 1.14

require (
	github.com/e421083458/gin_scaffold v0.0.0-20200502033629-13733b8c6903
	github.com/e421083458/golang_common v1.1.2
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/gogf/gf v1.15.7
	gopkg.in/go-playground/validator.v9 v9.29.0
)
```

#### Special thanks
<img width="80" height="50" src="http://liumurong.org/images/avatar.png"/>[Gin框架中使用pprof](http://liumurong.org/2019/12/gin_pprof/)

<img width="80" height="50" src="https://cdn2.jianshu.io/assets/web/nav-logo-4c7bbafe27adc892f3046e6978459bac.png"/>[Golang 大杀器之性能剖析 PProf](https://www.jianshu.com/p/4e4ff6be6af9)

