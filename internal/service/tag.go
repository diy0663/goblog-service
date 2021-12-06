package service

import (
	"github.com/diy0663/goblog-service/internal/model"
	"github.com/diy0663/goblog-service/internal/requests"
	"github.com/diy0663/goblog-service/pkg/app"
)

// 路由->控制器->service-> dao->model
func (svc *Service) CountTag(param *requests.CountTagRequest) (int64, error) {
	return svc.dao.CountTag(param.Name, param.State)
}
func (svc *Service) GetTagList(param *requests.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *requests.CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *requests.UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *requests.DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
