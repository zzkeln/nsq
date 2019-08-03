package statsd

import (
	"strings"
)

//将ip:port中的.和:都换成_

//h主要是127.0.0.1:8888-->127_0_0_1_8888
func HostKey(h string) string {
	return strings.Replace(strings.Replace(h, ".", "_", -1), ":", "_", -1)
}
