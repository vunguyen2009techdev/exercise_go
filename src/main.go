package src

import (
	"exercise_go/src/services"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type handler struct {
	translator services.Translator
}

func newHandler(db *xorm.Engine) *handler {
	return &handler{
		translator: services.NewTranslatorService(db),
	}
}

func RegisterRoutes(r *gin.Engine, db *xorm.Engine) {
	h := newHandler(db)

	routes := r.Group("/api/translations")
	routes.GET("/", h.translator.GetTranslates)
}
