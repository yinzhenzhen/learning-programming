/**
 * @Time : 2020-05-26 13:20
 * @Author : yz
 */

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func decodeAtIndex(S string, K int) string {

	var newS string

	for i := 0; i < len(S); i++ {
		str := fmt.Sprintf("%c", S[i])
		// 如果是字母，写入newS
		if str >= "a" && str <= "z" {
			newS = newS + str
		// 如果是数字，进行重复
		} else {
			count, _ := strconv.Atoi(str)
			temp := newS
			for count > 1 {
				// 运行环境，panic
				newS = newS + temp
				count--
			}

			if len(newS) >= K {
				//fmt.Println(newS)
				return fmt.Sprintf("%c", newS[K-1])
			}
		}
	}

	//fmt.Println(newS)

	return fmt.Sprintf("%c", newS[K-1])
}

func main()  {
	//fmt.Println(decodeAtIndex("y959q969u3hb22odq595",222280369))

	demo := Demo{Name:"22", Addr:"33", School:"44"}

	fmt.Println(demo.Age)

	//setByTags(demo)
	setByPoint(demo)
}

// ------------------------------------------------------------------------------

type Demo struct {
	Age int `json:"Age"`
	Name string `json:"Name"`
	Addr string `json:"Addr"`
	School string `json:"School"`
}

// 通过反射改变值
func setByTags(demo Demo) {
	fmt.Println("before : ", demo)

	// 设置结构体值需要获取的
	s := reflect.ValueOf(&demo).Elem()

	// 类型
	t := reflect.TypeOf(demo)

	// 字段数量
	num := t.NumField()

	for i := 0; i < num; i++ {
		f := t.Field(i)
		tagName := f.Tag.Get("json")
		fType := f.Type.Kind()

		switch fType {
		case reflect.Int:
			s.FieldByName(tagName).SetInt(11111)
		case reflect.String:
			s.FieldByName(tagName).SetString("yy")
		}

		fmt.Printf("%d: %s %s = %v\n", i, f.Name, fType.String(), s.FieldByName(tagName).String())
	}

	fmt.Println("after : ", demo)
}

// 通过指针改变值
func setByPoint(demo Demo)  {
	fmt.Println("before : ", demo)

	// 结构体起始位置
	demoPointer := unsafe.Pointer(&demo)

	// 类型
	t := reflect.TypeOf(demo)

	// 字段数量
	num := t.NumField()

	for i := 0; i < num; i++ {
		f := t.Field(i)
		fType := f.Type.Kind()

		fmt.Println(f.Index)

		// 偏移量
		offset := f.Offset

		switch fType {
		case reflect.Int:
			fieldPointer := (*int)(unsafe.Pointer(uintptr(demoPointer) + offset))
			*fieldPointer = 22222
		case reflect.String:
			fieldPointer := (*string)(unsafe.Pointer(uintptr(demoPointer) + offset))
			*fieldPointer = "zz"
		}
		fmt.Printf("%d: %s %s = %v\n", i, f.Name, fType.String(), offset)
	}

	fmt.Println("after : ", demo)
}

