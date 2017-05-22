package service

import (
	"model"

	daoApi "api/dao_service"
)

type ActionService struct {
}

func (this *ActionService) Create(action *model.Action) error {
	var err error

	err = daoApi.ActionDaoApi.Create(action)
	if err != nil {
		return err
	}

	return err
}

func (this *ActionService) GetAll(userId int64) ([]*model.Action, error) {
	var err error
	var actions []*model.Action

	actions, err = daoApi.ActionDaoApi.GetAll(userId)
	if err != nil {
		return nil, err
	}

	return actions, err
}
