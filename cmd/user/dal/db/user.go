package db

import (
	"context"
	"errors"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateUser(ctx context.Context, user *User) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", user.Username).First(&userResp).Error

	if err == nil {
		return nil, errno.ParamError
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = DB.WithContext(ctx).Create(user).Error; err != nil {
		// add some logs
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", username).First(&userResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ParamError
		}
		return nil, err
	}

	return userResp, nil
}

func GetUserByID(ctx context.Context, userid int64) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("id = ?", userid).First(&userResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ParamError
		}
		return nil, err
	}

	return userResp, nil
}

func SaveUser(ctx context.Context, user *User) error {

	if err := DB.WithContext(ctx).Save(user).Error; err != nil {
		// add some logs
		return err
	}

	return nil
}
