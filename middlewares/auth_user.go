package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oj/help"
)

func AuthUserCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.GetHeader("Authorization")
		userClaim, err := help.AnalyseToken(auth)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error() + "unauthorized",
			})
			return
		}
		if userClaim.IsAdmin != 1 {
			context.Abort()
			context.JSON(http.StatusOK, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "forbidden",
			})
			return
		}
		context.Next()

	}
}
