package main

import "fmt"

/*
	选项模式
	New时支持传递不确定个参数，不区分顺序，有较好的兼容性
*/

// ==================== 模式代码 ====================

type Robot struct {
	Name    string
	Age     int
	Address string
	// ....
}

type Option func(r *Robot)

func WithAge(age int) Option {
	return func(r *Robot) {
		r.Age = age
	}
}

func WithAddress(address string) Option {
	return func(r *Robot) {
		r.Address = address
	}
}

func NewRobot(name string, ops ...Option) *Robot {
	r := &Robot{
		Name:    name,
		Address: "黑龙江",
	}

	for i := range ops {
		ops[i](r)
	}

	return r
}

// ==================== 测试代码 ====================

func main() {
	r1 := NewRobot("艾AA", WithAddress("浙江"), WithAge(10))
	fmt.Println(r1.Name, r1.Address, r1.Age)
	r2 := NewRobot("艾BB")
	fmt.Println(r2.Name, r2.Address, r2.Age)
}
