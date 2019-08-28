package models

import (
    "github.com/astaxie/beego/orm"
    "reflect"
    "strings"
    "kjlive-service/redis"
    "strconv"
    "time"
)

type LiveAdInfo struct {
    MsgSkipUrl  string `json:"msg_skip_url" orm:"column(msg_skip_url)"`
    MsgImageUrl string `json:"msg_image_url" orm:"column(msg_image_url)"`
    HasPoster   string `json:"has_poster" orm:"column(has_poster)"`
    Links       map[string]string `json:"links"`
    
}

func GetLivePoster(liveId int) (liveAdInfo LiveAdInfo) {
    jsonLivePoster := redis.GetMap("LIVE_POSTER_INFO_" +  strconv.Itoa(liveId))
    if jsonLivePoster == nil {
        o := orm.NewOrm()
        sql := `
            SELECT
                msg_skip_url,
                msg_image_url,
                has_poster
            FROM live_liveadinfo
            WHERE live_info_id = ?
                  AND (has_poster IS NOT NULL
                        OR msg_image_url IS NOT NULL
                        OR msg_skip_url IS NOT NULL)
        `
        err := o.Raw(sql, liveId).QueryRow(&liveAdInfo)
        if err == nil {
            if liveAdInfo.HasPoster == "1" {
                if liveAdInfo.MsgImageUrl == "" || liveAdInfo.MsgSkipUrl == "" {
                    liveAdInfo.setLinks()
                }
            } else {
                liveAdInfo.setLinks()
            }
        } else {
            liveAdInfo.setLinks()
        }
        livePosterInfo := getMapLiveAdInfo(liveAdInfo)
        redis.Set("LIVE_POSTER_INFO_" +  strconv.Itoa(liveId), livePosterInfo, 24*time.Hour)
    } else {
        liveAdInfo.mapTransStruct(jsonLivePoster)
    }
    return
}

func getMapLiveAdInfo(serializer LiveAdInfo) map[string]interface{}{
    t := reflect.TypeOf(serializer)
    v := reflect.ValueOf(serializer)
    
    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[strings.ToLower(t.Field(i).Tag.Get("json"))] = v.Field(i).Interface()
    }
    return data
}

func (liveAdInfo *LiveAdInfo) mapTransStruct(jsonLivePoster map[string]interface{})  {
    liveAdInfo.MsgSkipUrl = jsonLivePoster["msg_skip_url"].(string)
    liveAdInfo.MsgImageUrl = jsonLivePoster["msg_image_url"].(string)
    liveAdInfo.HasPoster = jsonLivePoster["has_poster"].(string)
    mapJsonLivePoster := jsonLivePoster["links"].(map[interface{}]interface{})
    stringJsonLivePoster := make(map[string]string)
    if mapJsonLivePoster != nil {
        for k, v := range mapJsonLivePoster {
            stringJsonLivePoster[k.(string)] = v.(string)
        }
        liveAdInfo.Links = stringJsonLivePoster
    }
}


func (liveAdInfo *LiveAdInfo) setLinks()  {
    parameters := GetParameter("CMS_LIVE_POSTER_DICT")
    if parameters != nil {
        liveAdInfo.HasPoster = parameters["has_poster"].(string)
        if liveAdInfo.HasPoster != "0" {
            links := parameters["links"].(string)
            linksList := strings.Fields(links)
            if len(linksList) == 4 {
                msgSkipUrlValue := strings.Replace(strings.Replace(linksList[1], "u'", "", 1), "',", "", 1)
                msgImageUrlValue := strings.Replace(strings.Replace(linksList[3], "u'", "", 1), "'}", "", 1)
                linksMap := map[string]string{"msg_skip_url": msgSkipUrlValue, "msg_image_url": msgImageUrlValue}
                liveAdInfo.Links = linksMap
            }
        }
    }
    
}