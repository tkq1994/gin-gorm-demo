package dao

import (
	"gorm.io/gorm"
	"iTypeService/global"
)

type BaseDao struct {
	Orm *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		Orm: global.DB,
	}
}
