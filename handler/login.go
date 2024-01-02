package handler

import (
	"fmt"
	"net/http"

	"github.com/floriwan/srcm/pkg/config"
	"github.com/gin-gonic/gin"
)

func LoginPost(c *gin.Context) {
	fmt.Printf("http post\n")
}

func LoginGet(c *gin.Context) {

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title":    config.GlobalConfig.HomepageName,
		"subtitle": "login",
	})

}
