package controllers

import (
    "github.com/gin-gonic/gin"
    "kjlive-service/models"
    "kjlive-service/utils"
    "kjlive-service/logs"
)

func UserInfoController() gin.HandlerFunc {
    return func(c *gin.Context) {
        user_obj, _ := c.Get("user")
        user := user_obj.(models.UserInfo)
        accessToken := utils.GetAccessToken()
        if accessToken != "" {
            c.SetCookie("access_token", accessToken, 60*60*24, "/", "", false, false)
        }
        if user.IsAuthenticated {
            c.JSON(200, user)
        } else {
            c.Writer.WriteHeader(403)
        }
    }
}
func NewUserInfoController() gin.HandlerFunc {
    return func(c *gin.Context) {
        cookie, err := c.Request.Cookie("liveactivitesessionid")
        var jsonRpc models.JosnRpc
        if err != nil {
            logs.Error(err)
        } else if sessionId := cookie.Value; sessionId != "" {
            jsonRpc.GetNewUserInfo(sessionId)
            if jsonRpc.Result.IsAuthenticated {
                c.JSON(200, jsonRpc.Result)
                return
            }
        }
        c.Writer.WriteHeader(403)
    }
}