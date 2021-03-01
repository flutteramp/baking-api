package user

import (
	"github.com/flutter-amp/baking-api/entity"
)

type UserService interface {
	User(id uint) (*entity.User, []error)
	UserByEmail(email string) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
	EmailExists(email string) bool
}
