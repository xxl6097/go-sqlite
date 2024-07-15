package model

import (
	"fmt"
	"gorm.io/gorm"
)

type TestRepository struct {
	db *gorm.DB
}

func NewTestRepository(db *gorm.DB) *TestRepository {
	//if db.Migrator().HasTable(&TestModel{}) {
	//	db.Debug().Migrator().CreateTable(&TestModel{})
	//}
	err := db.Debug().AutoMigrate(&TestModel{})
	if err != nil {
		fmt.Println("TestModel created failed", err)
	} else {
		fmt.Println("TestModel created")
	}
	return &TestRepository{
		db: db,
	}
}

func (this *TestRepository) AddJson(v map[string]interface{}) error {
	return this.db.Debug().Model(TestModel{}).Create(&v).Error
}

func (this *TestRepository) Add(testModel TestModel) error {
	return this.db.Create(&testModel).Error
}

func (this *TestRepository) Update(key map[string]interface{}) error {
	return this.db.Debug().Model(TestModel{}).Omit("cardcode").Where("cardcode = ?", key["cardcode"]).Updates(key).Error
}

func (this *TestRepository) Delete(Key TestModel) error {
	return this.db.Where("cardcode = ?", Key.CardCode).Unscoped().Delete(&Key).Error
}

func (this *TestRepository) Find(nmodel TestModel) (*TestModel, error) {
	var model TestModel
	err := this.db.
		Debug().
		Model(&TestModel{}).
		Where(&nmodel).
		Take(&model).Error
	return &model, err
}

func (this *TestRepository) FindAll() (*[]TestModel, error) {
	var users []TestModel
	err := this.db.Find(&users).Error
	return &users, err
}
