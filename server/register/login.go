package register

import (
	"codies-server/server/authorize"
	"codies-server/server/user"
)

func loginByEmail(email, pwd string) (*user.User, error) {
	uuid, err := authorize.CheckAuthByEmail(email, pwd)
	if err != nil {
		return nil, err
	}
	user, err := user.UserByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
