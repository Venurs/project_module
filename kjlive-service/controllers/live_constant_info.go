package controllers

import (
    "github.com/gin-gonic/gin"
    "time"
    "kjlive-service/redis"
    "kjlive-service/models"
    "kjlive-service/utils"
    "strconv"
    "encoding/json"
    "fmt"
    "kjlive-service/exception"
)

func LiveConstantInfoController() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 取得ID
        liveId := c.Query("live_id")
        id, err := strconv.Atoi(liveId)
        if err != nil {
            err := exception.ValidateError{
                Err: err,
                Code: "validate",
                Message:map[string]string{"live_id": "required"}}
            c.JSON(200, err)
            return
        }
        liveConstantInfo, err := getLiveInfo(id)
        if err != nil {
            err := exception.ValidateError{
                Err: err,
                Code: "error:data_does_not_exist",
                Message:map[string]string{"error:data_does_not_exist": "详细数据不存在"}}
            c.JSON(200, err)
        } else {
            liveConstantInfo.setOnlineNumber(id)
            liveConstantInfo.setMakeNumber(id)
            liveConstantInfo.setLiveStatus(id)
            c.JSON(200, liveConstantInfo)
        }
    }
}

type LiveConstantInfo struct {
    RoomId              string `json:"room_id"`
    IsLive              int    `json:"is_live"`
    IsOver              int    `json:"is_over"`
    Broadcast           string `json:"broadcast"`
    IsHalfTime          int    `json:"is_half_time"`
    HalfTimeImageUrl    string `json:"half_time_image_url"`
    WatchCountIncrement int    `json:"watch_count_increment"`
    AdSmallImageUrl     string `json:"ad_small_image_url"`
    AdBigImageUrl       string `json:"ad_big_image_url"`
    AdIsShow            int8   `json:"ad_is_show"`
    AdSkipUrl           string `json:"ad_skip_url"`
    MakeNumber          int    `json:"make_number"`
    UserNumber          int    `json:"user_number"`
}

// 在线人数设定
func (serializer *LiveConstantInfo) setOnlineNumber(liveId int) {
    serializer.UserNumber = models.GetLiveAccessNumber(liveId, false) + serializer.WatchCountIncrement
}

func (serializer *LiveConstantInfo) setMakeNumber(liveId int) {
    makeNumber := redis.GetInt(fmt.Sprintf("GO_LIVE_CONSTASNT_INFO_MAKE_NUMBER_%d", liveId))
    if makeNumber != 0 {
        serializer.MakeNumber = int(makeNumber)
    } else {
        serializer.MakeNumber = models.GetMakeNumberByLiveId(liveId)
        redis.Set(fmt.Sprintf("GO_LIVE_CONSTASNT_INFO_MAKE_NUMBER_%d", liveId), serializer.MakeNumber, 60*60*24*time.Second)
    }
}

func (serializer *LiveConstantInfo) setLiveStatus(liveId int)  {
    if serializer.IsOver == 0 && serializer.IsHalfTime == 0 {
        IsLive, _ := utils.GetLiveStatus(serializer.RoomId)
        serializer.IsLive = IsLive
    }
}

func getLiveInfo(id int) (liveConstantInfo LiveConstantInfo, err error) {
    serializerData := redis.GetMap("LIVE_INFO_DETAIL_" + strconv.Itoa(id))
    if serializerData == nil {
        // 取得直播页信息
        liveInfoSerializer, err := models.GetLiveLiveinfoById(id)
        if err != nil {
            return liveConstantInfo, err
        }
        
        jsonSerializer, _ := json.Marshal(liveInfoSerializer)
        m := make(map[string]interface{})
        err = json.Unmarshal([]byte(string(jsonSerializer)), &m)
        if err != nil {
            return liveConstantInfo, err
        }
        redis.Set("LIVE_INFO_DETAIL_"+strconv.Itoa(id), m, 24*time.Hour)
        serializerData = m
    }
    serializerData["user_number"] = serializerData["online_number"]
    jsonSerializerData, err := json.Marshal(serializerData)
    json.Unmarshal(jsonSerializerData, &liveConstantInfo)
    
    return
}
