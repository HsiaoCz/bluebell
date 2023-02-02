package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type myClaims struct {
	// 可根据需要自行添加字段
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义过期时间
const TokenExpirDuration = time.Hour * 24 * 3

var mySecret = []byte("小樊宜")

// GenToken 生成JWT
func GenToken(userID int64, username string) (token string, err error) {
	// 创建一个我们自己的声明数据
	claims := myClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpirDuration).Unix(),
			Issuer:    "bluebell", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
	if err != nil {
		zap.L().Error("token failed", zap.Error(err))
		return
	}
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token, err
}

// ParseToken 解析JWT
// ParseToken 解析JWT
func ParseToken(tokenString string) (*myClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
