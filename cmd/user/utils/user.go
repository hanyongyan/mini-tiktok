package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"

	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
)

func CheckUser(q *query.Query, username, password string) (res *model.TUser, err error) {
	tuser := q.TUser
	res, _ = tuser.WithContext(context.Background()).
		Where(tuser.Name.Eq(username)).
		First()
	if res == nil {
		err = fmt.Errorf("用户不存在: %v", username)
		return
	}
	if pwd := ScryptPwd(password); pwd != res.Password {
		err = fmt.Errorf("密码错误: %v", password)
		return
	}
	return
}

// ScryptPwd 加密
func ScryptPwd(password string) string {
	const PwdHashByte = 10
	salt := make([]byte, 8)
	salt = []byte{200, 20, 9, 29, 15, 50, 80, 7}

	key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, PwdHashByte)
	if err != nil {
		log.Fatal(err)
	}
	FinPwd := base64.StdEncoding.EncodeToString(key)
	return FinPwd
}
