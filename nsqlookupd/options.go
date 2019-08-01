package nsqlookupd

import (
	"log"
	"os"
	"time"

	"github.com/nsqio/nsq/internal/lg"
)

type Options struct {
	LogLevel  lg.LogLevel `flag:"log-level"`           //日志等级
	LogPrefix string      `flag:"log-prefix"`          //日志前缀
	Logger    Logger

	TCPAddress       string `flag:"tcp-address"`       //监听的tcp地址
	HTTPAddress      string `flag:"http-address"`      //监听的http地址
	BroadcastAddress string `flag:"broadcast-address"` //广播地址

	InactiveProducerTimeout time.Duration `flag:"inactive-producer-timeout"`  //生产者超时时间，默认300s
	TombstoneLifetime       time.Duration `flag:"tombstone-lifetime"`
}

//返回nsqlookupd的配置选项
func NewOptions() *Options {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	return &Options{
		LogPrefix:        "[nsqlookupd] ",  //日志前缀
		LogLevel:         lg.INFO,          //日志等级是info
		TCPAddress:       "0.0.0.0:4160",
		HTTPAddress:      "0.0.0.0:4161",
		BroadcastAddress: hostname,

		InactiveProducerTimeout: 300 * time.Second, //300s超时
		TombstoneLifetime:       45 * time.Second,
	}
}
