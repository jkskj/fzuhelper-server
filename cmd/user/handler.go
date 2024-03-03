package main

import (
	"context"
	"github.com/west2-online/fzuhelper-server/cmd/user/pack"
	"github.com/west2-online/fzuhelper-server/cmd/user/service"
	user "github.com/west2-online/fzuhelper-server/kitex_gen/user"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
	"github.com/west2-online/fzuhelper-server/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).CreateUser(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	token, err := utils.CreateToken(userResp.Id)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.UserId = userResp.Id
	resp.Token = token
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).CheckUser(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	token, err := utils.CreateToken(userResp.Id)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.User = pack.User(userResp)
	resp.Token = token
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	resp = new(user.InfoResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthFailedError)
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).GetUser(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.User = userResp
	return
}

// Change implements the UserServiceImpl interface.
func (s *UserServiceImpl) Change(ctx context.Context, req *user.ChangeRequest) (resp *user.ChangeResponse, err error) {
	resp = new(user.ChangeResponse)

	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	err = service.NewUserService(ctx).ChangePassword(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}
