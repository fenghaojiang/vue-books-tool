package model

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Book struct {
	Id         int32
	Title      string
	Author     string
	Translator string
	Press      string
	Date       string
	Price      string
	Star       string
	Comment    string
	Quote      string
}

func BookQuery(context *gin.Context, query string) {
	var resultList []Book
	dsn := "root:fenghaojiang97@tcp(localhost:3306)/douban?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	rows, err2 := db.Query(query)
	if err2 == nil {
		for rows.Next() {
			var book Book
			rows.Scan(&book.Id, &book.Title, &book.Author, &book.Translator, &book.Press, &book.Date, &book.Price, &book.Star, &book.Comment, &book.Quote)
			resultList = append(resultList, book)
		}
	}
	context.JSON(http.StatusOK, resultList)
}
