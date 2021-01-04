package main

import "fmt"

// 测试如果struct未实现全部的interface接口，但是赋值给了interface，编译是否报错
type TestI interface {
	Test() error
}

type TestS struct {
	name string
}

// 知识点err可以提前被定义, 好处当执行defer时，不至于err未定义
func (t *TestS) Test() error {
	err := func() error {
		fmt.Errorf("Test error")
		return nil
	}()
	defer func() {
		fmt.Printf("%v\n", err)
	}()
	return err
}

var _ TestI = &TestS{}

type TestI1 interface {
	Test1() error
}

type TestS1 struct {
	TestI1
}

var t1 TestI1 = &TestS1{}

func main() {
	fmt.Println("vim-go")
	fmt.Printf("%v\n", t1)

	var t TestI = &TestS{}
	t.Test()
	r1, r2 := NameReturn()
	fmt.Printf("%s, %d\n", r1, r2)
}

//./main.go:14:5: cannot use &tests literal (type *tests) as type testi in assignment:
//     *tests does not implement testi (missing test method)

// 代码中的var _ Client = &ovndb{}, 目的是校验ovndb struct是否全部实现了Client接口的所有方法，否则编译报错。

// 如何保证struct不需要实现全部的interface 函数同样能够转换成Interface类型呢？将interface作为struct的匿名变量

// https://www.jianshu.com/p/a5bc8add7c6e, struct中嵌入结构体的解答

// 函数定义
//函数无须前置声明

//不支持命名嵌套定义，支持匿名嵌套

//函数只能判断是否为nil，不支持其它比较操作

//支持多返回值

//支持命名返回值

//支持返回局部变量指针

//支持匿名函数和闭包
func NameReturn() (r1 string, r2 int) {
	r1 = "r1"
	r2 = 2
	return r1, r2
}
