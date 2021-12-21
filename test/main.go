package main

import (
	"fmt"
	"os"
	"reflect"
)

type User struct {
	Id int
	Name string
}

func ChangeSlice(s interface{}) {
	sT := reflect.TypeOf(s)
	//fmt.Println(sT.Kind())

	if sT.Kind() != reflect.Ptr {
		fmt.Println("参数必须是ptr类型")
		os.Exit(-1)
	}

	// 取得数组中元素的类型
	sTEE := sT.Elem().Elem()
	fmt.Println(sT.Elem())
	fmt.Println(sTEE)

	// 数组的值
	sV := reflect.ValueOf(s)
	sVE := sV.Elem()
	fmt.Println(sV)
	fmt.Println(sVE)

	// new一个数组中的元素对象
	sON := reflect.New(sTEE)
	// 对象的值
	sONE := sON.Elem()
	// 给对象复制
	sONEId := sONE.FieldByName("Id")
	sONEName := sONE.FieldByName("Name")
	sONEId.SetInt(10)
	sONEName.SetString("李四")

	// 创建一个新数组并把元素的值追加进去
	newArr := make([]reflect.Value, 0)
	newArr = append(newArr, sON.Elem())

	// 把原数组的值和新的数组合并
	resArr := reflect.Append(sVE, newArr...)

	// 最终结果给原数组
	sVE.Set(resArr)
}

func main() {
	var users []User
	ChangeSlice(&users)
	// 这里希望让Users指向ChangeSlice函数中的那个新数组
	fmt.Println(users)


}
