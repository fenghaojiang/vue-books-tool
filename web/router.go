package web

import (
	"demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	r := gin.Default()
	r.GET("/douban", func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		model.BookQuery(context, "SELECT * FROM douban_book") //查出所有结果返回
	})
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exist",
		})
	})
	r.Run(":8000")
}
