// short for "log"
package lg

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*提供来打日志的封装，定义了日志等级，提供Logf接口，用于打日志*/

//定义了5种日志类型，依次从低到高是debug、info、warn、error、fatal。等级最低为1级，0是无效等级
const (
	DEBUG = LogLevel(1)
	INFO  = LogLevel(2)
	WARN  = LogLevel(3)
	ERROR = LogLevel(4)
	FATAL = LogLevel(5)
)

type AppLogFunc func(lvl LogLevel, f string, args ...interface{})

// Logger是一个接口，提供Output方法。实现了这个接口的都是一种logger
type Logger interface {
	Output(maxdepth int, s string) error
}

//空日志，实现Output方法，啥都没做
type NilLogger struct{}
func (l NilLogger) Output(maxdepth int, s string) error {
	return nil
}

//日志等级
type LogLevel int

//日志等级的几个方法
//获得当前等级
func (l *LogLevel) Get() interface{} { return *l }
//设置等级，参数是不区分大小写的字符串，转换成debug、info、warn、error和fatal
func (l *LogLevel) Set(s string) error {
	lvl, err := ParseLogLevel(s)
	if err != nil {
		return err
	}
	*l = lvl
	return nil
}
//返回等级的字符串形式
func (l *LogLevel) String() string {
	switch *l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "invalid"
}
//将登记的字符串形式转换成数字形式
func ParseLogLevel(levelstr string) (LogLevel, error) {
	//将字符串全部转换为小写
	switch strings.ToLower(levelstr) {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	}
	//其它无效情况返回0等级
	return 0, fmt.Errorf("invalid log level '%s' (debug, info, warn, error, fatal)", levelstr)
}

//logger是一个日志，cfgLevel是配置文件种的日志登记，msgLevel是当前消息的日志等级，f是format字符串，args都是参数
func Logf(logger Logger , cfgLevel LogLevel, msgLevel LogLevel, f string, args ...interface{}) {
	//如果配置文件等级大于当前消息等级直接返回。例如配置文件等级是warn，那么只有warn, error和fatal日志才会被打出来
	//低于这一等级的日志都不用输出
	if cfgLevel > msgLevel {
		return
	}
	//将msgLevel和参数拼接成字符串然后输出
	logger.Output(3, fmt.Sprintf(msgLevel.String()+": "+f, args...))
}

//提供默认的LogFatal的实现
func LogFatal(prefix string, f string, args ...interface{}) {
	//打到错误输出的日志
	logger := log.New(os.Stderr, prefix, log.Ldate|log.Ltime|log.Lmicroseconds)
	Logf(logger, FATAL, FATAL, f, args...)  //打印日志
	os.Exit(1) //直接挂掉进程
}
