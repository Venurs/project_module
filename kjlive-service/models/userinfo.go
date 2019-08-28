package models

import (
    "strings"
    "strconv"
    "bytes"
    "io/ioutil"
    "time"
    "net/http"
    "encoding/json"
    
    "kjlive-service/conf"
    "kjlive-service/redis"
    "kjlive-service/logs"
)

type UserInfo struct {
    NotReadCount       int      `json:"not_read_count"`
    IsAuthenticated    bool     `json:"is_authenticated"`
    IsPrimeUser        bool     `json:"is_prime_user"`
    Id                 int      `json:"id"`
    IsSuperuser        bool     `json:"is_superuser"`
    IsOver             int      `json:"is_over"`
    UsedToBePrimeUser  string   `json:"used_to_be_prime_user"`
    Pk                 int      `json:"pk"`
    Email              string   `json:"email"`
    Username           string   `json:"username"`
    NicknameExport     string   `json:"nickname_export"`
    IsActive           bool     `json:"is_active"`
    Name               string   `json:"name"`
    IsAnonymous        bool     `json:"is_anonymous"`
    UserName           string   `json:"user_name"`
    CasUserId          int      `json:"cas_user_id"`
    IsStaff            bool     `json:"is_staff"`
    Avatar             string   `json:"avatar"`
    AllPermissions     []string `json:"all_permissions"`
    ProfileUserUuidHex string   `json:"profile__user_uuid_hex"`
    PrimeWeight        int      `json:"prime_weight"`
}

func (user *UserInfo) SetUserName()  {
    user.UserName = user.NicknameExport
}

func (user *UserInfo) GetName() string {
    if user.UserName != "Hack" {
        return user.UserName
    } else if user.Name != "" {
        return user.Name
    } else {
        return user.Username
    }
}


func (user *UserInfo) GetLiveId() string {
    if user.Id == 0 {
        return "1000000"
    } else if user.Id > 9999999 {
        return strconv.Itoa(user.Id)[:7]
    } else if user.Id < 1000000{
        userId := strconv.Itoa(user.Id)
        return strings.Repeat("1", 7 -len(userId)) + userId
    } else {
        return strconv.Itoa(user.Id)
    }
}

type JosnRpc struct {
    Jsonrpc string   `json:"jsonrpc"`
    Id      int      `json:"id"`
    Result  UserInfo `json:"result"`
}

type LiveUserInfo struct {
    UserId string `json:"user_id"`
    UserName string `json:"user_name"`
    UserAvatarUrl string `json:"user_avatar_url"`
    UserIsSurvey int `json:"user_is_survey"`
}


func (jsonRpc *JosnRpc) GetUserInfo(sessionId string) {
    body := redis.Get("USER_INFO_" + sessionId)
    if body == "" {
        jsonRpc.getUserInfo(sessionId)
    } else {
        if err := json.Unmarshal([]byte(body), &jsonRpc.Result); err != nil {
            logs.Error(err)
        }
    }
}

func (jsonRpc *JosnRpc) GetNewUserInfo(sessionId string) {
    jsonRpc.getUserInfo(sessionId)
}

func (jsonRpc *JosnRpc) getUserInfo(sessionId string) {
    client := &http.Client{}
    inputMap := make(map[string]interface{})
    inputMap["method"] = "account.user_by_session_key"
    inputMap["params"] = []string{sessionId}
    inputMap["jsonrpc"] = "2.0"
    inputMap["id"] = 0
    bytesData, err := json.Marshal(inputMap)
    reader := bytes.NewReader(bytesData)
    response, err := client.Post(conf.Settings.UnicornServiceRpcUrl, "application/json", reader)
    if err != nil {
        logs.Error(err)
    } else {
        defer response.Body.Close()
        body, _ := ioutil.ReadAll(response.Body)
        if err := json.Unmarshal([]byte(string(body)), &jsonRpc); err != nil {
            logs.Error(err)
        } else {
            jsonRpc.Result.SetUserName()
            jsonUserInfo, _ := json.Marshal(jsonRpc.Result)
            redis.Set("USER_INFO_" + sessionId, string(jsonUserInfo), 60*60*24*time.Second)
        }
    }
}