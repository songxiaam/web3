package models

import (
	"fmt"
	"gin-study/database"
	"gin-study/utils"
)

type User struct {
	ID          int    `json:"id"`
	LoginName   string `json:"login_name"`
	NickName    string `json:"nick_name"`
	AvatarUrl   string `json:"avatar_url"`
	PhoneNumber string `json:"phone_number"`
}

func (User) TableName() string {
	return "tb_uvs_user"
}

func FindUserByLoginNameAndPassword(loginName, password string) (*User, error) {
	var user User
	str, _ := utils.Encrypt(utils.SecretKey, password)
	md5Password := utils.MD5Hash("EN:" + str)
	fmt.Println(md5Password)
	md5Password = password
	result := database.UvsDB.Where("login_name = ?", loginName).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
