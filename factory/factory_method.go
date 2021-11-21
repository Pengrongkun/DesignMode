package factory

type ParserFactoryMethod interface {
	CreateParser()*parser
}

type BoolParserFactoryMethod struct {

}

func (BoolParserFactoryMethod)CreateParser()*parser  {
	var b parser
	b=boolParser{}
	return &b
}

type IntParserFactoryMethod struct {

}

func (IntParserFactoryMethod)CreateParser()*parser  {
	var i parser
	i=intParser{}
	return &i
}

type FloatParserFactoryMethod struct {

}

func (FloatParserFactoryMethod)CreateParser()*parser  {
	var f parser
	f=floatParser{}
	return &f
}

//工厂的工厂，不包含成员变量，可以服用
//不需要每次都创建新的工厂类对象
var ParserFactoryMap map[string]ParserFactoryMethod

func init() {
	ParserFactoryMap=make(map[string]ParserFactoryMethod)
	ParserFactoryMap["bool"]=BoolParserFactoryMethod{}
	ParserFactoryMap["int"]=IntParserFactoryMethod{}
	ParserFactoryMap["float"]=FloatParserFactoryMethod{}
}

func getParserFactory(s string)*ParserFactoryMethod{
	if p,ok:=ParserFactoryMap[s];ok{
		return &p
	}
	return nil
}