package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-CSRF-Token", nosurf.Token(c.Request))
	}
}
