package main

import (
	"flag"
	"fmt"
	"gin-temp/config"
	"gin-temp/extends/mysql"
	"gin-temp/extends/redis"
	"gin-temp/router"
	"gin-temp/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var (
	configFile = flag.String("c", "./config/config.yaml", "配置文件地址")
)

func main() {
	flag.Parse()
	if err := config.InitConfig(*configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化mysql数据库
	if err := mysql.DB.InitDB(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化redis
	if err := redis.Init(); err != nil {
		fmt.Println(err)
	}

	g := gin.New()
	router.LoadRouter(
		g,
		// 中间件
		middleware.Options,
		middleware.NoCache,
	)
	http.ListenAndServe(viper.GetString("server.addr")+":"+viper.GetString("server.port"), g)
}
