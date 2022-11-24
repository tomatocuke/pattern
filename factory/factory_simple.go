package simplefactory

/*
	简单工厂方法
	通过对一个工厂方法传递参数控制返回不同对象
*/

// ==================== 模式代码 ====================

type API interface {
	Do()
}

type A struct{}

type B struct{}

func (A) Do() {}

func (B) Do() {}

// 工厂方法
func NewAPI(t int) API {
	if t == 1 {
		return &A{}
	} else if t == 2 {
		return &B{}
	}
	return nil
}
