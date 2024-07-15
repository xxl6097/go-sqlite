package main

import (
	"fmt"
	"go-sqlite/model"
	"go-sqlite/sqlite"
)

func main() {
	db := sqlite.InitMysql("sqlite.db")
	repo := model.NewTestRepository(db)
	repo.Add(model.TestModel{
		CardCode:    "001-cardcode",
		UserName:    "001-uuxia",
		Points:      100,
		ExpiresType: 9,
	})
	arr, err := repo.FindAll()
	fmt.Println(err, arr)
	fmt.Println(repo)
}
