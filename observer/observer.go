package observer

import "fmt"

//被观察者
type Subject interface {
	//注册观察者，观察改对象
	registerObserver(observer Observer)
	//删除观察者
	removeObserver(observer Observer)
	//通知观察者
	notifyObserver(message string)
}
//观察者
type Observer interface {
	update(message string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (c *ConcreteSubject) removeObserver(observer Observer) {
	for i,o:=range c.observers{
		if o==observer{
			c.observers=append(c.observers[:i],c.observers[i+1:]...)
		}
	}
}

func (c *ConcreteSubject) notifyObserver(message string) {
	for _,r:=range c.observers{
		r.update(message)
	}
}

func (c *ConcreteSubject)registerObserver(observer Observer){
	c.observers=append(c.observers, observer)
}

type ConcreteObserver struct {

}

func (c ConcreteObserver) update(message string) {
	fmt.Println(message)
}
