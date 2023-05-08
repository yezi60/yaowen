package handler

import (
	"fmt"
	"gorm.io/gorm"
	"soleaf.xyz/yaowen/global"
	"soleaf.xyz/yaowen/model"
)

// Paginate 分页函数
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		// 从下标 offset 拿 pageSize条数据
		return db.Offset(offset).Limit(pageSize)
	}
}


func GetList(pn,ps int)[]model.Data{

	var list []model.Data

	global.DB.Scopes(Paginate(pn, ps)).Order("health desc").Find(&list)

	return list
}

func Save(item model.Data){

	res := global.DB.Save(&item)

	if res.Error != nil{
		fmt.Println(res.Error)
	}

}