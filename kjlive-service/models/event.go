package models

import (
    "github.com/astaxie/beego/orm"
    "time"
    "kjlive-service/redis"
    "strconv"
    "encoding/json"
    "kjlive-service/logs"
)

type BaseEvent struct {
    Id        int       `json:"id" orm:"column(id)"`
    Name      string    `json:"name" orm:"column(name)"`
    StartTime time.Time `json:"starttime" orm:"column(starttime)"`
    EndTime   time.Time `json:"endtime" orm:"column(endtime)"`
    SignUrl   string    `json:"sign_url" orm:"column(sign_url)"`
}

func GetEventInfoByLiveId(liveId int) (baseEvent BaseEvent) {
    jsonBaseEvent := redis.Get("LIVE_INFO_BASE_EVENT_" +  strconv.Itoa(liveId))
    if jsonBaseEvent == "" {
        o := orm.NewOrm()
        sql := `
        SELECT
          T1.id        id,
          T1.name      name,
          T1.starttime starttime,
          T1.endtime   endtime,
          T1.sign_url  sign_url
        FROM live_liveinfo T0
          LEFT JOIN event_baseevent T1
            ON T0.baseevent_id = T1.id
        WHERE T0.id = ?
              AND T0.delete_flag = 0
    `
        err := o.Raw(sql, liveId).QueryRow(&baseEvent)
        if err != nil {
            logs.Error(err)
        } else {
            if jsonBaseEvent, err := json.Marshal(baseEvent); err == nil {
                redis.Set("LIVE_INFO_BASE_EVENT_" +  strconv.Itoa(liveId), string(jsonBaseEvent), 60 * 12 * time.Second)
            } else {
                logs.Error(err)
            }
        }
    } else {
        json.Unmarshal([]byte(jsonBaseEvent), &baseEvent)
    }
    return
}
