package middlewares

import (
	"ladipage_server/common/utils"
	"ladipage_server/core/constant"
	"ladipage_server/core/services"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareJwt struct {
	jwtService *services.JwtService
	user       *services.UserService
}

func NewMiddlewareJwt(jwtService *services.JwtService, user *services.UserService) *MiddlewareJwt {
	return &MiddlewareJwt{
		jwtService: jwtService,
		user:       user,
	}
}

func (m *MiddlewareJwt) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		// Format: "Bearer <token>"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Expected Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]

		claims, err := m.jwtService.VerifyToken(c, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		user, err := m.user.Profile(c, claims.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		if utils.FormatTime(user.UpdatedAt) != utils.FormatTime(claims.UpdatedAccountUser) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  1,
				"error": "User information has been updated. Please log in again."})
			c.Abort()
			return
		}

		c.Set("userId", claims.Id)
		c.Set("userName", claims.UserName)
		c.Set("updatedAt", claims.UpdatedAccountUser)
		c.Set("role", user.Role)

		c.Next()
	}
}

func (m *MiddlewareJwt) AuthorizationIsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		// Format: "Bearer <token>"
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Expected Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]

		claims, err := m.jwtService.VerifyToken(c, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		user, err := m.user.Profile(c, claims.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		if utils.FormatTime(user.UpdatedAt) != utils.FormatTime(claims.UpdatedAccountUser) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  1,
				"error": "User information has been updated. Please log in again."})
			c.Abort()
			return
		}
		if user.Role != constant.RoleIsAdmin {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("userId", claims.Id)
		c.Set("userName", claims.UserName)
		c.Set("updatedAt", claims.UpdatedAccountUser)
		c.Set("role", user.Role)
		c.Next()
	}
}
