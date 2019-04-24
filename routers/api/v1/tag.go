package v1

import (
	"blog/pkg/logging"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	//参数检测
	valider := validation.Validation{}

	valider.Required(name, "name").Message("名称不能为空")
	valider.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valider.Required(createdBy, "created_by").Message("创建人不能为空")
	valider.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valider.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS

	if !valider.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, createdBy, state)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valider.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		//"msg":  name,
		"data": make(map[string]string),
	})

}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	name := c.Query("name")

	modifiedBy := c.Query("modified_by")

	valider := validation.Validation{}

	state := -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valider.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	//log.Println(name, state, modifiedBy)

	valider.Required(id, "id").Message("ID不能为空")
	valider.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valider.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valider.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS

	if !valider.HasErrors() {
		code = e.SUCCESS

		//标签存在
		if models.ExistTagByID(id) {
			//更新
			data := make(map[string]interface{})

			data["modified_by"] = modifiedBy

			if name != "" {
				data["name"] = name
			}

			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)

		} else {
			//标签不存在
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valider.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valider := validation.Validation{}

	valider.Required(id, "id").Message("ID不能为空")

	code := e.INVALID_PARAMS

	if !valider.HasErrors() {
		if models.ExistTagByID(id) {
			code = e.SUCCESS
			//删除
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valider.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
