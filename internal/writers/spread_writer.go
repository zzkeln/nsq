package writers

import (
	"io"
	"time"
)

//封装了一个writer
type SpreadWriter struct {
	w        io.Writer       //输出流
	interval time.Duration	 //这个输出流需要多少时间来刷新完缓存到目标地
	buf      [][]byte        //内存缓存
	exitCh   chan int        //停止刷新标志
}

//返回一个新的输出流
func NewSpreadWriter(w io.Writer, interval time.Duration, exitCh chan int) *SpreadWriter {
	return &SpreadWriter{
		w:        w,
		interval: interval,
		buf:      make([][]byte, 0),
		exitCh:   exitCh,
	}
}

//写p字节到缓存中
func (s *SpreadWriter) Write(p []byte) (int, error) {
	//将p深拷贝一份，并添加到buf中
	b := make([]byte, len(p))
	copy(b, p)
	s.buf = append(s.buf, b)
	return len(p), nil
}

//将缓存全部刷新到目的地，interval控制刷新的频率
func (s *SpreadWriter) Flush() {
	//每次刷新完需要interval时间，共要刷新len(s.buf)个内存块，那么每个内存块刷新完等待interval/len时间
	sleep := s.interval / time.Duration(len(s.buf))
	//创建一个定时器，这个定时器每到sleep秒后就醒来一次
	ticker := time.NewTicker(sleep)
	//遍历所有内存块
	for _, b := range s.buf {
		s.w.Write(b) //刷新缓存到输出流
		//等待sleep秒，或者停止标志被写入
		select {
		case <-ticker.C:
		case <-s.exitCh: // skip sleeps finish writes
		}
	}
	ticker.Stop() //停止定时器
	s.buf = s.buf[:0] //清空缓存
}
