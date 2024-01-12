package handler

import (
	"github.com/floriwan/srcm/pkg/config"
	"github.com/floriwan/srcm/pkg/templates"
	"gorm.io/gorm"
)

type (
	Handler struct {
		DB     *gorm.DB
		Config config.Config
		Tmpl   *templates.Tmpl
	}
)
