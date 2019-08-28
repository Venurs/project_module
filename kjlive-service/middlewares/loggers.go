package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"

	"kjlive-service/models"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		ip := c.Request.RemoteAddr
		requestMethod := c.Request.Method
		requestUrl := c.Request.URL.Path
		statusCode := c.Writer.Status()
		referrer := c.Request.Referer()
		ua := c.Request.UserAgent()
		
		sgBid, err := c.Request.Cookie("SGBID")
		sgBidValue := ""
		if err == nil {
			sgBidValue = sgBid.Value
		}
		sgUuid, err := c.Request.Cookie("SGUUID")
		sgUuidValue := ""
		if err == nil {
			sgUuidValue = sgUuid.Value
		}
		sgSid, err := c.Request.Cookie("SGSID")
		sgSidValue := ""
		if err == nil {
			sgSidValue = sgSid.Value
		}
		user_obj, _ := c.Get("user")
		casUserId := user_obj.(models.UserInfo).CasUserId
		glg.Warnf("IP:%v | cas_user_Id :%v | Method:%v | Url:%v | Code:%v | Referrer:%v | UserAgent:%v | SGBID:%v | SGUUID:%v | SGSID:%v",
			ip, casUserId, requestMethod, requestUrl, statusCode, referrer, ua, sgBidValue, sgUuidValue, sgSidValue)
	}
}
