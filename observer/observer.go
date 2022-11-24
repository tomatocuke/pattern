package main

import "fmt"

/*
	观察者模式
	用于触发联动，一个对象的改变会触发其观察者的相关动作。
	通常是一个被观察者对应多个观察者，观察者作为一个接口各自实现逻辑。
	个人理解，该模式的观察者与被观察者之间，不应该具有强相关性，而是弱相关性。
*/

// ==================== 模式代码 ====================

// 声明观察者接口
type Observer interface {
	Do()
}

// 被观察者
type Subject struct {
	observers []Observer
}

// 添加观察者
func (s *Subject) Add(o Observer) {
	s.observers = append(s.observers, o)
}

// 通知订阅者
func (s *Subject) Notify(id int, msg string) {
	for i := range s.observers {
		s.observers[i].Do()
	}
}

// ==================== 测试代码 ====================

// 观察者-图书馆相关人员
type LibraryMember interface {
	listenLibrary(*Library)
}

// 被观察者-图书馆
type Library struct {
	members []LibraryMember
	IsOpen  bool
}

// 图书馆添加成员
func (l *Library) AddMember(m ...LibraryMember) {
	l.members = append(l.members, m...)
}

// 图书馆发布通知
func (l *Library) Notify() {
	for i := range l.members {
		l.members[i].listenLibrary(l)
	}
}

// 观察者-学生
type Student struct {
}

func (s *Student) listenLibrary(l *Library) {
	fmt.Println("学生收到通知，图书馆状态:", l.IsOpen)
}

// 观察者-管理员
type Manager struct {
}

func (m *Manager) listenLibrary(l *Library) {
	fmt.Println("管理员收到通知，图书馆状态:", l.IsOpen)
}

func main() {
	library := &Library{}
	s := &Student{}
	m := &Manager{}
	library.AddMember(s, m)

	library.IsOpen = false
	library.Notify()
	library.IsOpen = true
	library.Notify()
}
