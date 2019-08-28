package models

import (
    "time"
    "encoding/json"
    
    "github.com/astaxie/beego/orm"
    
    "kjlive-service/redis"
    "kjlive-service/logs"
)

type LiveChargePopSerializer struct {
    Id             int     `json:"id" orm:"column(id)"`
    IsUsed         int       `json:"is_used" orm:"column(is_used)"`
    StartTime      time.Time `json:"start_time" orm:"column(start_time)"`
    EndTime        time.Time `json:"end_time" orm:"column(end_time)"`
    DiscountAmount float64   `json:"discount_amount" orm:"column(discount_amount)"`
}

func GetLiveChargePopById(pop_param string) (liveChargePopSerializer LiveChargePopSerializer, err error) {
    jsonLivePop := redis.Get("LIVE_CHARGE_POP_" + pop_param)
    if jsonLivePop == "" {
        o := orm.NewOrm()
        sql := `
        SELECT
          T0.id                                  id,
          T0.discount_amount                     discount_amount,
          T0.is_used                             is_used,
          T0.start_time                          start_time,
          T0.end_time                            end_time
        FROM live_livechargepop T0
        WHERE T0.pop_param = ?
              AND T0.delete_flag = 0
        `
        err = o.Raw(sql, pop_param).QueryRow(&liveChargePopSerializer)
        if err != nil {
            return
        } else if liveChargePopSerializer.EndTime.After(time.Now()){
            subTime := liveChargePopSerializer.EndTime.Sub(time.Now())
            if jsonLivePopByte, err := json.Marshal(liveChargePopSerializer); err == nil {
                redis.Set("LIVE_CHARGE_POP_" +  pop_param, string(jsonLivePopByte), subTime * time.Second)
            } else {
                logs.Error(err)
            }
        }
    } else {
        json.Unmarshal([]byte(jsonLivePop), &liveChargePopSerializer)
    }
    
    return liveChargePopSerializer, nil
}
