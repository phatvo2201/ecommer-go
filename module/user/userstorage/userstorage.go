package userstorage

import "gorm.io/gorm"

type sqlStore struct {
	userDb *gorm.DB
}

func NewUserStore(userDb *gorm.DB) *sqlStore {
	return &sqlStore{userDb: userDb}

}
