package factory

import "errors"

type ParserFactory struct {
}

func (*ParserFactory) CreateParser(s string) (*parser, error) {
	var b parser
	switch s{
	case "bool":
		b = boolParser{}
		return &b,nil
	case "int":
		b = intParser{}
		return &b,nil
	case "float64":
		b = floatParser{}
		return &b,nil
	}
	return nil,errors.New("connot create parser with put in")
}
