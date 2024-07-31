package main

import (
	"golang.org/x/crypto/bcrypt"
	"k8s.io/klog/v2"
	"log"
)

// BcryptEncode 加密密码
func BcryptEncode(str string) string {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		klog.Info("调用bcrypt,加密异常，HASH编码失败:%w", err)
	}
	return string(bcryptPassword)
}

func VerifyPassword(sourcePwd, bcryptPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptPwd), []byte(sourcePwd))
	if err != nil {
		return false
	}
	return true
}

func GetHashingCost(hashedPassword []byte) int {
	cost, _ := bcrypt.Cost(hashedPassword) // 为了简单忽略错误处理
	return cost
}

func main() {
	pwd := "go-by-example"
	pwdEnc := BcryptEncode(pwd)
	log.Println(VerifyPassword(pwd, pwdEnc))

	log.Println(GetHashingCost([]byte(pwdEnc)))
}
