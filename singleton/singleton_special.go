package singleton

import (
	"errors"
	"fmt"
	"math/rand"
)

const ex_count = 3

type example3 struct {
	Name string
}

var exMap=make(map[int]*example3)

func init() {
	exMap[0]=&example3{
		Name: "prk",
	}
	exMap[1]=&example3{
		Name: "why",
	}
	exMap[2]=&example3{
		Name: "dd",
	}
}

func GetInstance3(i int) (*example3,error) {
	if i>=0&&i<ex_count{
		return exMap[i],nil
	}
	return nil,errors.New("index i is out of bound")
}

func GetRandomInstance3() (*example3) {
	r:=rand.Intn(ex_count)
	fmt.Println("random int is",r)
	return exMap[r]
}
