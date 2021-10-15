package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func Judge(i interface{}) {
	fmt.Println(reflect.TypeOf(i))        // *main.User
	fmt.Println(reflect.TypeOf(i).Elem()) // main.User
	for index := 0; index < reflect.TypeOf(i).Elem().NumField(); index++ {
		fmt.Println(reflect.TypeOf(i).Elem().Field(index))
	}
	v := reflect.ValueOf(i) // &{Liu 100}
	fmt.Println(v)          // {Liu 100}
	ve := v.Elem()
	fmt.Println(ve)
	for index := 0; index < ve.NumField(); index++ {
		fmt.Println(ve.Field(index)) // Liu 100
	}

	fmt.Println(ve.FieldByName("Age"))             // int类型
	fmt.Println(ve.FieldByName("Age").Interface()) // 接口类型

	objVal := reflect.Indirect(v)
	fmt.Println(objVal) //

	fmt.Println("----------------------------")
	var users []User
	users = append(users, User{Name: "l2", Age: 12})
	fmt.Println(reflect.ValueOf(users))
	fmt.Println(reflect.ValueOf(users).Type().Kind())
	val := reflect.Indirect(reflect.ValueOf(users))
	fmt.Println(val)
	fmt.Println(val.Type())        // slice type []User
	fmt.Println(val.Type().Elem()) // Elem User
	u1 := reflect.New(val.Type().Elem())
	if u2, ok := u1.Interface().(*User); ok {
		u2.Age = 13
		fmt.Printf("%d\n", u2.Age)
	}

}

func main() {
	u := User{
		Name: "Liu",
		Age:  100,
	}
	Judge(&u)
	fmt.Println("vim-go")
}
