// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 16:30
package model

import (
	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	Mid     int
	Name    string
	time    string
	Address string
	Num     int
}

func RemoveMovie(mid int) error {

}

func GetMovieNum(mid int) (int, error) {
	var m Movie

	err := DB.Table("movies").Where("mid = ?", mid).First(&m).Error
	return m.Num, err
}
