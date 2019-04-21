package repository

import (
	"github.com/go_gin_sample/apps/configure/db"
	l "github.com/go_gin_sample/apps/configure/logger"
	m "github.com/go_gin_sample/apps/domain/model"
	"github.com/jinzhu/gorm"
	"time"
)
var logger  = l.GetLogger()

type UserRepository struct {}

// IDを元にユーザーを取得します
func (c *UserRepository) FindByUserId(id uint64) (*m.User, error) {

	var user = m.User{}
	user.Id = id
	db := db.GetDB()
	if err := db.First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Error("ユーザーの取得処理でエラーが発生しました: ", err)
			return nil, err
		} else {
			logger.Info("ユーザーが存在しません。ID: ", id)
			return nil, nil
		}
	}

	return &user, nil
}

// Emailを元にユーザーの取得を行います。
func (c *UserRepository) FindByEmail(email string) (*m.User, error) {

	var user = m.User{}
	user.Email = email
	db := db.GetDB().Where("email = ?", email)

	// Emailを元にユーザー取得
	if err := db.First(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.Error("ユーザーの取得処理でエラーが発生しました: ", err)
			return nil, err
		} else {
			logger.Info("ユーザーが存在しません。: ", email)
			return nil, nil
		}
	}

	return &user, nil
}

// ユーザー情報を作成します
func (c *UserRepository) Save(user *m.User) (*m.User, error) {

	db := db.GetDB()
	if err := db.Create(&user).Error; err != nil {
		logger.Error("ユーザーの作成処理でエラーが発生しました: ", err)
		return nil, err
	}

	return user, nil
}

// ユーザー情報を更新します
func (c *UserRepository) Update(user *m.User) (*m.User, error) {

	db := db.GetDB()
	if err := db.Save(&user).Error; err != nil {
		logger.Error("ユーザーの更新処理でエラーが発生しました: ", err)
		return nil, err
	}

	return user, nil
}

// ユーザー情報を削除します
func (c *UserRepository) Delete(user *m.User) (*m.User, error) {

	db := db.GetDB()
	// 論理削除
	user.Status = m.UserStatusInValid
	current := time.Now()
	user.DeletedAt = &current

	if err := db.Save(&user).Error; err != nil {
		logger.Error("ユーザーの削除処理でエラーが発生しました: ", err)
		return nil, err
	}

	return user, nil
}