package handler

import (
	"net/http"

	"github.com/floriwan/srcm/pkg/config"
	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": config.GlobalConfig.HomepageName,
	})
}
