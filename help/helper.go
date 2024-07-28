package help

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/smtp"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("gingormkey")

func GenerateJwt(identity string, name string) (string, error) {
	userClaim := UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		log.Panic(err)
		return "", err
	}
	return tokenString, nil
}
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("analyse token fail", err)
	}
	if !claims.Valid {
		return nil, err
	}
	return userClaim, nil
}

func SendCode(toUserEmail string, code string) error {
	e := email.NewEmail()
	e.From = "yangchen<yangchenyc@mail.ustc.edu.cn>"
	e.To = []string{
		toUserEmail,
	}

	e.Subject = "验证码发送测试"

	e.HTML = []byte("您的验证码:<b>" + code + "</b>")
	return e.SendWithTLS("mail.ustc.edu.cn:465", smtp.PlainAuth("", "yangchenyc@mail.ustc.edu.cn", "qFTKVz3AUGKk5Zb6", "mail.ustc.edu.cn"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "mail.ustc.edu.cn"})

}
func GetUUID() string {
	return uuid.NewV4().String()
}
