package factory

import "strconv"

type parser interface {
	parse(s string) interface{}
}

type intParser struct {

}
type floatParser struct {

}
type boolParser struct {

}

func (intParser) parse(s string) interface{} {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return -1
	}
	return int(i)
}

func (floatParser) parse(s string) interface{} {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1
	}
	return i
}

func (boolParser) parse(s string) interface{} {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return -1
	}
	return i
}
