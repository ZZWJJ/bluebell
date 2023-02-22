package service

import (
	"bluebell/dao/mysql"
	"bluebell/model"
	"bluebell/model/database"
	"bluebell/pkg/snowflake"
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

func Login(p *model.ParamLogin) (err error) {
	u := database.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(&u, p.Password)
}
