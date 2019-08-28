package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "kjlive-service/redis"
    "strconv"
    "encoding/json"
    
    "kjlive-service/logs"
)

type LiveChargeSerializer struct {
    Id                   int     `json:"id" orm:"column(id)"`
    Name                 string  `json:"name" orm:"column(name)"`
    Summary              string  `json:"summary" orm:"column(summary)"`
    OriginalPrice        float64 `json:"original_price" orm:"column(original_price)"`
    PresentPrice         float64 `json:"present_price" orm:"column(present_price)"`
    Sale                 float64 `json:"sale" orm:"column(sale)"`
    SmallImageUrl        string  `json:"small_image_url" orm:"column(small_image_url)"`
    StartTime            string  `json:"start_time" orm:"column(start_time)"`
    EndTime              string  `json:"end_time" orm:"column(end_time)"`
    LiveStartTime        string  `json:"live_start_time" orm:"column(live_start_time)"`
    IsDiscount           int8    `json:"is_discount" orm:"column(is_discount)"`
    EnrollCountIncrement int     `json:"enroll_count_increment" orm:"column(enroll_count_increment)"`
}

func GetLiveChargeInfoById(id int) (liveChargeSerializer LiveChargeSerializer, err error) {
    jsonLiveInfo := redis.Get("LIVE_CHARGE_INFO_" + strconv.Itoa(id))
    if jsonLiveInfo == "" {
        o := orm.NewOrm()
        sql := `
            SELECT
              T0.id                                                id,
              T0.name                                              name,
              T0.summary                                           summary,
              T0.original_price                                    original_price,
              T0.present_price                                     present_price,
              T0.sale                                              sale,
              T0.small_image_url                                   small_image_url,
              DATE_FORMAT(T0.start_time, '%Y-%m-%d %H:%i:%s')      start_time,
              DATE_FORMAT(T0.end_time, '%Y-%m-%d %H:%i:%s')        end_time,
              DATE_FORMAT(T0.live_start_time, '%Y-%m-%d %H:%i:%s') live_start_time,
              T0.is_discount                                       is_discount,
              T0.enroll_count_increment                            enroll_count_increment
            FROM live_livechargeinfo T0
            WHERE T0.live_info_id = ?
                  AND T0.delete_flag = 0
                  AND T0.start_time <= ?
                  AND T0.end_time >= ?
            `
        _ = o.Raw(sql, id, time.Now(), time.Now()).QueryRow(&liveChargeSerializer)
        if liveChargeSerializer.Id != 0 {
            loc, _ := time.LoadLocation("Local")
            theTime, err := time.ParseInLocation("2006-01-02 15:04:05", liveChargeSerializer.EndTime, loc)
            var subTime time.Duration
            if err != nil {
                logs.Error(err)
                subTime = 60
            } else {
                subTime = theTime.Sub(time.Now())
            }
            if jsonLiveInfo, err := json.Marshal(liveChargeSerializer); err == nil {
                redis.Set("LIVE_CHARGE_INFO_"+strconv.Itoa(id), string(jsonLiveInfo), subTime*time.Second)
            } else {
                logs.Error(err)
            }
        } else {
            if jsonLiveInfo, err := json.Marshal(liveChargeSerializer); err == nil {
                redis.Set("LIVE_CHARGE_INFO_"+strconv.Itoa(id), string(jsonLiveInfo), 60*time.Second)
            } else {
                logs.Error(err)
            }
        }
    } else {
        json.Unmarshal([]byte(jsonLiveInfo), &liveChargeSerializer)
    }
    
    return liveChargeSerializer, nil
}

type LiveChargeInfoSerializer struct {
    Id                   int     `json:"id" orm:"column(id)"`
    Name                 string  `json:"name" orm:"column(name)"`
    Summary              string  `json:"summary" orm:"column(summary)"`
    OriginalPrice        float64 `json:"original_price" orm:"column(original_price)"`
    PresentPrice         float64 `json:"present_price" orm:"column(present_price)"`
    Sale                 float64 `json:"sale" orm:"column(sale)"`
    SmallImageUrl        string  `json:"small_image_url" orm:"column(small_image_url)"`
    StartTime            string  `json:"start_time" orm:"column(start_time)"`
    EndTime              string  `json:"end_time" orm:"column(end_time)"`
    LiveStartTime        string  `json:"live_start_time" orm:"column(live_start_time)"`
    IsDiscount           int8    `json:"is_discount" orm:"column(is_discount)"`
    EnrollCountIncrement int     `json:"enroll_count_increment" orm:"column(enroll_count_increment)"`
    LecturerId           int     `json:"lecturer" orm:"column(lecturer)"`
    LecturerName         string  `json:"lecturer_name" orm:"column(lecturer_name)"`
    LecturerHonour       string  `json:"lecturer_honour" orm:"column(lecturer_honour)"`
    LecturerPictureUrl   string  `json:"lecturer_picture_url" orm:"column(lecturer_picture_url)"`
    LecturerSummary      string  `json:"lecturer_summary" orm:"column(lecturer_summary)"`
    LikedCount           int     `json:"liked_count" orm:"column(liked_count)"`
    WechatShareImageUrl  string  `json:"wechat_share_image_url" orm:"column(wechat_share_image_url)"`
    WechatShareSummary   string  `json:"wechat_share_summary" orm:"column(wechat_share_summary)"`
    RoomName             string  `json:"room_name" orm:"column(room_name)"`
    EnrollStartTime      string  `json:"enroll_start_time" orm:"column(enroll_start_time)"`
    EnrollEndTime        string  `json:"enroll_end_time" orm:"column(enroll_end_time)"`
    FreeIsShow           int8    `json:"free_is_show" orm:"column(free_is_show)"`
    FreeImageUrl         string  `json:"free_image_url" orm:"column(free_image_url)"`
    IsCharge             int     `json:"is_charge"`
    ShowPrimeExchange    int     `json:"show_prime_exchange"`
    IsPay                int     `json:"is_pay"`
    PageView             int     `json:"page_view"`
    IsUsedUp             int     `json:"is_used_up"`
    PreferPrice          float64 `json:"prefer_price"`
    LiveSeriesId         int     `json:"live_series_id" orm:"column(live_series_id)"`
    LiveInfoId           int     `json:"live_info_id" orm:"column(live_info_id)"`
    EnrollCount          int     `json:"enroll_count"`
    WatchCountIncrement  int     `json:"watch_count_increment" orm:"column(watch_count_increment)"`
}

func GetLiveLiveChargeInfoByLiveId(liveId int) (liveChargeInfoSerializer LiveChargeInfoSerializer, err error) {
    jsonLiveInfo := redis.Get("LIVE_CHARGE_PAGE_INFO_" + strconv.Itoa(liveId))
    if jsonLiveInfo == "" {
        o := orm.NewOrm()
        sql := `
            SELECT
              T0.id                                                id,
              CASE T0.name
              WHEN NULL
                THEN T1.name
              ELSE T0.name END                                     name,
              CASE T0.summary
              WHEN ''
                THEN T1.summary
              ELSE T0.summary END                                  summary,
              T1.lecturer_id                                       lecturer,
              CASE T1.lecturer_name
              WHEN ''
                THEN T2.lecturer_name
              ELSE T1.lecturer_name END                            lecturer_name,
              CASE T1.lecturer_summary
              WHEN NULL
                THEN T2.lecturer_summary
              ELSE T1.lecturer_summary END                         lecturer_summary,
              CASE T1.lecturer_picture_url
              WHEN NULL
                THEN T2.lecturer_picture_url
              ELSE T1.lecturer_picture_url END                     lecturer_picture_url,
              CASE T1.lecturer_honour
              WHEN NULL
                THEN T2.lecturer_honour
              ELSE T1.lecturer_honour END                          lecturer_honour,
              DATE_FORMAT(T1.start_time, '%Y-%m-%d %H:%i:%s')      start_time,
              DATE_FORMAT(T0.end_time, '%Y-%m-%d %H:%i:%s')        end_time,
              T1.liked_count                                       liked_count,
              T1.wechat_share_image_url                            wechat_share_image_url,
              T1.wechat_share_summary                              wechat_share_summary,
              T1.room_name                                         room_name,
              T0.original_price                                    original_price,
              T0.present_price                                     present_price,
              T0.sale                                              sale,
              T0.small_image_url                                   small_image_url,
              DATE_FORMAT(T0.start_time, '%Y-%m-%d %H:%i:%s')      enroll_start_time,
              DATE_FORMAT(T0.end_time, '%Y-%m-%d %H:%i:%s')        enroll_end_time,
              DATE_FORMAT(T0.live_start_time, '%Y-%m-%d %H:%i:%s') live_start_time,
              T0.enroll_count_increment                            enroll_count_increment,
              T3.free_is_show                                      free_is_show,
              T3.free_image_url                                    free_image_url,
              T0.live_series_id                                    live_series_id,
              T0.live_info_id                                      live_info_id,
              T1.watch_count_increment                             watch_count_increment
            FROM live_livechargeinfo T0
              LEFT JOIN live_liveinfo T1 ON T0.live_info_id = T1.id
              LEFT JOIN video_lecturer T2 ON T1.lecturer_id = T2.id
              LEFT JOIN live_liveadinfo T3 ON T1.id = T3.live_info_id
            WHERE T0.live_info_id = ?
                  AND T0.delete_flag = 0
                  AND T0.start_time <= ?
                  AND T0.end_time >= ?
            ORDER BY ID DESC
            `
        
        err = o.Raw(sql, liveId, time.Now(), time.Now()).QueryRow(&liveChargeInfoSerializer)
        if err == nil {
            loc, _ := time.LoadLocation("Local")
            theTime, err := time.ParseInLocation("2006-01-02 15:04:05", liveChargeInfoSerializer.EndTime, loc)
            var subTime time.Duration
            if err != nil {
                logs.Error(err)
                subTime = 60
            } else {
                subTime = theTime.Sub(time.Now())
            }
            if jsonLiveInfo, err := json.Marshal(liveChargeInfoSerializer); err == nil {
                redis.Set("LIVE_CHARGE_PAGE_INFO_"+strconv.Itoa(liveId), string(jsonLiveInfo), subTime*time.Second)
            } else {
                logs.Error(err)
            }
        }
    } else {
        json.Unmarshal([]byte(jsonLiveInfo), &liveChargeInfoSerializer)
    }
    liveChargeInfoSerializer.setField()
    return
}

func (liveChargeInfoSerializer *LiveChargeInfoSerializer) setField() {
    if liveChargeInfoSerializer.PresentPrice == 0 {
        liveChargeInfoSerializer.IsCharge = 0
    } else {
        liveChargeInfoSerializer.IsCharge = 1
    }
    if liveChargeInfoSerializer.PresentPrice > 0 && liveChargeInfoSerializer.LiveSeriesId == 0 {
        liveChargeInfoSerializer.ShowPrimeExchange = 1
    }
    count := PayCountLiveChargeInfoId(liveChargeInfoSerializer.Id)
    liveChargeInfoSerializer.EnrollCount = liveChargeInfoSerializer.EnrollCountIncrement + count

    liveChargeInfoSerializer.PageView = GetLiveAccessNumber(liveChargeInfoSerializer.LiveInfoId, false) + liveChargeInfoSerializer.WatchCountIncrement
}
