package user

import (
	"mirror/models"
	"mirror/common"
	"log"
	"mirror/Databases"
)

/**
 执行登录
 */
func DoLogin(userName string, password string, loginType uint8) interface{} {

	var user models.User
	var pass string
	var userId int
	if loginType == 2 {

		user, err := user.FindByName(userName)
		if err != nil {
			return nil
		}
		userName = user.Email
		pass = user.Password
		userId = user.Id
	} else {
		user, err := user.FindByEmail(userName)
		log.Println(user)
		if err != nil {
			log.Println(err)
			return nil
		}
		userId = user.Id
		pass = user.Password

	}

	if pass != common.Encode(password) {
		log.Println(common.Encode(password))
		log.Println(password)
		log.Println(pass)
		return nil
	}

	return afterLogin(userName, userId)
}

func GetPassword(email string) {
	var pass = common.GenValidateCode(6)

	body := `您的密码：<br> ` + pass + `
            <h3></h3>
             <br>`

	common.SendEmail(email, "【心魔】找回密码", body)
}

func GetRegisterCode(email string) bool {
	var pass = common.GenValidateCode(6)

	var register models.Register
	register.Email = email
	register.Code = pass

	_, err := register.FindByEmail(email)
	if err == nil {

		return false
	}

	register.Insert()
	body := `您的注册码：<br> ` + pass + `
            <h3></h3>
             <br>`

	common.SendEmail(email, "【心魔】欢迎来到魔界", body)

	return true
}

func DoRegister(email string, code string) interface{} {
	var register models.Register
	c, err := register.FindByEmail(email)

	if err != nil || c.Code != code {
		return 1
	}

	var user models.User

	_, err = user.FindByEmail(email)
	if err == nil {
		return 2
	}
	user.Password = common.Encode(code)
	user.Email = email
	user.Insert()
	log.Println(user.Id)

	response := afterLogin(email, user.Id)
	return response
}

func afterLogin(email string, id int) interface{} {
	var info = make(map[string]interface{})
	info["email"] = email
	info["userId"] = id

	token := common.CreateToken("token", info)
	var response = make(map[string]interface{})
	response["token"] = token

	return response
}

func EditUser(u models.UserInfo, info map[string]interface{}) {
	log.Println(info["userId"])
	var user models.User
	id := int(info["userId"].(float64))
	res := user.FindById(id)

	if u.Name != "" && u.Name != res.Name {
		res.Name = u.Name
	}

	if u.Password != "" {
		ps := common.Encode(u.Password)

		if ps != u.Password {
			res.Password = ps
		}
	}

	Databases.DB.Save(&res)
}
