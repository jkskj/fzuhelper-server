package pack

import (
	"errors"
	"github.com/west2-online/fzuhelper-server/kitex_gen/user"
	"github.com/west2-online/fzuhelper-server/pkg/errno"
)

func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.InternalServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
