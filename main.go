package main

import (
	"edulx/hax_woiden_vc_xqtx/api"
	"edulx/hax_woiden_vc_xqtx/service"
	"embed"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
)

//go:embed template/*
var f embed.FS

func main() {
	go service.Ticker()
	go service.TickerTX()
	go service.Write()
	r := gin.Default()
	tmpl := template.Must(template.New("").ParseFS(f, "template/*.html"))
	r.SetHTMLTemplate(tmpl)
	r.GET("/", api.Loginx)
	r.POST("/login", api.Login)
	r.GET("/login", api.Login)
	r.POST("/", api.Add)
	r.POST("/del", api.Del)
	r.POST("/put", api.Put)
	if len(os.Args) > 1 {
		r.Run(":" + os.Args[1])
	} else {
		r.Run(":8080")
	}

}

// Ticker 定时任务

func init() {
	err := service.Date.ReadJson()
	if err != nil {
		fmt.Println("static error")
		return
	}
}
