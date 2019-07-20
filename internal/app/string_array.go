package app

import (
	"strings"
)

/*提供字符串切片的封装，提供复制切片、添加字符串等方法*/

type StringArray []string

//复制一份字符串切片
func (a *StringArray) Get() interface{} { return []string(*a) }

//添加一个字符串到切片中
func (a *StringArray) Set(s string) error {
	*a = append(*a, s)
	return nil
}

//将切片元素连接起来，以，分隔
func (a *StringArray) String() string {
	return strings.Join(*a, ",")
}
