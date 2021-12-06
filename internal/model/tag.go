package model

import (
	"github.com/diy0663/goblog-service/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (Tag) TableName() string {
	return "blog_tag"
}

// tag.go 为了让swagger注解中的返回数据更贴合实际
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// sql查询,获取总数(传参放到自身结构体上面去了)
func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64

	db = getConditionDB(t, db)

	if err := db.Model(&t).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 查列表, 跟 查列表总数总是要用到的一致查询条件放这里
func getConditionDB(t Tag, db *gorm.DB) *gorm.DB {
	if t.Name != "" {
		db = db.Where("name = ? ", t.Name)
	}
	db = db.Where("state = ? ", t.State)
	db = db.Where("is_del = ? ", 0)
	return db
}

// todo 返回值为何是  []*Tag ???
func (t Tag) List(db *gorm.DB, pageOffset int, pageSize int) ([]*Tag, error) {
	var tags []*Tag

	var err error
	db = getConditionDB(t, db)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	// 获取多条数据用 Find
	if err = db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

// 插入一条数据
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error

}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	// 更新多列 ?  这种用法是否有效需要验证
	//	return db.Model(&t).Where("id = ? and is_del = ? ", t.ID, 0).Updates(t).Error
	// GORM 中使用 struct 类型传入进行更新时，GORM 是不会对值为零值的字段进行变更
	err := db.Model(t).Where("id =? and is_del = ? ", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	// 假如是真删除数据,那么  is_del 的意义何在.....
	return db.Where("id = ? and is_del = ? ", t.ID, 0).Delete(&t).Error
}
