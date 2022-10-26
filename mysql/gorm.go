package mysql

import (
	"fmt"
	"gorm.io/gorm"
)

func NewGormClient() {
	db := gorm.DB{}
	fmt.Println(db)
}
