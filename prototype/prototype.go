package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	原型模式
	克隆一个相同的对象，被克隆的对象就叫做原型。
	这大概是最简单的模式了
*/

// ==================== 模式代码 ====================

type Robot struct {
	ID   int    // 偶尔有那么一个唯一值
	Name string // 克隆属性
}

func NewRobot(name string) *Robot {
	return &Robot{
		ID:   rand.Intn(1000),
		Name: name,
	}
}

func (r *Robot) Clone() *Robot {
	c := *r
	c.ID = rand.Intn(1000)
	return &c
}

// ==================== 测试代码 ====================

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	a := NewRobot("小明")
	fmt.Println(a.ID, a.Name)

	b := a.Clone()
	fmt.Println(b.ID, b.Name)
}
