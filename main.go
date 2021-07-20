package main

import (
	"designmode/singleton"
	"fmt"
)

func main()  {
	s:=singleton.GetInstance()
	fmt.Println(s.Name)
	s.Name="why"
	s2:=singleton.GetInstance()
	fmt.Println(s2.Name)
}
