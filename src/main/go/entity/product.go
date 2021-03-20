package product

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// 商品エンティティ
type Product struct {
	gorm.Model
	Name  string
	Price int
}

func NewProduct(name string, price int) *Product {
	p := new(Product)
	p.Name = name
	p.Price = price

	return p
}

// DBを初期化します
func DbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。（init）")
	}
	db.AutoMigrate(&Product{})
	defer db.Close()
}

// 商品を追加 or 編集します
func (p Product) Save() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic(err)
	}
	var product Product
	result := db.First(&product, "name = ?", p.Name)

	if result.Error != nil {
		db.Create(&p)
	} else {
		product.Name = p.Name
		product.Price = p.Price
		db.Save(product)
	}
	defer db.Close()
}

// 商品を削除します
func Delete(id int) {
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
func FindAll() []Product {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースが開けませんでした。(getAll)")
	}
	var products []Product
	db.Order("created_at desc").Find(&products)
	db.Close()
	return products
}
