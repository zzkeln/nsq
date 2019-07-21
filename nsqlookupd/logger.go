package nsqlookupd

import (
	"github.com/nsqio/nsq/internal/lg"
)

/*实现了nsqlookupd包里的logf接口，使用options的日志等级作为配置日志等级*/

//声明了Logger的类型
type Logger lg.Logger

//定义了5种log类型
const (
	LOG_DEBUG = lg.DEBUG
	LOG_INFO  = lg.INFO
	LOG_WARN  = lg.WARN
	LOG_ERROR = lg.ERROR
	LOG_FATAL = lg.FATAL
)

//实现logf接口
func (n *NSQLookupd) logf(level lg.LogLevel, f string, args ...interface{}) {
	//配置日志等级使用options里的
	lg.Logf(n.opts.Logger, n.opts.LogLevel, level, f, args...)
}
