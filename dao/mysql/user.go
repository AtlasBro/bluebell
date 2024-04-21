package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "liwenzhou.com"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorNotExist        = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlstr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlstr, user.UserID, user.UserName, user.Password)
	return
}

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login 登录
func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登录的密码
	sqlstr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlstr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库失败
		return err
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return ErrorNotExist
	}

	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return
	}
	return ErrorInvalidPassword
}
