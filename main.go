package main

import (
	"github.com/Albert-Zhan/httpc"
	"github.com/unknwon/goconfig"
	"github.com/ztino/jd_seckill/cmd"
	"github.com/ztino/jd_seckill/common"
	"log"
	"net/http"
	"os"
	"runtime"
)

func init()  {
	//客户端设置初始化
	common.Client=httpc.NewHttpClient()
	common.CookieJar=httpc.NewCookieJar()
	common.Client.SetCookieJar(common.CookieJar)
	common.Client.SetRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})

	//配置文件初始化
	confFile:="./conf.ini"
	var err error
	if common.Config,err=goconfig.LoadConfigFile(confFile);err!=nil {
		log.Println("配置文件不存在，程序退出")
		os.Exit(0)
	}

	//抢购状态管道
	common.SeckillStatus=make(chan bool)
}

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}