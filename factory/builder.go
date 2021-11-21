package factory

import "fmt"

type resourcePooling struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

const (
	default_max_total = 8
	default_max_idle  = 8
	default_min_idle  = 0
)

type Builder interface {
	AddName(string) Builder
	AddMaxTotal(int) Builder
	AddMaxIdle(int) Builder
	AddMinIdle(int) Builder
	Build() *resourcePooling
}

type ResourcePoolingBuilder struct {
	rp *resourcePooling
}

func NewResourcePoolingBuilder() *ResourcePoolingBuilder {
	return &ResourcePoolingBuilder{
		rp: &resourcePooling{
			name:     "",
			maxTotal: default_max_total,
			maxIdle:  default_max_idle,
			minIdle:  default_min_idle,
		},
	}
}

func (rpb *ResourcePoolingBuilder) AddName(s string) Builder {
	rpb.rp.name = s
	return rpb
}

func (rpb *ResourcePoolingBuilder) AddMaxTotal(m int) Builder {
	if m > 0 {
		rpb.rp.maxTotal = m
	} else {
		fmt.Println("设置minIdle失败，不能小于0")
	}
	return rpb
}

func (rpb *ResourcePoolingBuilder) AddMaxIdle(m int) Builder {
	if m > 0 {
		rpb.rp.maxIdle = m
	} else {
		fmt.Println("设置minIdle失败，不能小于0")
	}
	return rpb
}

func (rpb *ResourcePoolingBuilder) AddMinIdle(m int) Builder {
	if m > 0 {
		rpb.rp.minIdle = m
	} else {
		fmt.Println("设置minIdle失败，不能小于0")
	}
	return rpb
}

func (rpb *ResourcePoolingBuilder) Build() *resourcePooling {
	if rpb.rp.name == "" {
		fmt.Println("用户名不能为空")
		return nil
	}
	if rpb.rp.maxIdle < rpb.rp.minIdle || rpb.rp.maxTotal < rpb.rp.minIdle {
		fmt.Println("minIdle设置过大")
		return nil
	}
	if rpb.rp.maxIdle > rpb.rp.maxTotal {
		fmt.Println("maxIdle必须小于maxTotal")
		return nil
	}
	return rpb.rp
}
