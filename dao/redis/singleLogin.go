package redis

import "errors"

// SingleLogin 单点登录
func SingleLogin(oToken, userID string) (err error) {
	var redisToken string
	redisToken, err = rdb.Get(rdb.Context(), userID).Result()
	if err != nil {
		return
	}
	if redisToken != oToken {
		return errors.New("redis token not equal oToken")
	}
	return nil
}
