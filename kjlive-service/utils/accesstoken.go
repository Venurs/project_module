package utils

import (
    //"io/ioutil"
    //"strings"
    //"time"
    "kjlive-service/redis"
    //"net/http"
    //"kjlive-service/conf"
    //"bytes"
    //"encoding/json"
    //"fmt"
)

// 取得accessToken
func GetAccessToken() string {
    accessToken := redis.Get("SELLERUC_CLIENT_ACCESS_TOKEN")
    //if accessToken == "" {
    //    client := &http.Client{}
    //    inputMap := make(map[string]interface{})
    //    inputMap["grant_type"] = "client_credentials"
    //    bytesData, err := json.Marshal(inputMap)
    //    reader := bytes.NewReader(bytesData)
    //    req, _ := http.NewRequest("POST", conf.Settings.APP_MANAGE_URL, reader)
    //    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    //    response, err := client.Do(req)
    //    if err != nil {
    //        panic(err)
    //    }
    //    defer response.Body.Close()
    //    body, _ := ioutil.ReadAll(response.Body)
    //    accessToken = string(body[:])
    //    fmt.Println(accessToken)
    //    redis.Set("SELLERUC_CLIENT_ACCESS_TOKEN", accessToken, 60 * 2*time.Second)
    //}
    return accessToken
}