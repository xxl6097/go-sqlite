package model

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	CardCode    string `json:"cardcode" gorm:"column:cardcode;unique;not null;comment:'兑换码'"`
	UserName    string `json:"username" gorm:"column:username;default:null;comment:'注册用户'"`
	Points      int64  `json:"points" gorm:"column:points;default:10;not null;comment:'积分'"`
	ExpiresType int    `json:"expirestype" gorm:"column:expirestype;default:0;comment:'类型'"`
}

func (this *TestModel) Response() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["cardcode"] = this.CardCode
	resp["username"] = this.UserName
	resp["points"] = this.Points
	resp["expirestype"] = this.ExpiresType
	return resp
}
