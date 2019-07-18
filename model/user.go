package model

import (
	"api/pkg/auth"
	"api/pkg/constvar"
	"fmt"
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

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Self.Save(&u).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
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
