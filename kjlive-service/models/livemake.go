package models

import "github.com/astaxie/beego/orm"

func IsMakeLive(liveId int, userId int) bool {
    o := orm.NewOrm()
    sql := `
        SELECT
          Count(*) as count
        FROM live_livemake T0
        WHERE T0.live_info_id = ?
              AND T0.delete_flag = 0
              AND T0.user_id = ?
    `
    var count int
    err := o.Raw(sql, liveId, userId).QueryRow(&count)
    if err == nil && count > 0{
        return true
    } else {
        return false
    }
}
func GetMakeNumberByLiveId(liveId int) int {
    o := orm.NewOrm()
    sql := `
        SELECT
          Count(*) as count
        FROM live_livemake T0
        WHERE T0.live_info_id = ?
              AND T0.delete_flag = 0
    `
    var count int
    err := o.Raw(sql, liveId).QueryRow(&count)
    if err == nil && count > 0 {
        return count
    } else {
        return 0
    }
}
