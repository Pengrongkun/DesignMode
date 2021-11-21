package template

import "fmt"

type Template interface {
	TemplateMethod()
	Method1()
	Method2()
}
//抽象结构体
type AbstractTemplate struct {
	Template
}

func (a AbstractTemplate)TemplateMethod(){
	//骨架...
	fmt.Println("骨架1")
	//子类实现method1
	a.Method1()
	//骨架...
	fmt.Println("骨架2")
	//子类实现method2
	a.Method2()
	//骨架...
	fmt.Println("骨架3")
}

//实际
type ConcreteTemplate struct {
	AbstractTemplate
}

func (c ConcreteTemplate)Method1()  {
	fmt.Println("子方法1")
}

func (c ConcreteTemplate)Method2()  {
	fmt.Println("子方法2")
}