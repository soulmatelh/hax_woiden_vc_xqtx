package service

import (
	"encoding/json"
	"os"
)

type X struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Woiden  []struct {
		Uuid    string `json:"uuid"`
		Name    string `json:"name"`
		Cookie  string `json:"cookie"`
		Time    string `json:"time"`
		Code    string `json:"code"`
		TC      string `json:"tc"`
		TO      string `json:"to"`
		VpsNAME string `json:"vpsNAME"`
	} `json:"woiden"`
	Hax []struct {
		Uuid    string `json:"uuid"`
		Name    string `json:"name"`
		Cookie  string `json:"cookie"`
		Time    string `json:"time"`
		Code    string `json:"code"`
		TC      string `json:"tc"`
		TO      string `json:"to"`
		VpsNAME string `json:"vpsNAME"`
	} `json:"hax"`
	Vc []struct {
		Uuid    string `json:"uuid"`
		Name    string `json:"name"`
		Cookie  string `json:"cookie"`
		Time    string `json:"time"`
		Code    string `json:"code"`
		TC      string `json:"tc"`
		TO      string `json:"to"`
		VpsNAME string `json:"vpsNAME"`
	} `json:"vc"`
	Admin struct {
		Password string `json:"password"`
	} `json:"admin"`
	Email struct {
		From string `json:"from"`
		Key  string `json:"key"`
		Smtp string `json:"smtp"`
		Port string `json:"port"`
	} `json:"email"`
}

//读取配置文件

func (x *X) ReadJson() error {
	data, err := os.ReadFile("static.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &x)
	if err != nil {
		return err
	}
	return nil
}

func (x *X) WriteJson() error {
	jsonStr, err := json.Marshal(x)
	if err != nil {
		return err
	}
	file, err := os.Create("static.json")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	_, err = file.Write(jsonStr)
	if err != nil {
		return err
	}
	return nil
}

func (x *X) Del(uuid string) {
	var news X
	for _, v := range x.Hax {
		if v.Uuid == uuid {
			continue
		}
		news.Hax = append(news.Hax, v)
	}
	for _, v := range x.Woiden {
		if v.Uuid == uuid {
			continue
		}
		news.Woiden = append(news.Woiden, v)
	}
	for _, v := range x.Vc {
		if v.Uuid == uuid {
			continue
		}
		news.Vc = append(news.Vc, v)
	}
	x.Hax = news.Hax
	x.Woiden = news.Woiden
	x.Vc = news.Vc
}
