package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/natanista/go-api/services"
)


func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		const Bearer_schemas = "Bearer "
		header := c.GetHeader("Authorization")
		if header == ""{
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schemas):]

		if!services.NewJWTService().ValidateToken(token){
			c.AbortWithStatus(401)
		}
	}
}