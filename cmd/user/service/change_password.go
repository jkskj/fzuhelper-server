package service

import (
	"github.com/west2-online/fzuhelper-server/cmd/user/dal/db"
	"github.com/west2-online/fzuhelper-server/kitex_gen/user"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
	"golang.org/x/crypto/bcrypt"
)

// ChangePassword change user password
func (s *UserService) ChangePassword(req *user.ChangeRequest) error {
	userModel, err := db.GetUserByUsername(s.ctx, req.Username)

	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)) != nil {
		return errno.AuthFailedError
	}

	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword_), bcrypt.DefaultCost)

	userModel.Password = string(hashBytes)

	return db.SaveUser(s.ctx, userModel)
}
