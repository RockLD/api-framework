package model

import (
	"sync"
	"time"
)

//数据库基类
type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"_"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"_"`
	UpdateAt  time.Time  `gorm:"column:updateAt" json:"_"`
	DeleteAt  *time.Time `gorm:"column:deleteAt" sql:"index" json:"_"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}
