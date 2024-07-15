package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code        string
	Price       uint
	ApiKey      string `json:"apikey" gorm:"column:apikey;unique;not null;comment:'密钥'"`
	UseCount    int64  `json:"usecount" gorm:"column:usecount;default:0;comment:'已经使用次数'"`
	MaxCount    int64  `json:"maxcount" gorm:"column:maxcount;not null;default:10;comment:'最大使用次数'"`
	ExpiresType int    `json:"expirestype" gorm:"column:expirestype;default:0;comment:'类型'"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}
