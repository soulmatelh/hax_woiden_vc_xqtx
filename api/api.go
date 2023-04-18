package api

import (
	"edulx/hax_woiden_vc_xqtx/service"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type Status struct {
	Name        string `json:"Name"`
	DelName1    string `json:"DelName1"`
	DelName2    string `json:"DelName2"`
	PutName1    string `json:"PutName1"`
	PutName2    string `json:"PutName2"`
	Cookie      string `json:"Cookie"`
	Web         string `json:"Web"`
	Email       string `json:"Email"`
	Uuid        string `json:"Uuid"`
	Value       string `json:"Value"`
	Description string `json:"description"`
	Time        string `json:"Time"`
	VpsNAME     string `json:"VpsNAME"`
}

func Index(c *gin.Context) {
	if IP != c.ClientIP() {
		c.HTML(200, "login.html", gin.H{
			"Error": "",
		})
		return
	}
	x := service.X{}
	err := x.ReadJson()
	if err != nil {
		c.String(201, "<h1>static error</h1>")
		//查询文件是否存在
		_, err1 := os.Stat("static.json")
		if err1 != nil {
			c.String(201, "<h2>static.json not found</h2>")
			c.String(201, "<h2>create static.json</h2>")
			x := service.X{
				Name:    "",
				Version: "",
				Woiden:  nil,
				Hax:     nil,
				Vc:      nil,
				Admin: struct {
					Password string `json:"password"`
				}{},
			}
			err := x.WriteJson()
			if err != nil {
				c.String(201, "<h2>create static.json error</h2>")
				return
			}
			c.String(201, "<h2>create static.json success</h2>")
			c.String(201, "<h2>please refresh</h2>")
		}
		return
	}
	var list []Status
	for _, v := range x.Hax {
		if len(v.Time) < 10 {
			v.Time = "当前未获取到过期时间请等待定时任务进程"
		}
		list = append(list, Status{
			Name:        v.Name,
			Value:       "OK",
			DelName1:    "#del" + v.Uuid,
			DelName2:    "del" + v.Uuid,
			PutName1:    "#put" + v.Uuid,
			PutName2:    "put" + v.Uuid,
			Web:         "hax.co.id",
			Cookie:      v.Cookie,
			VpsNAME:     "HAX-" + v.VpsNAME,
			Uuid:        v.Uuid,
			Description: v.Code,
			Email:       v.TO,
			Time:        v.Time[0:10],
		})
	}
	for _, v := range x.Woiden {
		if len(v.Time) < 10 {
			v.Time = "当前未获取到过期时间请等待定时任务进程"
		}
		list = append(list, Status{
			Name:        v.Name,
			Value:       "OK",
			DelName1:    "#del" + v.Uuid,
			DelName2:    "del" + v.Uuid,
			PutName1:    "#put" + v.Uuid,
			PutName2:    "put" + v.Uuid,
			Cookie:      v.Cookie,
			VpsNAME:     "WOIDEN-" + v.VpsNAME,
			Web:         "woiden.id",
			Uuid:        v.Uuid,
			Email:       v.TO,
			Description: v.Code,
			Time:        v.Time[0:10],
		})
	}
	for _, v := range x.Vc {
		if len(v.Time) < 10 {
			v.Time = "当前未获取到过期时间请等待定时任务进程"
		}
		list = append(list, Status{
			Name:        v.Name,
			Value:       "OK",
			DelName1:    "#del" + v.Uuid,
			DelName2:    "del" + v.Uuid,
			PutName1:    "#put" + v.Uuid,
			PutName2:    "put" + v.Uuid,
			Cookie:      v.Cookie,
			VpsNAME:     "VC-" + v.VpsNAME,
			Web:         "free.vps.vc",
			Uuid:        v.Uuid,
			Description: v.Code,
			Email:       v.TO,
			Time:        v.Time[0:10],
		})
	}
	c.HTML(200, "status1.html", gin.H{
		"title":      "续期提醒",
		"statusList": list,
	})
}

var IP string

// Add 添加监控
func Add(c *gin.Context) {
	if IP != c.ClientIP() {
		c.HTML(200, "login.html", gin.H{
			"Error": "",
		})
		return
	}
	//service.Tickers.Stop()
	//停止计时器
	name := c.PostForm("name")
	cookie := c.PostForm("cookie")
	web := c.PostForm("web")
	tx := c.PostForm("tx")
	mubiao := c.PostForm("mubiao")
	var x = service.Date
	var a struct {
		Uuid    string `json:"uuid"`
		Name    string `json:"name"`
		Cookie  string `json:"cookie"`
		Time    string `json:"time"`
		Code    string `json:"code"`
		TC      string `json:"tc"`
		TO      string `json:"to"`
		VpsNAME string `json:"vpsNAME"`
	}
	a.Name = name
	a.Cookie = cookie
	a.Uuid = service.Uuid()
	a.TC = tx
	a.TO = mubiao
	switch web {
	case "1":
		t := service.Woiden(a.Cookie)
		if t == nil {
			a.Code = "Warning"
		}
		layout := "January 02, 2006"
		date, err := time.Parse(layout, t[0])
		if err != nil {
			a.Code = "error"
		}

		// 转换为时间戳
		timestamp := date.String()
		a.Time = timestamp
		a.Code = "OK"
		if len(t) >= 2 {
			a.VpsNAME = t[1]
		}
		service.Date.Woiden = append(service.Date.Woiden, a)
	case "2":
		t := service.Hax(a.Cookie)
		if t == nil {
			a.Code = "Warning"
		}
		layout := "January 02, 2006"
		date, err := time.Parse(layout, t[0])
		if err != nil {
			a.Code = "error"
		}

		// 转换为时间戳
		timestamp := date.String()
		a.Time = timestamp
		a.Code = "OK"
		if len(t) >= 2 {
			a.VpsNAME = t[1]
		}
		service.Date.Hax = append(service.Date.Hax, a)
	case "3":
		t := service.Vc(a.Cookie)
		if t == nil {
			a.Code = "Warning"
		}
		layout := "January 02, 2006"
		date, err := time.Parse(layout, t[0])
		if err != nil {
			a.Code = "error"
		}

		// 转换为时间戳
		timestamp := date.String()
		a.Time = timestamp
		a.Code = "OK"
		if len(t) >= 2 {
			a.VpsNAME = t[1]
		}
		service.Date.Vc = append(x.Vc, a)
	}

	err := x.WriteJson()
	if err != nil {
		c.HTML(200, "add.html", gin.H{
			"code": "添加失败",
		})

	}
	c.HTML(200, "add.html", gin.H{
		"code": "添加成功",
	})
}

func Del(c *gin.Context) {
	if IP != c.ClientIP() {
		c.HTML(200, "login.html", gin.H{
			"Error": "",
		})
		return
	}
	uuid := c.PostForm("uuid")
	if uuid == "" {
		c.HTML(200, "add.html", gin.H{
			"code": "删除失败",
		})
	}
	service.Date.Del(uuid)
	err := service.Date.WriteJson()
	if err != nil {
		c.HTML(200, "add.html", gin.H{
			"code": "修改配置失败",
		})
		return
	}
	c.HTML(200, "add.html", gin.H{
		"code": "删除成功",
	})
}

func Put(c *gin.Context) {
	if IP != c.ClientIP() {
		c.HTML(200, "login.html", gin.H{
			"Error": "",
		})
		return
	}
	name := c.PostForm("name")
	cookie := c.PostForm("cookie")
	uuid := c.PostForm("uuid")
	tx := c.PostForm("tx")
	mubiao := c.PostForm("mubiao")
	if uuid == "" {
		c.HTML(200, "add.html", gin.H{
			"code": "修改失败",
		})
		return
	}
	for i, v := range service.Date.Woiden {
		if v.Uuid == uuid {
			service.Date.Woiden[i].Name = name
			service.Date.Woiden[i].Cookie = cookie
			service.Date.Woiden[i].TC = tx
			service.Date.Woiden[i].TO = mubiao
		}
	}
	for i, v := range service.Date.Hax {
		if v.Uuid == uuid {
			service.Date.Hax[i].Name = name
			service.Date.Hax[i].Cookie = cookie
			service.Date.Hax[i].TC = tx
			service.Date.Hax[i].TO = mubiao
		}
	}
	for i, v := range service.Date.Vc {
		if v.Uuid == uuid {
			service.Date.Vc[i].Name = name
			service.Date.Vc[i].Cookie = cookie
			service.Date.Vc[i].TC = tx
			service.Date.Vc[i].TO = mubiao
		}
	}

	err := service.Date.WriteJson()
	if err != nil {
		c.HTML(200, "add.html", gin.H{
			"code": "修改配置成功，存储配置失败",
		})
		return
	}
	c.HTML(200, "add.html", gin.H{
		"code": "修改成功",
	})
}

func Login(c *gin.Context) {
	password := c.PostForm("password")
	var x service.X
	err := x.ReadJson()
	if err != nil {
		return
	}
	if IP == c.ClientIP() {
		Index(c)
		return
	}
	if password == x.Admin.Password {
		//Index(c)
		//记录登录IP
		c.HTML(200, "login.html", gin.H{
			"Meta":  "0;url=/",
			"Error": "",
		})
		IP = c.ClientIP()
	} else {
		c.HTML(200, "login.html", gin.H{
			"Meta":  "",
			"Error": "登录失败",
		})
	}
}

func Loginx(c *gin.Context) {
	if IP == c.ClientIP() {
		Index(c)
		return
	}
	c.HTML(200, "login.html", gin.H{
		"Error": "",
	})
}
