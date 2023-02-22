package mysql

import (
	"bluebell/model/database"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const (
	secret = "zzw.com"
)

func CheckUserExist(username string) (err error) {
	// 执行SQL语句入库
	sqlStr := `SELECT count(user_id) FROM user WHERE username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}
	return
}

// InsertUser 插入用户数据
func InsertUser(user *database.User) (err error) {
	// 对密码进行加密
	// 执行SQL语句入库
	sqlStr := `INSERT INTO user (user_id,username, password) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, encryptPassword(user.Password))
	return err
}

// Login 用户登录
func Login(user *database.User, oPassword string) (err error) {
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	if user.Password != encryptPassword(oPassword) {
		return errors.New("用户名密码不正确")
	}
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
