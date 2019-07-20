package stringy

/*提供一些字符串切片的方法，例如2个字符串切片的并集，对字符串切片去重，添加一个字符串到切片中*/

//将字符串a添加到字符串切片s中（只有不存在的情况下），返回新的切片
func Add(s []string, a string) []string {
	for _, existing := range s {
		if a == existing {
			return s
		}
	}
	return append(s, a)

}

//将a的元素并入s中，返回s。这里会改变参数s，而且遍历s未必高效，也可以试试map来check-exist
func Union(s []string, a []string) []string {
	for _, entry := range a {
		found := false
		for _, existing := range s {
			if entry == existing {
				found = true
				break
			}
		}
		if !found {
			s = append(s, entry)
		}
	}
	return s
}

//对s进行去重，返回去重后的结果，这里并不改变参数s
func Uniq(s []string) (r []string) {
	for _, entry := range s {
		found := false
		for _, existing := range r {
			if existing == entry {
				found = true
				break
			}
		}
		if !found {
			r = append(r, entry)
		}
	}
	return
}
