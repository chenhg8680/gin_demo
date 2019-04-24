package routers

import (
	"blog/pkg/setting"
	"blog/routers/api"
	"blog/routers/api/v1"
	"blog/routers/page"

	"blog/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())

	{
		//---标签---
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		//新建标签
		apiv1.POST("/tags", v1.AddTag)

		//更新制定标签
		apiv1.PUT("/tags/:id", v1.EditTag)

		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//---文章---
		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.GET("/articles", v1.GetArticles)

		apiv1.POST("/articles", v1.AddArticle)

		apiv1.PUT("/articles/:id", v1.EditArticle)

		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	//加载html
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", page.WebsiteIndex)

	return r
}
