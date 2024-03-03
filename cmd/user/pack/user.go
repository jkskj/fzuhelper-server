package pack

import (
	"github.com/west2-online/fzuhelper-server/cmd/user/dal/db"
	"github.com/west2-online/fzuhelper-server/kitex_gen/user"
)

func User(data *db.User) *user.User {
	if data == nil {
		return nil
	}

	return &user.User{
		Id:   data.Id,
		Name: data.Username,
	}

}
