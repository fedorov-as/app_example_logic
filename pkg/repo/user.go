package repo

import (
	"github.com/fedorov-as/app_example_logic/pkg/model"
)

type UsersRepo interface {
	AddUser(user model.User) error
	GetUser(nickname string) (model.User, error)
}
