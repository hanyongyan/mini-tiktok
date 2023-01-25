package utils

import (
	"context"
	"fmt"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
)

func CheckUser(q *query.Query, username, password string) (res *model.TUser, err error) {
	tuser := q.TUser
	res, err = tuser.WithContext(context.Background()).
		Where(tuser.Name.Eq(username)).
		First()
	if res == nil {
		err = fmt.Errorf("用户不存在: %v", username)
		return
	}
	if res.Password != password {
		err = fmt.Errorf("密码错误: %v", password)
		return
	}
	return
}
