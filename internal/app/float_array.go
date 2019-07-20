package app

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*提供浮点数切片的封装，为这个切片提供一些方法*/

type FloatArray []float64

//赋值一份浮点数数组的拷贝，注意返回参数是interface{}
func (a *FloatArray) Get() interface{} { return []float64(*a) }

//参数是"1.2,2.3,3.4"这样的字符串，解析出字符串并赋值给浮点数切片并排序
func (a *FloatArray) Set(param string) error {
	for _, s := range strings.Split(param, ",") {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatalf("Could not parse: %s", s)
			return nil
		}
		*a = append(*a, v)
	}
	sort.Sort(*a)
	return nil
}
//交换浮点数切片2个元素
func (a FloatArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//比较2个元素大小
func (a FloatArray) Less(i, j int) bool { return a[i] > a[j] }
//返回浮点数切片长度
func (a FloatArray) Len() int           { return len(a) }
//返回string形式，以逗号分隔
func (a *FloatArray) String() string {
	var s []string
	for _, v := range *a {
		s = append(s, fmt.Sprintf("%f", v))
	}
	return strings.Join(s, ",")
}
