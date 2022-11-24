package main

import (
	"fmt"
	"sync"
)

/*
	单例模式
	声明一个包内私有变量，只通过Get函数暴露此唯一变量
	可以分为懒汉式和饿汉式。（个人觉得两种方式主要影响本地开发，例如某第三方包账号连接，推荐懒汉式）
		1. 饿汉式在init函数中加载。消耗的是程序启动的时间，
		2. 懒汉式在获取唯一对象时使用sync.Once现加载。第一次获取耗时。
*/

// ==================== 模式代码 ====================

// 单例接口。(个人觉得声明单例接口不是很有必要，GetInstance直接返回实例就好)
type Singleton interface {
	init()
}

// 举例的单例对象
type King struct{}

func (k *King) init() {
	// 假设耗时30ms
	k = &King{}
}

var (
	instance *King
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance.init()
	})

	return instance
}

// ==================== 使用 ====================

func main() {
	fmt.Println(instance == nil)
	k1 := GetInstance()
	fmt.Println(k1 != nil)
	k2 := GetInstance()
	fmt.Println(k1 == k2)

}
