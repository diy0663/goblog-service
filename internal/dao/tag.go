package dao

import (
	"github.com/diy0663/goblog-service/internal/model"
	"github.com/diy0663/goblog-service/pkg/app"
)

// dao 这一层, 把 db 和已赋值 的结构体一起使用, 去调用model方法

func (d *Dao) CountTag(name string, state uint8) (int64, error) {

	// 为结构体赋值
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	// 去调用数据库操作方法,结构体以及在上面赋值了,参数传的其实是 *grom.DB 类型
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint64, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	}

	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint64) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

//  操作路径:  路由->控制器->service-> dao->model SQL操作
