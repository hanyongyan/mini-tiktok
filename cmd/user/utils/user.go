package utils

import (
	"context"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
)

func CheckUser(q *query.Query, username, passwrod string) (*model.TUser, error) {
	tuser := q.TUser
	return tuser.WithContext(context.Background()).
		Where(tuser.Name.Eq(username), tuser.Password.Eq(passwrod)).
		First()
}
