package main

import (
	"fmt"
)

/*
	策略模式
	举例。去学校使用哪种交通方式是不同的策略；使用哪种方式提现是不同的策略；
*/

// ==================== 模式简单示例 ====================

// 定义策略接口
type Strategy interface {
	do()
}

// 策略A
type StrategyA struct{}

func (StrategyA) do() {}

// 策略B
type StrategB struct{}

func (StrategB) do() {}

// 策略使用对象
type Model struct {
	Strategy
}

// ==================== 模式代码举例 ====================

// 提现策略接口
type withdrawStrategy interface {
	withdraw(account string, money int) error
}

// 策略1: 支付宝提现
type Alipay struct {
	// ...
}

func (a *Alipay) withdraw(account string, money int) error {
	// 调用sdk
	fmt.Printf("提现宝账号：%s，提现金额：%d \n", account, money)
	return nil
}

// 策略2: 微信提现
type WeChat struct {
	// ...
}

func (w *WeChat) withdraw(account string, money int) error {
	// 调用sdk
	fmt.Printf("微信账号：%s，提现金额：%d \n", account, money)
	return nil
}

// 校验接口实现
var (
	_ withdrawStrategy = (*Alipay)(nil)
	_ withdrawStrategy = (*WeChat)(nil)
)

// 提现
type Withdraw struct {
	withdrawStrategy
	account string
	money   int
}

func (w *Withdraw) Do() error {
	return w.withdrawStrategy.withdraw(w.account, w.money)
}

func NewWithdraw(account string, money int, strategy withdrawStrategy) *Withdraw {
	return &Withdraw{
		withdrawStrategy: strategy,
		account:          account,
		money:            money,
	}
}

// ==================== 调用代码举例 ====================

func main() {
	var ws withdrawStrategy
	ws = &Alipay{}
	withdraw := NewWithdraw("110", 1, ws)
	withdraw.Do()

	ws = &WeChat{}
	withdraw = NewWithdraw("110", 1, ws)
	withdraw.Do()
}
