package service

import (
	"github.com/west2-online/fzuhelper-server/cmd/user/dal/db"
	"github.com/west2-online/fzuhelper-server/cmd/user/pack"
	"github.com/west2-online/fzuhelper-server/kitex_gen/user"
)

// GetUser check token and get user's info
func (s *UserService) GetUser(req *user.InfoRequest) (*user.User, error) {
	var userResp *user.User
	// 获取用户基本信息
	userModel, err := db.GetUserByID(s.ctx, req.UserId)
	userResp = pack.User(userModel)

	if err != nil {
		return nil, err
	}

	return userResp, nil
}
