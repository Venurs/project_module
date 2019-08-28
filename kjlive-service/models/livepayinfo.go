package models

import (
    "github.com/astaxie/beego/orm"
    "kjlive-service/redis"
    "strconv"
    "time"
    "fmt"
)


func IsPayByLiveChargeInfoId(liveChargeInfoId int, userId int) (isPay bool) {
    o := orm.NewOrm()
    sql := `
        SELECT
          Count(*) as count
        FROM live_livepayuserinfo T0
        WHERE T0.live_charge_info_id = ?
              AND T0.delete_flag = 0
              AND T0.user_id = ?
    `
    var count int
    err := o.Raw(sql, liveChargeInfoId, userId).QueryRow(&count)
    if err == nil && count > 0{
        return true
    } else {
        return false
    }
}

func PayCountLiveChargeInfoId(liveChargeInfoId int) (count int) {
    count = int(redis.GetInt("LIVE_CHARGE_INFO_PAY_COUNT_" +  strconv.Itoa(liveChargeInfoId)))
    if count == 0 {
        o := orm.NewOrm()
        sql := `
        SELECT
          Count(*) as count
        FROM live_livepayuserinfo T0
        WHERE T0.live_charge_info_id = ?
              AND T0.delete_flag = 0
        `
        err := o.Raw(sql, liveChargeInfoId).QueryRow(&count)
        if err == nil {
            redis.Set(fmt.Sprintf("LIVE_CHARGE_INFO_PAY_COUNT_%d", liveChargeInfoId), count, 60 * time.Second)
        }
    }
    return
}