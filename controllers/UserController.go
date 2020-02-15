package controllers

import (
	"github.com/gin-gonic/gin"
	"mirror/services/user"
	"log"
	"net/http"
	"mirror/models"
	"github.com/dgrijalva/jwt-go"
)

type UserLogin struct {
	Name     string
	Password string
	Type     uint8
}

type UserRegister struct {
	Email string
	Code  string
}

func Login(c *gin.Context) {
	var p UserLogin
	c.BindJSON(&p)
	log.Println(p)
	res := user.DoLogin(p.Name, p.Password, p.Type)
	if res != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data":    res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    2001,
			"message": "用户名或密码错误",
		})
	}
}

func Register(c *gin.Context) {
	var p UserRegister
	c.BindJSON(&p)
	log.Println(p.Email)
	res := user.DoRegister(p.Email, p.Code)

	if res == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    2000,
			"message": "口令错误",
		})
	} else if res == 2 {
		c.JSON(http.StatusOK, gin.H{
			"code":    2003,
			"message": "邮箱已注册",
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "成功加入",
			"data":    res,
		})
	}

}

func GetPassword(c *gin.Context) {
	email := c.Query("email")

	user.GetPassword(email)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}

func GetRegisterCode(c *gin.Context) {
	email := c.Query("email")

	res := user.GetRegisterCode(email)

	if res == true {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    2003,
			"message": "邮箱已注册",
		})
	}

}

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}

func EditInfo(c *gin.Context) {
	var u models.UserInfo
	c.BindJSON(&u)
	userInfo, ex := c.Get("user")
	if ex == true {
		us := userInfo.(jwt.MapClaims)
		user.EditUser(u, us)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}
