// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 16:31
package model

import (
	"github.com/jinzhu/gorm"
)

// Order
type Order struct {
	gorm.Model
	Oid  int
	Time int64
	Uid  int
	Mid  int
}

// AddOrder
func AddOrder(o Order) error {
	return DB.Create(&o).Error
}
