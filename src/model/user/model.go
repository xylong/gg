package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id" form:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name      string         `json:"name" form:"name" binding:"min=4" gorm:"column:name;type:varchar(20);"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;type:datetime;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:datetime;index"`
}

func NewUser(attrs ...UserAttrFunc) *User {
	user := &User{}
	UserAttrFuncs(attrs).Apply(user)
	return user
}
