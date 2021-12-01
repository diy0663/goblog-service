package model

// 创建模型基类 把共用字段放这里
type Model struct {
	ID         uint64 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  uint64 `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint64 `json:"modified_on"`
	DeletedOn  uint64 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
