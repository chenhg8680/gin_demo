package page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebsiteIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
