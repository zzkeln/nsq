package statsd

import (
	"fmt"
	"io"
)

//这个接口就是搞一套输出的模版，让value按照指定格式写入到输出流中

//Client包括一个输出流和前缀
type Client struct {
	w      io.Writer
	prefix string
}
//创建一个Client的指针
func NewClient(w io.Writer, prefix string) *Client {
	return &Client{
		w:      w,
		prefix: prefix,
	}
}
//假设prefix是test

//假设stat是”STAT“，count=11，那么这个函数将"testSTAT:11"写入到输出流中
func (c *Client) Incr(stat string, count int64) error {
	return c.send(stat, "%d|c", count)
}

func (c *Client) Decr(stat string, count int64) error {
	return c.send(stat, "%d|c", -count)
}

func (c *Client) Timing(stat string, delta int64) error {
	return c.send(stat, "%d|ms", delta)
}

func (c *Client) Gauge(stat string, value int64) error {
	return c.send(stat, "%d|g", value)
}

func (c *Client) send(stat string, format string, value int64) error {
	//format就是前缀+stat参数+format，形成一个输出的模版
	format = fmt.Sprintf("%s%s:%s\n", c.prefix, stat, format)
	//按照模版来将参数value输出到输出流中
	_, err := fmt.Fprintf(c.w, format, value)
	return err
}
