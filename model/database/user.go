package database

type User struct {
	UserId   uint64 `db:"user_id,string"` // 在结构体tag中添加string来告诉json包从字符串中解析相应字段的数据
	Username string `db:"username"`
	Password string `db:"password"`
}
