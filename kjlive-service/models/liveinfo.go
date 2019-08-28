package models

import (
	"time"
	"github.com/astaxie/beego/orm"

	"kjlive-service/utils"
	"kjlive-service/conf"
	"kjlive-service/logs"
)

type LiveInfoSerializer struct {
	Id                         int     `json:"id" orm:"column(id)"`
	RoomId                     string  `json:"room_id" orm:"column(room_id)"`
	Password                   string  `json:"password" orm:"column(password)"`
	LecturerId                 int     `json:"lecturer_id" orm:"column(lecturer_id)"`
	LecturerLecturerName       string  `json:"lecturer_lecturer_name" orm:"column(lecturer_lecturer_name)"`
	LecturerLecturerSummary    string  `json:"lecturer_lecturer_summary" orm:"column(lecturer_lecturer_summary)"`
	LecturerLecturerPictureUrl string  `json:"lecturer_lecturer_picture_url" orm:"column(lecturer_lecturer_picture_url)"`
	LecturerLecturerHonour     string  `json:"lecturer_lecturer_honour" orm:"column(lecturer_lecturer_honour)"`
	TemplateId                 int     `json:"template_id" orm:"column(template_id)"`
	TemplateIsRequired         bool    `json:"template_is_required" orm:"column(template_is_required)"`
	Name                       string  `json:"name" orm:"column(name)"`
	Summary                    string  `json:"summary" orm:"column(summary)"`
	StartTime                  string  `json:"start_time" orm:"column(start_time)"`
	LikedCount                 int     `json:"liked_count" orm:"column(liked_count)"`
	VideoId                    int     `json:"video_id" orm:"column(video_id)"`
	Video2Id                   int     `json:"video_2_id" orm:"column(video_2_id)"`
	Video3Id                   int     `json:"video_3_id" orm:"column(video_3_id)"`
	VideoVideoName             string  `json:"video_name" orm:"column(video_video_name)"`
	Video2VideoName            string  `json:"video_2_name" orm:"column(video_2_video_name)"`
	Video3VideoName            string  `json:"video_3_name" orm:"column(video_3_video_name)"`
	VideoImageUrl              string  `json:"video_image_url" orm:"column(video_image_url)"`
	Video2ImageUrl             string  `json:"video_2_image_url" orm:"column(video_2_image_url)"`
	Video3ImageUrl             string  `json:"video_3_image_url" orm:"column(video_3_image_url)"`
	VideoPresentPrice          float64 `json:"video_present_price" orm:"column(video_present_price)"`
	Video2PresentPrice         float64 `json:"video_2_present_price" orm:"column(video_2_present_price)"`
	Video3PresentPrice         float64 `json:"video_3_present_price" orm:"column(video_3_present_price)"`
	WechatShareImageUrl        string  `json:"wechat_share_image_url" orm:"column(wechat_share_image_url)"`
	BaseEventId                int     `json:"event_id" orm:"column(event_id)"`
	IsOver                     int     `json:"is_over" orm:"column(is_over)"`
	Broadcast                  string  `json:"broadcast" orm:"column(broadcast)"`
	WechatShareSummary         string  `json:"wechat_share_summary" orm:"column(wechat_share_summary)"`
	HalfTimeImageUrl           string  `json:"half_time_image_url" orm:"column(half_time_image_url)"`
	IsHalfTime                 int     `json:"is_half_time" orm:"column(is_half_time)"`
	RoomName                   string  `json:"room_name" orm:"column(room_name)"`
	BroadcastUrl               string  `json:"broadcast_url" orm:"column(broadcast_url)"`
	ImageUrl                   string  `json:"image_url" orm:"column(image_url)"`
	WatchCountIncrement        int     `json:"watch_count_increment" orm:"column(watch_count_increment)"`
	LecturerName               string  `json:"lecturer_name" orm:"column(lecturer_name)"`
	LecturerHonour             string  `json:"lecturer_honour" orm:"column(lecturer_honour)"`
	LecturerPictureUrl         string  `json:"lecturer_picture_url" orm:"column(lecturer_picture_url)"`
	LecturerSummary            string  `json:"lecturer_summary" orm:"column(lecturer_summary)"`
	IsShowQidianQq             int     `json:"is_show_qidian_qq" orm:"column(is_show_qidian_qq)"`
	AdSmallImageUrl            string  `json:"ad_small_image_url" orm:"column(ad_small_image_url)"`
	AdBigImageUrl              string  `json:"ad_big_image_url" orm:"column(ad_big_image_url)"`
	AdIsShow                   int8    `json:"ad_is_show" orm:"column(ad_is_show)"`
	AdSkipUrl                  string  `json:"ad_skip_url" orm:"column(ad_skip_url)"`
	FreeIsShow                 int8    `json:"free_is_show" orm:"column(free_is_show)"`
	FreeImageUrl               string  `json:"free_image_url" orm:"column(free_image_url)"`
	IsLive                     int     `json:"is_live"`
	OnlineNumber               int     `json:"online_number"`
	AlreadyStart               int8    `json:"already_start"`
	VideoBuyStatus             string  `json:"video_buy_status"`
	Video2BuyStatus            string  `json:"video_2_buy_status"`
	Video3BuyStatus            string  `json:"video_3_buy_status"`
	IsCharge                   int     `json:"is_charge"`
	PresentPrice               float64 `json:"present_price"`
	ShowPrimeExchange          float64 `json:"show_prime_exchange"`
	Token                      string  `json:"token"`
	UserId                     string  `json:"userId"`
	PopularizeSkipUrl          string  `json:"popularize_skip_url" orm:"column(popularize_skip_url)"`
	PopularizeImageUrl         string  `json:"popularize_image_url" orm:"column(popularize_image_url)"`
	PopularizeIsShow           int8    `json:"popularize_is_show" orm:"column(popularize_is_show)"`
}

// 取得直播情报
func GetLiveLiveinfoById(id int) (liveInfoSerializer LiveInfoSerializer, err error) {
	o := orm.NewOrm()
	sql := `
        SELECT
          T0.id                     id,
          T0.room_id                room_id,
          T0.password               password,
		  T0.template_id            template_id,
          T0.template_is_required   template_is_required,
          T1.id                     lecturer_id,
          T1.lecturer_name          lecturer_lecturer_name,
          T1.lecturer_summary       lecturer_lecturer_summary,
          T1.lecturer_picture_url   lecturer_lecturer_pictureUrl,
          T1.lecturer_honour        lecturer_lecturer_honour,
          T0.name                   name,
          T0.summary                summary,
          DATE_FORMAT(T0.start_time, '%Y-%m-%d %H:%i')             start_time,
          T0.liked_count            liked_count,
          T2.id                     video_id,
          T3.id                     video_2_id,
          T4.id                     video_3_id,
          T2.video_name             video_video_name,
          T3.video_name             video_2_video_name,
          T4.video_name             video_3_video_name,
          T2.image_url              video_image_url,
          T3.image_url              video_2_image_url,
          T4.image_url              video_3_image_url,
          T5.present_price          video_present_price,
          T6.present_price          video_2_present_price,
          T7.present_price          video_3_present_price,
          T0.wechat_share_image_url wechat_share_image_url,
          T0.baseevent_id           event_id,
          T0.is_over                is_over,
          T0.broadcast              broadcast,
          T0.wechat_share_summary   wechat_share_summary,
          T0.half_time_image_url    half_time_image_url,
          T0.is_half_time           is_half_time,
          T0.room_name              room_name,
          T0.broadcast_url          broadcast_url,
          T0.image_url              image_url,
          T0.watch_count_increment  watch_count_increment,
          T0.lecturer_name          lecturer_name,
          T0.lecturer_honour        lecturer_honour,
          T0.lecturer_picture_url   lecturer_picture_url,
          T0.lecturer_summary       lecturer_summary,
          T0.is_show_qidian_qq      is_show_qidian_qq,
          T8.small_image_url        ad_small_image_url,
          T8.big_image_url          ad_big_image_url,
          T8.is_show                ad_is_show,
          T8.skip_url               ad_skip_url,
          T8.free_is_show           free_is_show,
          T8.free_image_url         free_image_url,
          T8.popularize_is_show     popularize_is_show,
          T8.popularize_image_url   popularize_image_url,
          T8.popularize_skip_url    popularize_skip_url
        FROM live_liveinfo T0
          INNER JOIN video_lecturer T1 ON T1.id = T0.lecturer_id
          LEFT JOIN video_video T2 ON T2.id = T0.video_id
          LEFT JOIN video_video T3 ON T3.id = T0.video_2_id
          LEFT JOIN video_video T4 ON T4.id = T0.video_3_id
          LEFT JOIN video_coursevideoinfo T5 ON T5.video_id = T2.id
          LEFT JOIN video_coursevideoinfo T6 ON T6.video_id = T3.id
          LEFT JOIN video_coursevideoinfo T7 ON T7.video_id = T4.id
          INNER JOIN live_liveadinfo T8 ON T0.id = T8.live_info_id
        WHERE T0.id = ?
          AND T0.delete_flag=0
    `

	err = o.Raw(sql, id).QueryRow(&liveInfoSerializer)
	if err == nil {
		// 直播状态设定
		liveInfoSerializer.setLiveStatus()
		// 直播是否已经开始设定
		liveInfoSerializer.setAlreadyStart()
		// 购买状态设定
		liveInfoSerializer.setVideoBuyStatus()
		// 讲师信息设定
		liveInfoSerializer.setLecturerInfo()
		return liveInfoSerializer, nil
	}
	return
}

// 直播Token
func (serializer *LiveInfoSerializer) setLiveStatus() {
	serializer.Token = utils.GetLiveTelecastToken()
	serializer.UserId = conf.Settings.PolyvLiveUserId
}

// 直播是否已经开始设定
func (serializer *LiveInfoSerializer) setAlreadyStart() {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04", serializer.StartTime, loc)
	if err != nil {
		serializer.AlreadyStart = 0
		logs.Error(err)
		return
	}
	if theTime.Before(time.Now()) {
		serializer.AlreadyStart = 1
	} else {
		serializer.AlreadyStart = 0
	}
}

// 购买状态
func (serializer *LiveInfoSerializer) setVideoBuyStatus() {
	if serializer.VideoId > 0 {
		if serializer.VideoPresentPrice > 0 {
			serializer.VideoBuyStatus = "0"
		} else {
			serializer.VideoBuyStatus = "2"
		}
	}
	if serializer.Video2Id > 0 {
		if serializer.Video2PresentPrice > 0 {
			serializer.Video2BuyStatus = "0"
		} else {
			serializer.Video2BuyStatus = "2"
		}
	}
	if serializer.Video3Id > 0 {
		if serializer.Video3PresentPrice > 0 {
			serializer.Video3BuyStatus = "0"
		} else {
			serializer.Video3BuyStatus = "2"
		}
	}
}

// 讲师信息修改
func (serializer *LiveInfoSerializer) setLecturerInfo() {
	if serializer.LecturerName == "" {
		serializer.LecturerName = serializer.LecturerLecturerName
	}
	if serializer.LecturerSummary == "" {
		serializer.LecturerSummary = serializer.LecturerLecturerSummary
	}
	if serializer.LecturerPictureUrl == "" {
		serializer.LecturerPictureUrl = serializer.LecturerLecturerPictureUrl
	}
	if serializer.LecturerHonour == "" {
		serializer.LecturerHonour = serializer.LecturerLecturerHonour
	}
}
