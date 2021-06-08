package util

import (
	"block-service/internal/conf"
	"bytes"
	"fmt"
	"html/template"

	"github.com/go-gomail/gomail"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strconv"
	"strings"
)

type (
	EmailNotify struct {
		SmtpS   string
		SmtpP   int
		Fromer  string
		Toers   []string
		Ccers   []string
		EUser   string
		Epasswd string
	}
)

func NewEmailNotify(mailConf *conf.Email) *EmailNotify {
	smtp_s_str := mailConf.Host
	smtp_p_str := mailConf.Port
	sender_str := mailConf.Sender
	passwd_str := mailConf.Password

	receivers := []string{}
	receiversStr := ""
	for _, receiverStr := range strings.Split(receiversStr, ";") {
		receivers = append(receivers, strings.TrimSpace(receiverStr))
	}

	smtp_p_int, _ := strconv.Atoi(smtp_p_str)

	email := &EmailNotify{
		SmtpS:   smtp_s_str,
		SmtpP:   smtp_p_int,
		Fromer:  sender_str,
		Toers:   receivers,
		Ccers:   []string{},
		EUser:   strings.Split(sender_str, "@")[0],
		Epasswd: passwd_str,
	}
	return email
}
func (en *EmailNotify) SendNotifyWithFile(title, content string) bool {
	return en.SendNotifyWithFileAndAttach(title, content, "", "")
}
func (en *EmailNotify) SendNotifyWithFileAndAttach(title, content, filePath, newName string) bool {
	msg := gomail.NewMessage(gomail.SetCharset("utf-8"))
	msg.SetHeader("From", en.Fromer)
	msg.SetHeader("To", en.Toers...)
	msg.SetHeader("Subject", title)

	msg.SetBody("text/html", en.renderNotify(content))

	//防止中文文件名乱码
	if filePath != "" {
		fileName, _ := Utf8ToGbk([]byte(newName))
		msg.Attach(filePath, gomail.Rename(string(fileName)))
	}

	mailer := gomail.NewDialer(en.SmtpS, en.SmtpP, en.EUser, en.Epasswd)
	if err := mailer.DialAndSend(msg); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return true
}

func (en *EmailNotify) renderNotify(content string) string {
	tplStr := `<html>
				<body>
				 {<!-- -->{.}}
				</table>
				</body>
				</html>`

	outBuf := &bytes.Buffer{}
	tpl := template.New("email notify template")
	tpl, _ = tpl.Parse(tplStr)
	tpl.Execute(outBuf, content)

	return outBuf.String()
}
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
