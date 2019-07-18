package model

import (
	"api/pkg/auth"
	validator "gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string
	Password string
}

//表名
func (u *UserModel) TableName() string {
	return "tb_users"
}

//Create creates a new user account
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
