package Router

import (
	"github.com/gin-gonic/gin"
	"log"
	"mirror/common"
	"net/http"
	"mirror/controllers"
)

func InitRouter() {
	router := gin.Default()
	// 无需鉴权
	user := router.Group("user")
	{
		user.POST("/login", controllers.Login)
		user.POST("/register", controllers.Register)
		user.GET("/getPassword", controllers.GetPassword)
		user.GET("/getRegisterCode", controllers.GetRegisterCode)
	}

	// 需要鉴权
	router.Use(ValidateTokenMiddleware())
	userRouter := router.Group("user")
	{
		userRouter.POST("/check", controllers.Check)
		userRouter.POST("/editInfo", controllers.EditInfo)
	}

	router.Run(":8626")
}

func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Token
		token := ctx.Request.Header.Get("Authorization")
		log.Println(token)
		if string(token) == "" {
			log.Println("auth is fail1")
			return
		}

		//defer func() {
		//	if p := recover(); p != nil {
		//		log.Println("auth is fail2")
		//	}
		//}()

		rs, ok := common.ParseToken(string(token), "token")
		if ok != true {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    1000,
				"message": "Unauthorized",
			})

			ctx.Abort()
		}

		ctx.Set("user", rs)
		ctx.Next()

	}
}
