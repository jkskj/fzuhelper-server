package service

import (
	"github.com/west2-online/fzuhelper-server/cmd/user/dal/db"
	"github.com/west2-online/fzuhelper-server/kitex_gen/user"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
	"golang.org/x/crypto/bcrypt"
)

// CheckUser check user is existed and it's password
func (s *UserService) CheckUser(req *user.LoginRequest) (*db.User, error) {
	userModel, err := db.GetUserByUsername(s.ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)) != nil {
		return nil, errno.AuthFailedError
	}

	return userModel, nil
}
