package api

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
)

type auth struct {
	Username string `valid:"Requird;MaxSize(50)"`
	Password string `valid:"Requird;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}

	a := auth{Username: username, Password: password}

	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})

	code := e.INVALID_PARAMS

	//log.Panicf("%s,%s", username, password)
	//log.Panicln(ok, username, password)

	if ok {
		isExist := models.CheckAuth(username, password)

		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				code = e.SUCCESS
				data["token"] = token
			}

		} else {
			code = e.ERROR_AUTH
		}

	} else {
		for _, err := range valid.Errors {
			log.Panicln(err.Key, err.Message)
			log.Panicf("err key:%s,err msg:%s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
