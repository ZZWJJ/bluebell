package service

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/model"
	"bluebell/model/database"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"fmt"
	"strconv"
)

func SignUp(p *model.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成uid
	userID, _ := snowflake.GetID()
	user := &database.User{
		UserId:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *model.ParamLogin) (token string, err error) {
	u := database.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(&u, p.Password); err != nil {
		return "", err
	}
	token, err = jwt.GenToken(u.UserId, u.Username)
	if err != nil {
		return "", err
	}
	if err := redis.Set(strconv.FormatUint(u.UserId, 10), token); err != nil {
		fmt.Println("redis set failed: ", err)
	}
	return
}
