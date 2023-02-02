package controllors

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxtUserIDKey = "userID"

var ErrorUserLogin = errors.New("用户未登录")

// 获取当前登录的ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxtUserIDKey)
	if !ok {
		err = ErrorUserLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserLogin
		return
	}
	return
}
