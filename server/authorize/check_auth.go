package authorize

import "codies-server/skeleton/common"

func CheckAuthByEmail(email, pwd string) (uuid string, err error) {
	localAuth := &LocalAuth{}
	if err := localAuth.GetByEmail(email); err != nil {
		return "", err
	}
	if password(localAuth.UUID, pwd) == localAuth.Password {
		return localAuth.UUID, nil
	}
	return "", common.InvalidArgumentErr("invalid password")
}
