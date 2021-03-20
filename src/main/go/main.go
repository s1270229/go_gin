package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/src/main/go/entity"
	"strconv"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*.html")

	product.DbInit()

	router.GET("/", func(ctx *gin.Context) {
		products := product.FindAll()
		ctx.HTML(200, "index.html", gin.H{"products": products})
	})

	router.POST("/create", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		price, err := strconv.Atoi(ctx.PostForm("price"))
		if err != nil {
			panic("ERROR")
		}
		newProduct := product.NewProduct(name, price)
		newProduct.Save()

		ctx.Redirect(302, "/")
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		product.Delete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
