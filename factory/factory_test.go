package factory

import "testing"

func TestParserFactory_CreateParser(t *testing.T) {
	a:=ParserFactory{}
	p1,err1:=a.CreateParser("bool")
	if err1!=nil{
		t.Fatal("an error happened")
	}
	p2,err2:=a.CreateParser("int")
	if err2!=nil{
		t.Fatal("an error happened")
	}
	p3,err3:=a.CreateParser("float64")
	if err3!=nil{
		t.Fatal("an error happened")
	}
	r1:=(*p1).parse("false")
	if r1!=false{
		t.Errorf("bool false,expected 123,but found %v",r1)
	}
	r2:=(*p2).parse("123")
	if r2!=123{
		t.Errorf("int false,expected 123,but found %v",r2)
	}
	r3:=(*p3).parse("11.11")
	if r3!=11.11{
		t.Errorf("float false,expected 11.11,but found %v",r3)
	}
	_,err4:=a.CreateParser("haha")
	if err4.Error()!="connot create parser with put in"{
		t.Error("error is not out put like expected,",err4)
	}
}
