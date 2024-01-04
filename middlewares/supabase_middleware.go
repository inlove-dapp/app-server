package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"inlove-app-server/constants/environments"
	"inlove-app-server/constants/keys"
	"inlove-app-server/types"
	"net/http"
	"strings"
)

// JWTAuthMiddleware is responsible for validating the JWT token sent in the header of the request from
// the client authenticated by `Supabase`.
//
// In the context of the `Supabase` authenticated app, the JWT token is sent in the cookie.
// However, we cannot access the cookie directly from our backend since it is a cross-origin request.
// Therefore, we need to send the JWT token in the header of the request.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("Authorization")
		jwtToken = strings.Replace(jwtToken, "Bearer ", "", 1)

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid signing algorithm"})
			}
			return []byte(environments.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set(keys.UserKey, types.User{Id: token.Claims.(jwt.MapClaims)["sub"].(string)})
		c.Next()
	}
}
