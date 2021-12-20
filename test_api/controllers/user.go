package controllers

import (
	// "fmt"
	"encoding/json"
	"test_api/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users

type UserController struct {
	beego.Controller
}

// @Title CreateUser

// @Description create users

// @Param body body models.User true "body for user content"

// @Success 200 {int} models.User.Id

// @Failure 403 body is empty

// @router / [post]

func (u *UserController) Post() {

	// u.TplName = "index.tpl"
	// u.Data["first_name"] = u.GetString("first_name")
	// u.Data["last_name"] = u.GetString("last_name")
	// u.Data["email"] = u.GetString("email")
	// u.Data["phone"] = u.GetString("phone")
	// u.Data["password"] = u.GetString("password")
	// u.Data["dob"] = u.GetString("dob")
	// fmt.Println(u.Data["first_name"], u.Data["last_name"], u.Data["email"], u.Data["phone"], u.Data["password"], u.Data["dob"])

	var user models.User

	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	uid := models.AddUser(user)

	u.Data["json"] = map[string]string{"api": uid}

	u.ServeJSON()

}

// @Title GetAll

// @Description get all Users

// @Success 200 {object} models.User

// @router / [get]

func (u *UserController) GetAll() {

	users := models.GetAllUsers()

	u.Data["json"] = users

	u.ServeJSON()

}
