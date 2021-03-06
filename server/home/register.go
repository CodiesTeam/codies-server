package home

import (
	"codies-server/server/authorize"
	"codies-server/server/user"
)

func regByEmail(username, email, pwd string) (*user.User, error) {
	// TODO: check wether email has registered
	// generate user, and insert to user table
	u := user.NewUser(username)
	u.Email = email
	// TODO: use rollback
	err := u.Insert()
	if err != nil {
		return nil, err
	}
	localAuth := authorize.NewLocalAuth(u.UUID, email, "", pwd)
	err = localAuth.Insert()
	if err != nil {
		return nil, err
	}
	return u, nil
}
