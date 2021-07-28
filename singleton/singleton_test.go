package singleton

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestGetInstance1(t *testing.T) {
	start:=time.Now()
	for i:=1;i<100;i++{
		go func(i int) {
			e:=GetInstance1()
			e.Name=string(i)
			fmt.Printf("%s, %s\n",
				e.Name, time.Since(start))
		}(i)
	}
	e:=GetInstance1()
	fmt.Printf("last one:%s, %s\n",
		e.Name, time.Since(start))
}


func TestGetInstance2(t *testing.T) {
	start:=time.Now()
	for i:=1;i<100;i++{
		go func(i int) {
			e:=GetInstance2()

			mu.Lock()
			fmt.Printf("%s, %s\n",
				e.Name, time.Since(start))
			e.Name=strconv.Itoa(i)
			mu.Unlock()
		}(i)
	}
}