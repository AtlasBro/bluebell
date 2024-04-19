package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// CtxUserIDKey 定义上下文key常量
const CtxUserIDKey = "userID"

// ErrorUserNotLogin 定义常量错误信息
var ErrorUserNotLogin = errors.New("用户未登录")

// 获取当前登录的用户ID
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}
