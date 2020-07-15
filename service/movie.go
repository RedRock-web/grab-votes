// @program: grab-votes 
// @author: edte
// @create: 2020-07-15 22:26
package service

import (
	"grab-votes/logs"
	"grab-votes/model"
)

func RemoveMovie(mid int) {
	num, err := model.GetMovieNum(mid)
	if err != nil {
		logs.Error.Printf("failed to get movie num:%s", err)
		return
	}
	if num == 0 {
		return
	}

	err = model.RemoveMovie(mid)
	if err != nil {
		logs.Error.Printf("failed to remove movie:%s", err)
	}
}
