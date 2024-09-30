package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"murmur-server/model/apperrors"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("userId")

		if id == nil {
			e := apperrors.NewAuthorization(apperrors.InvalidSession)
			c.JSON(e.Status(), gin.H{
				"error": e,
			})
			c.Abort()
			return
		}

		userId := id.(string)

		c.Set("userId", userId)

		// Recreate session to extend its lifetime
		session.Set("userId", id)
		if err := session.Save(); err != nil {
			log.Printf("Failed recreate the session: %v\n", err.Error())
		}

		c.Next()
	}
}
