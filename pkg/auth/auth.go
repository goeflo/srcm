package auth

import (
	"errors"

	srcm_db "github.com/floriwan/srcm/pkg/db"
	"github.com/floriwan/srcm/pkg/db/model"
)

var WrongPasswdError = errors.New("password does not match")
var UserNotExistsError = errors.New("error does not exists")

func ValidateUser(email string, passwd string) error {

	u := model.User{Email: email}
	res := srcm_db.Instance.First(&u)
	if res.Error != nil {
		return UserNotExistsError
	}

	if err := u.CheckPassword(passwd); err != nil {
		return WrongPasswdError
	}

	return nil
}
