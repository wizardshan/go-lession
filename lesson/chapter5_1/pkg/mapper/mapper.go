package mapper

import "github.com/jinzhu/copier"

func Map(toValue interface{}, fromValue interface{}) {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		panic(err)
	}
}
