package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

// 商品エンティティ
type Product struct {
	gorm.Model
	Name  string
	Price int
}

// DBを初期化します
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。（init）")
	}
	db.AutoMigrate(&Product{})
	defer db.Close()
}

// DBに商品を追加します
func insert(name string, price int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。（insert)")
	}
	db.Create(&Product{Name: name, Price: price})
	defer db.Close()
}

// 商品を削除します
func delete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。（delete)")
	}
	var product Product
	db.First(&product, id)
	db.Delete(&product)
	db.Close()
}

//　全商品を取得します
func getAll() []Product {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。(getAll)")
	}
	var products []Product
	db.Order("created_at desc").Find(&products)
	db.Close()
	return products
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*.html")

	dbInit()

	router.GET("/", func(ctx *gin.Context) {
		products := getAll()
		ctx.HTML(200, "index.html", gin.H{"products": products})
	})

	router.POST("/create", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		price, err := strconv.Atoi(ctx.PostForm("price"))
		if err != nil || len(name) == 0 {
			ctx.Redirect(302, "/")
		}
		insert(name, price)
		ctx.Redirect(302, "/")
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		name := ctx.PostForm("name")
		price, err := strconv.Atoi(ctx.PostForm("price"))
		if err != nil {
			panic("ERROR")
		}
		update(id, name, price)
		ctx.Redirect(302, "/")
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		delete(id)
		ctx.Redirect(302, "/")

	})

	router.Run()
}
