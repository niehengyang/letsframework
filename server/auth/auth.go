package auth

import "golang.org/x/crypto/bcrypt"

const (
	//无效的Token
	ExpireOrErrorToken = 40100
	//用户名或密码错误
	ErrorLoginOrPassword = 40101
	//用户被禁用
	UserIsDisabled = 40102
)

var AuthErrorMap = map[int]string{
	ExpireOrErrorToken:   "无效的Token",
	ErrorLoginOrPassword: "用户名或密码错误",
	UserIsDisabled:       "用户被禁用",
}

// 生成加密密码
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// 对比密码
func ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
