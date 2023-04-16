package service

import (
	"fmt"
	"time"
)

type VPSCookie struct {
	HAX    []string `yaml:"hax"`
	Woiden []string `yaml:"woiden"`
	VC     []string `yaml:"vc"`
}

type FormEmail struct {
	Name    string `json:"name"`
	VpsName string `json:"vpsName"`
	Times   string `json:"times"`
}

var Date X

func Run() {
	var x = Date
	for i := range x.Hax {
		t := Hax(x.Hax[i].Cookie)
		if t == nil {
			x.Hax[i].Code = "Warning"
			continue
		}
		layout := "January 02, 2006"
		date, err := time.Parse(layout, t[0])
		if err != nil {
			x.Hax[i].Code = "error"
			return
		}

		// 转换为时间戳
		timestamp := date.String()
		x.Hax[i].Time = timestamp
		x.Hax[i].Code = "OK"
		if len(t) >= 2 {
			x.Hax[i].VpsNAME = t[1]
		}
	}
	for i := range x.Woiden {
		t := Woiden(x.Woiden[i].Cookie)
		if t == nil {
			x.Woiden[i].Code = "Warning"
			continue
		}
		layout := "January 02, 2006"
		date, err := time.Parse(layout, t[0])
		if err != nil {
			x.Woiden[i].Code = "error"
			return
		}
		// 转换为时间戳
		timestamp := date.String()
		x.Woiden[i].Time = timestamp
		x.Woiden[i].Code = "OK"
		if len(t) >= 2 {
			x.Woiden[i].VpsNAME = t[1]
		}
	}
	for i := range x.Vc {
		t := Vc(x.Vc[i].Cookie)
		if t == nil {
			x.Vc[i].Code = "Warning"
			continue
		}
		layout := "January 02, 2006"
		date, err := time.Parse(layout, t[0])
		if err != nil {
			x.Vc[i].Code = "error"
			return
		}
		// 转换为时间戳
		timestamp := date.String()
		x.Vc[i].Time = timestamp
		x.Vc[i].Code = "OK"
		if len(t) >= 2 {
			x.Vc[i].VpsNAME = t[1]
		}
	}
	fmt.Println("success " + time.Now().String())
}

func TX(ts int64) {
	var x = Date
	mp := map[string][]FormEmail{}
	for _, v := range x.Hax {
		if len(v.Time) < 19 {
			continue
		}
		t, err1 := time.Parse("2006-01-02 15:04:05", v.Time[0:19])
		if err1 != nil {
			fmt.Println("解析时间字符串错误：", err1)
			continue
		}
		timestamp := t.Unix()
		if timestamp-time.Now().Unix() < ts {
			var e FormEmail
			e.Name = v.Name
			e.VpsName = "HAX-" + v.VpsNAME
			e.Times = v.Time[0:10]
			if v.TC == "email" {
				if v.TO == "" {
					mp[x.Email.From] = append(mp[x.Email.From], e)
				} else {
					mp[v.TO] = append(mp[v.TO], e)
				}
			}
		}
	}
	for _, v := range x.Woiden {
		if len(v.Time) < 19 {
			continue
		}
		t, err1 := time.Parse("2006-01-02 15:04:05", v.Time[0:19])
		if err1 != nil {
			fmt.Println("解析时间字符串错误：", err1)
			continue
		}
		timestamp := t.Unix()
		if timestamp-time.Now().Unix() < ts {
			var e FormEmail
			e.Name = v.Name
			e.VpsName = "WOIDEN-" + v.VpsNAME
			e.Times = v.Time[0:10]
			if v.TC == "email" {
				if v.TO == "" {
					mp[x.Email.From] = append(mp[x.Email.From], e)
				} else {
					mp[v.TO] = append(mp[v.TO], e)
				}
			}
		}
	}
	for _, v := range x.Vc {
		if len(v.Time) < 19 {
			continue
		}
		t, err1 := time.Parse("2006-01-02 15:04:05", v.Time[0:19])
		if err1 != nil {
			fmt.Println("解析时间字符串错误：", err1)
			continue
		}
		timestamp := t.Unix()
		if timestamp-time.Now().Unix() < ts {
			var e FormEmail
			e.Name = v.Name
			e.VpsName = "VC-" + v.VpsNAME
			e.Times = v.Time[0:10]
			if v.TC == "email" {
				if v.TO == "" {
					mp[x.Email.From] = append(mp[x.Email.From], e)
				} else {
					mp[v.TO] = append(mp[v.TO], e)
				}
			}
		}
	}
	mps := html(mp)
	for k, v := range mps {
		err := x.Emails(k, "VPS到期提醒", v)
		if err != true {
			fmt.Println("发送邮件失败：", k)
		} else {
			fmt.Println("发送邮件成功：", k)
		}
	}

}
