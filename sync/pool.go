package main

import (
	"fmt"
	"sync"
)

// 用途：
//sync.Pool是一个可以存或取的临时对象集合
//sync.Pool可以安全被多个线程同时使用，保证线程安全
//注意、注意、注意，sync.Pool中保存的任何项都可能随时不做通知的释放掉，所以不适合用于像socket长连接或数据库连接池。
//sync.Pool主要用途是增加临时对象的重用率，减少GC负担。

// 使用方法
// 1. 声明一个变量为sync.Pool{}实现New方法
// 2. 调用Get方法获取临时数据
// 3. 调用类型转换操作，将interface{}转成特定处理类型
// 4. 将临时数据放回到池子中

// 保存for循环100次的结果，然后依次打印出来

var pool = sync.Pool{
	New: func() interface{} {
		fmt.Println("make")
		return make([]int, 100)
	},
}

func DoFor(wg *sync.WaitGroup) {
	defer wg.Done()
	item := pool.Get()

	fmt.Println(item)
	// 类型转换
	for i := 0; i < len(item.([]int)); i++ {
		item.([]int)[i] = i
	}

	var st string
	// 统一处理结果
	for i := 0; i < len(item.([]int)); i++ {
		st = fmt.Sprintf("%s%d", st, item.([]int)[i])
	}

	fmt.Println(st)
	pool.Put(item)
	return
}
