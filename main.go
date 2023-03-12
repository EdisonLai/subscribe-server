package main

import (
	"server/route"

	"github.com/labstack/echo"
)

func main() {
	// 创建Echo实例
	e := echo.New()

	route.SetRoute(e)

	// 启动服务器
	e.Logger.Fatal(e.Start(":8080"))
}
