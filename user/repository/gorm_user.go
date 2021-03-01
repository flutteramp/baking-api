package repository

import (
	"github.com/flutter-amp/baking-api/entity"
	"github.com/flutter-amp/baking-api/user"

	"github.com/jinzhu/gorm"
)

type UserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates a new object of UserGormRepo
func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

func (userRepo *UserGormRepo) User(id uint) (*entity.User, []error) {
	user := entity.User{}
	errs := userRepo.conn.First(&user, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

// UserByEmail retrieves a user by its email address from the database
func (userRepo *UserGormRepo) UserByEmail(email string) (*entity.User, []error) {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

//UpdateUser updates a given user in the database
// func (userRepo *UserGormRepo) UpdateUserAmount(user *entity.User, Amount uint) *entity.User {
// 	usr := user
// 	userRepo.conn.entity(&usr).UpdateColumn("amount", Amount)

// 	return usr
// }

// DeleteUser deletes a given user from the database
func (userRepo *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = userRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (userRepo *UserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := userRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser stores a new user into the database
func (userRepo *UserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := userRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
func (userRepo *UserGormRepo) EmailExists(email string) bool {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}
