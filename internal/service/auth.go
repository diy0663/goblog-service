package service

import (
	"errors"

	"github.com/diy0663/goblog-service/internal/requests"
)

// service 去调用dao
func (svc *Service) CheckAuth(param *requests.AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
