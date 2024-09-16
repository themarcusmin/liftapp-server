package router

import (
	"liftapp/app/controller"
	gconfig "liftapp/config"
	gcontroller "liftapp/controller"
	gmiddleware "liftapp/lib/middleware"
	gservice "liftapp/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(configure *gconfig.Configuration) (*gin.Engine, error) {
	r := gin.Default()

	// API status
	r.GET("", controller.APIStatus)

	// API:v1.0
	v1 := r.Group("/api/v1/")
	{
		// RDBMS
		if gconfig.IsRDBMS() {
			// Register - no JWT required
			v1.POST("register", gcontroller.CreateUserAuth)

			// Login - app issues JWT
			// - if cookie management is enabled, save tokens on client browser
			v1.POST("login", gcontroller.Login)

			// Logout
			// - if cookie management is enabled, delete tokens from cookies
			// - if Redis is enabled, save tokens in a blacklist until TTL
			rLogout := v1.Group("logout")
			rLogout.Use(gmiddleware.JWT()).Use(gmiddleware.RefreshJWT()).Use(gservice.JWTBlacklistChecker())
			rLogout.POST("", gcontroller.Logout)

			// Refresh - app issues new JWT
			// - if cookie management is enabled, save tokens on client browser
			rJWT := v1.Group("refresh")
			rJWT.Use(gmiddleware.RefreshJWT()).Use(gservice.JWTBlacklistChecker())
			rJWT.POST("", gcontroller.Refresh)

			// Double authentication
			if gconfig.Is2FA() {
				r2FA := v1.Group("2fa")
				r2FA.Use(gmiddleware.JWT()).Use(gservice.JWTBlacklistChecker())
				r2FA.POST("setup", gcontroller.Setup2FA)
				r2FA.POST("activate", gcontroller.Activate2FA)
				r2FA.POST("validate", gcontroller.Validate2FA)
				r2FA.POST("validate-backup-code", gcontroller.ValidateBackup2FA)

				// r2FA.Use(gmiddleware.TwoFA(
				// 	configure.Security.TwoFA.Status.On,
				// 	configure.Security.TwoFA.Status.Off,
				// 	configure.Security.TwoFA.Status.Verified,
				// ))
				// get 2FA backup codes
				r2FA.POST("create-backup-codes", gcontroller.CreateBackup2FA)
				// disable 2FA
				r2FA.POST("deactivate", gcontroller.Deactivate2FA)
			}

			// Update/reset password
			rPass := v1.Group("password")
			// Reset forgotten password
			if gconfig.IsEmailService() {
				// send password recovery email
				rPass.POST("forgot", gcontroller.PasswordForgot)
				// recover account and set new password
				rPass.POST("reset", gcontroller.PasswordRecover)
			}
			rPass.Use(gmiddleware.JWT()).Use(gservice.JWTBlacklistChecker())
			// if gconfig.Is2FA() {
			// 	rPass.Use(gmiddleware.TwoFA(
			// 		configure.Security.TwoFA.Status.On,
			// 		configure.Security.TwoFA.Status.Off,
			// 		configure.Security.TwoFA.Status.Verified,
			// 	))
			// }
			// change password while logged in
			rPass.POST("edit", gcontroller.PasswordUpdate)

			// Change existing email
			rEmail := v1.Group("email")
			rEmail.Use(gmiddleware.JWT()).Use(gservice.JWTBlacklistChecker())
			// if gconfig.Is2FA() {
			// 	rPass.Use(gmiddleware.TwoFA(
			// 		configure.Security.TwoFA.Status.On,
			// 		configure.Security.TwoFA.Status.Off,
			// 		configure.Security.TwoFA.Status.Verified,
			// 	))
			// }
			// add new email to replace the existing one
			rEmail.POST("update", gcontroller.UpdateEmail)
			// retrieve the email which needs to be verified
			rEmail.GET("unverified", gcontroller.GetUnverifiedEmail)
			// resend verification code to verify the modified email address
			rEmail.POST("resend-verification-email", gcontroller.ResendVerificationCodeToModifyActiveEmail)

			// User
			rUsers := v1.Group("users")
			rUsers.GET("", controller.GetUsers)    // Non-protected
			rUsers.GET("/:id", controller.GetUser) // Non-protected
			rUsers.Use(gmiddleware.JWT()).Use(gservice.JWTBlacklistChecker())
			// if gconfig.Is2FA() {
			// 	rUsers.Use(gmiddleware.TwoFA(
			// 		configure.Security.TwoFA.Status.On,
			// 		configure.Security.TwoFA.Status.Off,
			// 		configure.Security.TwoFA.Status.Verified,
			// 	))
			// }
			rUsers.POST("", controller.CreateUser) // Protected
			rUsers.PUT("", controller.UpdateUser)  // Protected
			// rUsers.PUT("/hobbies", controller.AddHobby) // Protected

			// Log
			rLogs := v1.Group("logs")
			rLogs.Use(gmiddleware.JWT()).Use(gservice.JWTBlacklistChecker())
			rLogs.POST("", controller.CreateLog) // Protected

			// Test JWT
			// rTestJWT := v1.Group("test-jwt")
			// rTestJWT.Use(gmiddleware.JWT()).Use(gservice.JWTBlacklistChecker())
			// if gconfig.Is2FA() {
			// 	rTestJWT.Use(gmiddleware.TwoFA(
			// 		configure.Security.TwoFA.Status.On,
			// 		configure.Security.TwoFA.Status.Off,
			// 		configure.Security.TwoFA.Status.Verified,
			// 	))
			// }
			// rTestJWT.GET("", controller.AccessResource) // Protected
		}
	}

	return r, nil
}
