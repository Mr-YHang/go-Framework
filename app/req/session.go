package req

import "errors"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (loginReq *LoginReq) Check() error {
	if loginReq.Username == "" {
		return errors.New("用户名不能为空")
	}
	if loginReq.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}
