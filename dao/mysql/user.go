package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
)

// 把每一步数据库操作封装成函数
// 待logic 层根据业务需求调用

const secret = "xiaofanyi"

var (
	ErrorUserExist      = errors.New("用户已存在")
	ErrorUserNotExist   = errors.New("用户不存在")
	ErrorPasswordFailed = errors.New("密码错误")
	ErrorInvalidID      = errors.New("无效的ID")
)

// CheckUserExists 判断用户是否存在

func CheckUerExists(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser  向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	if err != nil {
		zap.L().Error("insert into user failed")
		return
	}
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	//用户登录时的密码
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorPasswordFailed
	}
	return err
}

// GetPostUser 根据用户id获取用户详情
func GetPostUser(uid int64) (user models.User, err error) {
	sqlStr := `select user_id,username from user where user_id=?`
	err = db.Get(&user, sqlStr, uid)
	return
}
