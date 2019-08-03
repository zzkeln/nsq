package version

import (
	"fmt"
	"runtime"
)

const Binary = "1.1.1-alpha"

//返回nsq当前版本号
func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, Binary, runtime.Version())
}
