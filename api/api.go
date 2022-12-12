package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"app/api/handler"
)

func SetUpApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandlerV1(db)

	r.POST("/user", handlerV1.Create)
	r.GET("/user/:id", handlerV1.GetById)
	r.GET("/user", handlerV1.GetList)
	r.DELETE("/user/:id",handlerV1.Delete)
	r.PUT("/user",handlerV1.Update)
	r.PATCH("/user",handlerV1.Patch)
}
