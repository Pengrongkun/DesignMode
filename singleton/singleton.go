package singleton

import "sync"

//饿汉模式，在包加载的时候初始化，而不是在调用的时候初始化，不支持延迟加载
//实例占用资源多，初始化时间长
type example1 struct {
	Name string
}

var instance *example1

func init() {
	instance = new(example1)
	instance.Name = "prk"
}

func GetInstance1() *example1 {
	return instance
}

//懒汉模式，支持延迟加载，但是会降低并发度
var loadIconsOnce sync.Once
var mu sync.Mutex
var sn ="模"

var instance2 *example2

type example2 struct {
	Name string
}

func loadExample2() {
	instance2 = &example2{Name: sn}
}

func GetInstance2() *example2 {
	sn=""
	loadIconsOnce.Do(loadExample2)
	return instance2
}
