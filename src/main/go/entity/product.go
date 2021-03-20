package entity

import (
	"github.com/jinzhu/gorm"
	"github.com/mattn/go-sqlite3"
)

// 商品エンティティ
type Product struct {
	gorm.Model
	Name  string
	Price int
}

// DBを初期化します
func init() {
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

// 商品を更新します
func update(id int, name string, price int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。（update)")
	}
	var product Product
	db.First(&product, id)
	product.Name = name
	product.Price = price
	db.Save(&product)
	db.Close()
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
func getAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。(getAll)")
	}
	var products []Product
	db.Order("created_at desc").Find(&products)
	db.Close()
	return products
}
