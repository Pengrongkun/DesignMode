package singleton

//饿汉模式，在包加载的时候初始化，而不是在调用的时候初始化，不支持延迟加载
//实例占用资源多，初始化时间长
type example struct {
	Name string
}

var instance *example

func init()  {
	instance=new(example)
	instance.Name="prk"
}

func GetInstance() *example {
	return instance
}

//懒汉模式，支持延迟加载，但是会降低并发度
type example2 struct {
	Name string
}

var instance2 *example2

func GetInstance2()*example2{
	if instance2==nil{
		instance2=&example2{Name: "why"}
	}
	return instance2
}