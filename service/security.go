package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"liftapp/config"
)

// IsTokenAllowed returns true when the token is not in the blacklist
//
// Dependency: JWT, Redis database + enable 'INVALIDATE_JWT' in .env
func IsTokenAllowed(jti string) bool {
	// verify that JWT service is enabled in .env
	if !config.IsJWT() {
		return true
	}

	// token blacklist management not enabled, abort
	if !config.InvalidateJWT() {
		return true
	}

	// key not found in blacklist
	return true
}

// JWTBlacklistChecker validates a token against the blacklist
func JWTBlacklistChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jti string
		jtiAccess := strings.TrimSpace(c.GetString("jtiAccess"))
		jtiRefresh := strings.TrimSpace(c.GetString("jtiRefresh"))

		if jtiAccess != "" {
			jti = jtiAccess
			goto CheckBlackList
		}
		if jtiRefresh != "" {
			jti = jtiRefresh
			goto CheckBlackList
		}

	CheckBlackList:
		if !IsTokenAllowed(jti) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
			return
		}

		c.Next()
	}
}
