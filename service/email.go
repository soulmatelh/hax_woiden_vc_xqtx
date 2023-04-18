package service

import (
	"fmt"
	"net/smtp"
)

func (x *X) Emails(to string, title string, content string) bool {
	subject := fmt.Sprintf("Subject: %s\r\n", title)
	send := fmt.Sprintf("From: %s 续期提醒\r\n", x.Email.From)
	receiver := fmt.Sprintf("To: %s\r\n", to)
	contentType := "Content-Type: text/html" + "; charset=UTF-8\r\n\r\n"
	msg := []byte(subject + send + receiver + contentType + content)
	addr := x.Email.Smtp + ":" + x.Email.Port
	auth := smtp.PlainAuth("", x.Email.From, x.Email.Key, x.Email.Smtp)
	from := x.Email.From
	var tos = []string{to}
	err := smtp.SendMail(addr, auth, from, tos, msg)
	if err != nil {
		fmt.Println("------------------------------------------")
		fmt.Println(err.Error())
		fmt.Println("------------------------------------------")
		return false
	}
	return true
}

func html(mp map[string][]FormEmail) map[string]string {
	rqs := make(map[string]string)
	for k, v := range mp {
		rq := "<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\">\n\t<title>续期通知</title>\n\t<style>\n\t\tbody {\n\t\t\tfont-family: Arial, sans-serif;\n\t\t\tbackground-color: #f8f8f8;\n\t\t\tmargin: 0;\n\t\t\tpadding: 0;\n\t\t}\n\t\th2 {\n\t\t\ttext-align: center;\n\t\t\tmargin-top: 30px;\n\t\t}\n\t\ttable {\n\t\t\tborder-collapse: collapse;\n\t\t\twidth: 80%;\n\t\t\tmargin: 0 auto;\n\t\t\tbackground-color: #fff;\n\t\t\tbox-shadow: 0 0 20px rgba(0, 0, 0, 0.1);\n\t\t\tborder-radius: 10px;\n\t\t\toverflow: hidden;\n\t\t\tmargin-bottom: 30px;\n\t\t}\n\t\tth, td {\n\t\t\tpadding: 12px 15px;\n\t\t\ttext-align: left;\n\t\t\tborder-bottom: 1px solid #ddd;\n\t\t}\n\t\tth {\n\t\t\tbackground-color: #008080;\n\t\t\tcolor: #fff;\n\t\t\tfont-weight: bold;\n\t\t}\n\t\ttr:hover {\n\t\t\tbackground-color: #f5f5f5;\n\t\t}\n\t</style></head>\n<body>\n\t<h2>服务器过期时间</h2>\n\t<table>\n\t\t<thead>\n\t\t\t<tr>\n\t\t\t\t<th>名称</th>\n\t\t\t\t<th>服务器名称</th>\n\t\t\t\t<th>过期时间</th>\n\t\t\t</tr>\n\t\t</thead>\n\t\t<tbody>"
		for _, l := range v {
			rq += fmt.Sprintf("\t\t\t<tr>\n\t\t\t\t<td>%s</td>\n\t\t\t\t<td>%s</td>\n\t\t\t\t<td>%s</td>\n\t\t\t</tr>", l.Name, l.VpsName, l.Times)
		}
		rq += "\t\t</tbody>\n\t</table>\n</body>\n</html>"
		rqs[k] = rq
	}
	return rqs
}
