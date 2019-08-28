package models

import (
	"github.com/astaxie/beego/orm"
)
//
//type BaseContenttypeuser struct {
//	Id             int       `orm:"column(id);auto"`
//	DeleteFlag     int8      `orm:"column(delete_flag)"`
//	CreateTime     time.Time `orm:"column(create_time);type(datetime);null"`
//	DeleteTime     time.Time `orm:"column(delete_time);type(datetime);null"`
//	LastUpdateTime time.Time `orm:"column(last_update_time);type(datetime);null"`
//	AppId          int       `orm:"column(app_id)"`
//	ObjectId       int       `orm:"column(object_id)"`
//	Type           int       `orm:"column(type)"`
//	ContentTypeId  int       `orm:"column(content_type_id)"`
//	UserId         int       `orm:"column(user_id)"`
//}


func GetContenttypeuserByIdList(idList [3]int, userId int) (objectList []int) {
	o := orm.NewOrm()
	sql := `
		SELECT
		  object_id
		FROM base_contenttypeuser
		WHERE object_id IN (?,?,?)
			  AND type = 3
			  AND delete_flag = 0
			  AND user_id=?
	`
	
	num, err := o.Raw(sql, idList, userId).QueryRows(&objectList)
	if err == nil && num > 0{
		return objectList
	}
	return
}