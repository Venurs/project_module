package middlewares

import (
    "github.com/gin-gonic/gin"

    "kjlive-service/models"
)


func UserInfoMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        cookie, err := c.Request.Cookie("liveactivitesessionid")
        var jsonRpc models.JosnRpc
        if err == nil {
            if sessionId := cookie.Value; sessionId != "" {
                jsonRpc.GetUserInfo(sessionId)
            }
        }
        c.Set("user", jsonRpc.Result)
    }
}
