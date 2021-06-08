package util

//
//import (
//	"github.com/go-gomail/gomail"
//	"github.com/google/wire"
//	"strings"
//)
//
//// ProviderSet is util providers.
//var ProviderSet = wire.NewSet(NewEmail)
//
//type EmailParam struct {
//	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
//	ServerHost string
//	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
//	ServerPort int
//	// FromEmail　发件人邮箱地址
//	FromEmail string
//	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
//	FromPasswd string
//	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
//	Toers string
//	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
//	CCers string
//}
//
//func NewEmail(ep *EmailParam) *gomail.Dialer{
//	d := gomail.NewDialer(ep.ServerHost, ep.ServerPort, ep.FromEmail, ep.FromPasswd)
//	return d
//}
//
//// SendEmail body支持html格式字符串
//func ()SendEmail(tos, subject, body string) {
//	m := gomail.NewMessage()
//	// 主题
//	m.SetHeader("Subject", subject)
//
//	// 正文
//	m.SetBody("text/html", body)
//
//	m := gomail.NewMessage()
//
//	if len(tos) == 0 {
//		return
//	}
//
//	for _, tmp := range strings.Split(tos, ",") {
//		tos = append(tos, strings.TrimSpace(tmp))
//	}
//
//	// 收件人可以有多个，故用此方式
//	m.SetHeader("To", toers...)
//
//	//抄送列表
//	if len(ep.CCers) != 0 {
//		for _, tmp := range strings.Split(ep.CCers, ",") {
//			toers = append(toers, strings.TrimSpace(tmp))
//		}
//		m.SetHeader("Cc", toers...)
//	}
//
//	// 发件人
//	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
//	m.SetAddressHeader("From", fromEmail, "")
//
//	// 发送
//	err := d.DialAndSend(m)
//	if err != nil {
//		panic(err)
//	}
//}