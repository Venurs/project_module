package utils

import (
	"io/ioutil"
	"strings"
	"time"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"encoding/json"
	
	"kjlive-service/conf"
	"kjlive-service/redis"
	"kjlive-service/logs"
)

// 取得直播状态
func GetLiveStatus(room_id string) (isLive int, err error) {
	isLiveString := redis.Get("POLYV_LIVE_ROOM_ID_IS_LIVE_" + room_id)
	if isLiveString == "" {
		stream, _ := getPolyvLiveStream(room_id)

		if stream != "" {

			client := &http.Client{}
			response, err := client.Get("http://api.live.polyv.net/live_status/query?stream=" + stream)
			if err != nil {
				logs.Error(err)
			}
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			if strings.Contains(string(body[:]), "end") {
				isLiveString = "0"
			} else {
				isLiveString = "1"
			}
		} else {
			isLiveString = "0"
		}
		redis.Set("POLYV_LIVE_ROOM_ID_IS_LIVE_" + room_id, isLiveString, 60 * 2*time.Second)
	}
	isLive, err = strconv.Atoi(isLiveString)
	return
}

// 取得直播stream
func getPolyvLiveStream(roomId string) (stream string, err error) {
	stream = redis.Get("POLYV_LIVE_ROOM_ID_STREAM_" + roomId)
	if stream == "" {
		ts := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[:10]
		signByte := bytes.Buffer{}
		signByte.WriteString(conf.Settings.PolyvLiveAppSecret)
		signByte.WriteString("appId")
		signByte.WriteString(conf.Settings.PolyvLiveAppId)
		signByte.WriteString("timestamp")
		signByte.WriteString(ts)
		signByte.WriteString("userId")
		signByte.WriteString(conf.Settings.PolyvLiveUserId)
		signByte.WriteString(conf.Settings.PolyvLiveAppSecret)
		h := md5.New()
		h.Write([]byte(signByte.String()))
		cipherStr := h.Sum(nil)
		sign := strings.ToUpper(hex.EncodeToString(cipherStr))
		requestUrl := bytes.Buffer{}
		requestUrl.WriteString("http://api.live.polyv.net/v1/channels/")
		requestUrl.WriteString(roomId)
		requestUrl.WriteString("/?appId=")
		requestUrl.WriteString(conf.Settings.PolyvLiveAppId)
		requestUrl.WriteString("&timestamp=")
		requestUrl.WriteString(ts)
		requestUrl.WriteString("&userId=")
		requestUrl.WriteString(conf.Settings.PolyvLiveUserId)
		requestUrl.WriteString("&sign=")
		requestUrl.WriteString(sign)

		client := &http.Client{}
		response, err := client.Get(requestUrl.String())
		if err != nil {
			logs.Error(err)
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		var prb = polyvReturnBody{}
		if err := json.Unmarshal([]byte(string(body)), &prb); err == nil {
			if prb.Status == "success" {
				stream = prb.Result.Stream
			}
		}
		redis.Set("POLYV_LIVE_ROOM_ID_STREAM_" +roomId, stream, 60*60*24*365*time.Second)
	}
	return
}

type polyvReturnBody struct {
	Status string `json:"status"`
	Result struct{
		Stream string `json:"stream"`
	} `json:"result"`
}


// 取得直播间聊天室token
func GetLiveTelecastToken() (token string){
	token = redis.Get("LIVE_POLVY_TELECAST_TOKEN")
	if token == "" {
		ts := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[:10]
		h := md5.New()
		h.Write([]byte(ts + "polyvsign"))
		cipherStr := h.Sum(nil)
		sign := hex.EncodeToString(cipherStr)
		client := &http.Client{}
		response, err := client.Get("http://api.live.polyv.net/watchtoken/gettoken?ts=" + ts + "&sign=" + sign)
		if err != nil {
			logs.Error(err)
		}
		if response.StatusCode == 200 {
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			stringBody := string(body)
			token = strings.Trim(strings.Replace(stringBody, "\n", " ", 1), " ")
			
		}
		redis.Set("LIVE_POLVY_TELECAST_TOKEN", token, 60 * 10 * time.Second)
	}

	return token
}

