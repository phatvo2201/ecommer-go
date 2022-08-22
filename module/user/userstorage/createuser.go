package userstorage

import (
	"context"
	"food-delivery-service/common"
	"food-delivery-service/module/user/usermodel"
)

func (s sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.userDb.Begin()
	data.PrepareForInsert()
	if err := db.Table(data.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)

	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)

	}
	return nil
}
