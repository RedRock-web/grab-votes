// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 16:31
package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Oid  int
	time string
	Uid  int
	Mid  int
}
