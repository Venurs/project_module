package controllers

import (
    "github.com/gin-gonic/gin"
    "kjlive-service/models"
    "strconv"
    "kjlive-service/utils"
    "math/rand"
    "encoding/json"
    "kjlive-service/redis"
    "time"
    "kjlive-service/mq"
    "strings"
    "kjlive-service/exception"
    "reflect"
    "kjlive-service/logs"
    "github.com/kisielk/og-rek"
)

type ResultData struct {
    LiveInfo     LiveInfoSerializer  `json:"live_info"`
    UserInfo     models.LiveUserInfo `json:"user_info"`
    Ch           string              `json:"ch"`
    IsChargePage int                 `json:"is_charge_page"`
    LivePoster   models.LiveAdInfo   `json:"live_poster"`
}
type ResultChargeData struct {
    LiveChargeInfo models.LiveChargeInfoSerializer `json:"live_info"`
    IsShowPopup    bool                            `json:"is_show_popup"`
    IsChargePage   int                             `json:"is_charge_page"`
    Ch             string                          `json:"ch"`
}

type LiveInfoSerializer struct {
    Id                  int                `json:"id"`
    RoomId              string             `json:"room_id"`
    Password            string             `json:"password"`
    LecturerId          int                `json:"lecturer_id"`
    LecturerName        string             `json:"lecturer_name"`
    LecturerHonour      string             `json:"lecturer_honour"`
    LecturerPictureUrl  string             `json:"lecturer_picture_url"`
    LecturerSummary     string             `json:"lecturer_summary"`
    Name                string             `json:"name"`
    Summary             string             `json:"summary"`
    IsLive              int                `json:"is_live"`
    OnlineNumber        int                `json:"online_number"`
    StartTime           string             `json:"start_time"`
    LikedCount          int                `json:"liked_count"`
    WechatShareImageUrl string             `json:"wechat_share_image_url"`
    AlreadyStart        int8               `json:"already_start"`
    IsOver              int                `json:"is_over"`
    BaseEventId         int                `json:"event_id"`
    Broadcast           string             `json:"broadcast"`
    WechatShareSummary  string             `json:"wechat_share_summary"`
    RoomName            string             `json:"room_name"`
    IsHalfTime          int                `json:"is_half_time"`
    HalfTimeImageUrl    string             `json:"half_time_image_url"`
    BroadcastUrl        string             `json:"broadcast_url"`
    ImageUrl            string             `json:"image_url"`
    WatchCountIncrement int                `json:"watch_count_increment"`
    AdSmallImageUrl     string             `json:"ad_small_image_url"`
    AdBigImageUrl       string             `json:"ad_big_image_url"`
    AdIsShow            int8               `json:"ad_is_show"`
    AdSkipUrl           string             `json:"ad_skip_url"`
    IsShowQidianQq      int                `json:"is_show_qidian_qq"`
    FreeIsShow          int8               `json:"free_is_show"`
    FreeImageUrl        string             `json:"free_image_url"`
    IsCharge            int                `json:"is_charge"`
    PresentPrice        float64            `json:"present_price"`
    ShowPrimeExchange   int                `json:"show_prime_exchange"`
    IsMake              int                `json:"is_make"`
    IsLike              int                `json:"is_like"`
    IsPay               int                `json:"is_pay"`
    VideoList           []models.VideoInfo `json:"video_list"`
    Token               string             `json:"token"`
    UserId              string             `json:"userId"`
    PopularizeSkipUrl   string             `json:"popularize_skip_url"`
    PopularizeImageUrl  string             `json:"popularize_image_url"`
    PopularizeIsShow    int8               `json:"popularize_is_show"`
    TemplateId          int8               `json:"template_id"`
    TemplateIsRequired  bool               `json:"template_is_required"`
}

func LiveInfoController() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 取得ID
        liveId := c.Query("live_id")
        pm := c.Query("pm")
        pop_param := c.Query("pop_param")
        id, err := strconv.Atoi(liveId)
        if err != nil {
            err := exception.ValidateError{
                Err:     err,
                Code:    "validate",
                Message: map[string]string{"live_id": "required"}}
            c.JSON(200, err)
            return
        }
        pmId, _ := strconv.Atoi(pm)
        // 取得user
        user_obj, _ := c.Get("user")
        user := user_obj.(models.UserInfo)
        // 取得收费信息
        liveChargeSerializer, _ := models.GetLiveChargeInfoById(id)
        
        // 不收费
        if liveChargeSerializer.Id == 0 {
            resultData, err := showLivePage(id, user, pmId)
            if err != nil {
                err := exception.ValidateError{
                    Err:     err,
                    Code:    "error:data_does_not_exist",
                    Message: map[string]string{"error:data_does_not_exist": "详细数据不存在"}}
                c.JSON(200, err)
                return
            }
            byteJson, _ := json.Marshal(resultData)
            c.Data(200, "text/plain", byteJson)
            return
            // 收费
        }
        // 用户未登录
        if !user.IsAuthenticated {
            // 去付费页面
            resultData, err := showChargePage(id, pop_param, user, pmId)
            if err != nil {
                err := exception.ValidateError{
                    Err:     err,
                    Code:    "error:charge_data_does_not_exist",
                    Message: map[string]string{"error:charge_data_does_not_exist": "付费数据不存在"}}
                c.JSON(200, err)
                return
            }
            c.JSON(200, resultData)
            return
            // 用户登录
        }
        // 有购买记录
        if isPay(id, liveChargeSerializer.Id, user.Id) {
            // 去直播页面
            // 取得直播页信息
            resultData, err := showLivePage(id, user, pmId)
            if err != nil {
                err := exception.ValidateError{
                    Err:     err,
                    Code:    "error:data_does_not_exist",
                    Message: map[string]string{"error:data_does_not_exist": "详细数据不存在"}}
                c.JSON(200, err)
                return
            }
            c.JSON(200, resultData)
        } else {
            // 去付费页面
            resultData, err := showChargePage(id, pop_param, user, pmId)
            if err != nil {
                err := exception.ValidateError{
                    Err:     err,
                    Code:    "error:charge_data_does_not_exist",
                    Message: map[string]string{"error:charge_data_does_not_exist": "付费数据不存在"}}
                c.JSON(200, err)
                return
            }
            c.JSON(200, resultData)
        }
    }
}

// 返回显示直播页面的信息
func showLivePage(id int, user models.UserInfo, pm int) (resultData ResultData, err error) {
    serializerData := redis.GetMap("LIVE_INFO_DETAIL_" + strconv.Itoa(id))
    if serializerData == nil {
        // 取得直播页信息
        liveInfoSerializer, err := models.GetLiveLiveinfoById(id)
        if err != nil {
            return resultData, err
        }
        m := getMapLiveInfo(liveInfoSerializer)
        redis.Set("LIVE_INFO_DETAIL_"+strconv.Itoa(id), m, 24*time.Hour)
        serializerData = m
    }
    serializerData = setVideoList(serializerData)
    var serializer LiveInfoSerializer
    jsonSerializerData, err := json.Marshal(serializerData)
    json.Unmarshal(jsonSerializerData, &serializer)
    // 用户登录
    if user.IsAuthenticated {
        serializer.setVideoBuyStatus(user.Id)
        serializer.OnlineNumber = models.GetLiveAccessNumber(id, false) + serializer.WatchCountIncrement
        resultData.UserInfo.UserId = user.GetLiveId()
        resultData.UserInfo.UserName = user.GetName()
        resultData.UserInfo.UserIsSurvey = models.GetUserQuestionnaireSurveyStatus(user.Id, serializer.TemplateId)
        resultData.UserInfo.UserAvatarUrl = user.Avatar
        if models.IsMakeLive(id, user.Id) {
            serializer.IsMake = 1
        }
        if models.IsLikedLive(id, user.Id) {
            serializer.IsLike = 1
        }
        _, err := models.InsertLiveAccessLog(user.Id, id, pm)
        if err != nil {
            logs.Error(err)
        }
        args := []interface{}{user.CasUserId,
            "http://www.selleruc.com/online_video/?id=" + strconv.Itoa(id),
            time.Now().Format("2006/01/02 15:04:05"),
            serializer.Name,
            serializer.WechatShareImageUrl}
        mq.SendMessageByManager("dynamic_information",
            "dynamic_information.live_watch",
            "dynamic_information.tasks.live_watch_information",
            args)
    } else {
        serializer.OnlineNumber = models.GetLiveAccessNumber(id, true) + serializer.WatchCountIncrement
        resultData.UserInfo.UserId = strconv.Itoa(rand.Intn(10000000))
        resultData.UserInfo.UserName = "游客"
        resultData.UserInfo.UserAvatarUrl = "https://oajua4pqj.qnssl.com/o_1c58cquc11n9q1tj7i491vnr1lv57.jpg"
        
        liveAccessLogMap := map[string]interface{}{"live": id, "access_time": time.Now().Format("2006-01-02"), "promotion_manage_id": pm}
        
        notLoginUserAccessLog := redis.GetSlice("NOT_LOGIN_USER_ACCESS_LOG_" + strconv.Itoa(id))
        if len(notLoginUserAccessLog) == 0 {
            notLoginUserAccessLog = append(notLoginUserAccessLog, liveAccessLogMap)
        } else {
            notLoginUserAccessLog = append(notLoginUserAccessLog, liveAccessLogMap)
        }
        redis.Set("NOT_LOGIN_USER_ACCESS_LOG_" + strconv.Itoa(id), notLoginUserAccessLog, 24*30*time.Hour)
    }
    if serializer.IsOver == 0 && serializer.IsHalfTime == 0 {
        IsLive, _ := utils.GetLiveStatus(serializer.RoomId)
        serializer.IsLive = IsLive
    }
    resultData.LiveInfo = serializer
    resultData.Ch = models.GetEventInfoByLiveId(id).SignUrl
    resultData.LivePoster = models.GetLivePoster(id)
    return
}

func getMapLiveInfo(serializer models.LiveInfoSerializer) map[string]interface{} {
    t := reflect.TypeOf(serializer)
    v := reflect.ValueOf(serializer)
    
    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[strings.ToLower(t.Field(i).Tag.Get("json"))] = v.Field(i).Interface()
    }
    return data
}

// 返回显示收费页面的数据
func showChargePage(liveId int, pop_param string, user models.UserInfo, pm int) (resultData ResultChargeData, err error) {
    liveChargeInfo, err := models.GetLiveLiveChargeInfoByLiveId(liveId)
    if err != nil {
        return
    }
    if pop_param != "" {
        setLiveChargePop(pop_param, &liveChargeInfo)
    }
    // 判断是否弹窗（优惠信息）
    // 登录用户弹过一次，下次不再弹了
    // 未登录用户一直弹
    if pop_param != "" && liveChargeInfo.PreferPrice > 0 {
        resultData.IsShowPopup = true
        if user.IsAuthenticated {
            user_list := redis.GetSlice("LIVE_CHARGE_POP_IS_SHOW_POPUP_GO_" + pop_param)
            if user_list == nil {
                user_info := []int{user.Id}
                redis.Set("LIVE_CHARGE_POP_IS_SHOW_POPUP_GO_"+pop_param, user_info, time.Hour*12*30)
            } else {
                for _, user_info := range user_list {
                    if int(user_info.(int64)) == user.Id {
                        resultData.IsShowPopup = false
                    }
                }
                if resultData.IsShowPopup {
                    user_list = append(user_list, user.Id)
                    redis.Set("LIVE_CHARGE_POP_IS_SHOW_POPUP_GO_"+pop_param, user_list, time.Hour*12*30)
                }
            }
        }
    } else {
        resultData.IsShowPopup = false
    }
    
    if user.IsAuthenticated {
        _, err := models.InsertLiveAccessLog(user.Id, liveId, pm)
        if err != nil {
            logs.Error(err)
        }
    }
    
    resultData.LiveChargeInfo = liveChargeInfo
    resultData.IsChargePage = 1
    resultData.Ch = models.GetEventInfoByLiveId(liveId).SignUrl
    return
}

// 判断是否购买了直播
func isPay(liveId int, liveChargeInfoId int, userId int) bool {
    // 购买了姜涵套餐能免费观看姜涵直播
    params := models.GetParameter("CMS_SET_LIVE_RELATION")
    if params != nil {
        setLiveReList := strings.Fields(params["set_live_relation"].(string))
        var setIdList []int
        if strings.Replace(strings.Replace(strings.Replace(setLiveReList[1], "u", "", 1), "'", "", 2), ",", "", 1) == strconv.Itoa(liveId) {
            setId, err := strconv.Atoi(strings.Replace(setLiveReList[3], "}]", "", 1))
            if err == nil {
                setIdList = append(setIdList, setId)
            }
        }
        if len(setIdList) > 0 && models.IsBuySet(setIdList, userId) {
            return true
        }
    }
    // 有没有购买记录
    return models.IsPayByLiveChargeInfoId(liveChargeInfoId, userId)
}

// 设定视频的购买状态
func (serializer *LiveInfoSerializer) setVideoBuyStatus(userId int) {
    videoIdList := [3]int{}
    for i, video := range serializer.VideoList {
        videoIdList[i] = video.Id
    }
    if len(serializer.VideoList) > 0 {
        objectList := models.GetContenttypeuserByIdList(videoIdList, userId)
        for _, objectId := range objectList {
            for i, video := range serializer.VideoList {
                if video.Id == objectId {
                    serializer.VideoList[i].BuyStatus = "1"
                }
            }
        }
    }
    
}

// 设定视频list
func setVideoList(serializer map[string]interface{}) map[string]interface{} {
    var videoList []map[string]interface{}
    var video_id int
    switch serializer["video_id"].(type) {
    case ogórek.None:
        video_id = 0
    case int64:
        video_id = int(serializer["video_id"].(int64))
    case float64:
        video_id = int(serializer["video_id"].(float64))
    }
    var video_2_id int
    switch serializer["video_2_id"].(type) {
    case ogórek.None:
        video_2_id = 0
    case int64:
        video_2_id = int(serializer["video_2_id"].(int64))
    case float64:
        video_2_id = int(serializer["video_2_id"].(float64))
    }
    var video_3_id int
    switch serializer["video_3_id"].(type) {
    case ogórek.None:
        video_3_id = 0
    case int64:
        video_3_id = int(serializer["video_3_id"].(int64))
    case float64:
        video_3_id = int(serializer["video_3_id"].(float64))
    }
    if video_id > 0 {
        video := map[string]interface{}{
            "id":            serializer["video_id"],
            "present_price": serializer["video_present_price"],
            "image_url":     serializer["video_image_url"],
            "name":          serializer["video_name"],
            "buy_status":    serializer["video_buy_status"],
        }
        videoList = append(videoList, video)
    }
    if video_2_id > 0 {
        video := map[string]interface{}{
            "id":            serializer["video_2_id"],
            "present_price": serializer["video_2_present_price"],
            "image_url":     serializer["video_2_image_url"],
            "name":          serializer["video_2_name"],
            "buy_status":    serializer["video_2_buy_status"],
        }
        videoList = append(videoList, video)
    }
    if video_3_id > 0 {
        video := map[string]interface{}{
            "id":            serializer["video_3_id"],
            "present_price": serializer["video_3_present_price"],
            "image_url":     serializer["video_3_image_url"],
            "name":          serializer["video_3_name"],
            "buy_status":    serializer["video_3_buy_status"],
        }
        videoList = append(videoList, video)
    }
    serializer["video_list"] = videoList
    
    return serializer
}

// 设定优惠价格
func setLiveChargePop(pop_param string, liveChargeInfo *models.LiveChargeInfoSerializer) {
    liveChargePopSerializer, _ := models.GetLiveChargePopById(pop_param)
    if liveChargePopSerializer.Id == 0 {
        liveChargeInfo.IsUsedUp = 0
    } else {
        if liveChargePopSerializer.IsUsed == 0 && liveChargePopSerializer.StartTime.Before(time.Now()) && liveChargePopSerializer.EndTime.After(time.Now()) {
            liveChargeInfo.PresentPrice = liveChargeInfo.PresentPrice - liveChargePopSerializer.DiscountAmount
            if liveChargeInfo.PresentPrice < 0 {
                liveChargeInfo.PresentPrice = 0
            }
            liveChargeInfo.PreferPrice = liveChargePopSerializer.DiscountAmount
        } else if liveChargePopSerializer.IsUsed == 1 {
            liveChargeInfo.IsUsedUp = 2
        } else if liveChargePopSerializer.StartTime.After(time.Now()) || liveChargePopSerializer.EndTime.Before(time.Now()) {
            liveChargeInfo.IsUsedUp = 3
        } else {
            liveChargeInfo.IsUsedUp = 1
        }
    }
}
