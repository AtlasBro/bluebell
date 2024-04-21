package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个user实例
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (aToken, rToken string, err error) {
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID和user.Username
	if err := mysql.Login(user); err != nil {
		return "", "", nil
	}
	// 生成JWT
	return jwt.GenToken(user.UserID, user.UserName)
}
