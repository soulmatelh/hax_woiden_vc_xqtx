package service

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Woiden(cookie string) (rq []string) {
	url := "https://woiden.id/vps-info/"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.62")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Host", "woiden.id")
	req.Header.Add("cookie", cookie)
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("\rWoiden status code error: %d %s\n", res.StatusCode, res.Status)
		time.Sleep(5 * time.Second)
	}
	bodys, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	//提取过期时间
	index := strings.Index(string(bodys), "Valid until")
	if index == -1 {
		return nil
	}
	body := bodys[index:]
	index = strings.Index(string(body), "</div>")
	if index == -1 {
		return nil
	}
	body = body[:index]
	index = strings.Index(string(body), `">`)
	if index == -1 {
		return nil
	}
	body = body[index+5:]
	body = []byte(strings.ReplaceAll(string(body), "\t", ""))
	body = []byte(strings.ReplaceAll(string(body), "\r", ""))
	body = []byte(strings.ReplaceAll(string(body), "\n", ""))
	bodyTime := strings.TrimSpace(string(body))
	rq = append(rq, bodyTime)
	//提取机器所属区域
	index = strings.Index(string(bodys), "Location")
	if index == -1 {
		return rq
	}
	body = bodys[index:]
	index = strings.Index(string(body), "</div>")
	if index == -1 {
		return rq
	}
	body = body[:index]
	index = strings.Index(string(body), `">`)
	if index == -1 {
		return rq
	}
	body = body[index+2:]
	//去除\n \t 等转义字符
	body = []byte(strings.ReplaceAll(string(body), "\t", ""))
	body = []byte(strings.ReplaceAll(string(body), "\r", ""))
	body = []byte(strings.ReplaceAll(string(body), "\n", ""))
	bodyLocation := strings.TrimSpace(string(body))
	rq = append(rq, bodyLocation)
	return rq
}
