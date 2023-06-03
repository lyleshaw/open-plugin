package utils

import (
	"crypto/rand"
	"github.com/bytedance/gopkg/util/logger"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
)

// 生成随机数验证码
func generateCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		b := make([]byte, 1)
		if _, err := rand.Read(b); err != nil {
			logger.Errorf("generate code error: %v", err)
			return ""
		}
		code += strconv.Itoa(int(b[0]) % 10)
	}
	return code
}

// 发送邮件
func SendEmail(email string) (string, error) {
	// 读取环境变量中的邮件配置
	user := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	strPort := os.Getenv("EMAIL_PORT")

	// 生成随机验证码
	code := generateCode()

	// 构建邮件内容
	content := "<html>\n<head>\n<link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css\">\n</head>\n<body>\n<div class=\"container\">\n<h1 class=\"display-4 text-center my-5\">欢迎注册</h1>  \n<div class=\"d-flex justify-content-center\">\n    <div class=\"col-sm-6\">\n        <p class=\"lead text-center\">您的注册验证码是:</p>\n        <h1 class=\"text-center mb-5\">" + code + "</h1> \n        <p class=\"text-center\">请不要将验证码透露给他人。</p>   \n    </div>\n</div>\n<p class=\"text-center\">感谢您的注册!</p>  \n</div>\n</body> \n</html>"
	m := mail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "MLOps 系统注册验证码")
	m.SetHeader("Content-Type", "text/html")
	m.SetBody("text/html", content)
	port, _ := strconv.Atoi(strPort)

	// 连接邮箱服务器并发送
	d := mail.NewDialer(host, port, user, password)
	return code, d.DialAndSend(m)
}

func main_() {
	email := "xlyleshaw@gmail.com"
	if _, err := SendEmail(email); err != nil {
		logger.Errorf("send email error: %v", err)
	} else {
		logger.Infof("send email success")
	}
}
