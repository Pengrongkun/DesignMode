package main

import "designmode/template"

func main()  {
	//s1,err1:=singleton.GetInstance3(1)
	//if err1!=nil{
	//	panic(err1)
	//}
	//fmt.Println(s1.Name)
	//s1.Name="ppz"
	//
	//s1,err2:=singleton.GetInstance3(1)
	//if err2!=nil{
	//	panic(err2)
	//}
	//fmt.Println(s1.Name)
	//
	//fmt.Println(singleton.GetRandomInstance3().Name)
	//fmt.Println(singleton.GetRandomInstance3().Name)
	//fmt.Println(singleton.GetRandomInstance3().Name)
	//b:=factory.NewResourcePoolingBuilder()
	//rp:=b.AddName("莫大").AddMaxIdle(16).AddMaxTotal(8).AddMinIdle(2).Build()
	//if rp!=nil{
	//	fmt.Println(*rp)
	//}else {
	//	fmt.Println("创建失败")
	//}

	c:=template.ConcreteTemplate{
		AbstractTemplate: template.AbstractTemplate{
			Template: nil,
		},
	}
	c.TemplateMethod()

}
