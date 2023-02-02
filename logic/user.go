package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户存不存在
	var exists bool
	err = mysql.CheckUerExists(p.Username)
	if err != nil {
		//数据库查询出错
		return err
	}
	if exists {
		//用户已存在的错误
		return mysql.ErrorUserExist
	}

	//生成UID
	userID := snowflake.GenID()

	//构造一个User 实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(user)
}

// Login 登录
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	//生成JWT
	token, err = jwt.GenToken(user.UserID, user.Username)
	return token, err
}
