package writers

import (
	"bufio"
	"io"
)
//bufio.Writer对writer加了一层缓存

//这个结构体表示有限缓存的writer
type BoundaryBufferedWriter struct {
	bw *bufio.Writer
}

//创建一个新的有限缓存writer，缓存大小是size，writer是w
func NewBoundaryBufferedWriter(w io.Writer, size int) *BoundaryBufferedWriter {
	return &BoundaryBufferedWriter{
		bw: bufio.NewWriterSize(w, size),
	}
}

func (b *BoundaryBufferedWriter) Write(p []byte) (int, error) {
	//Available()表示当前writer的缓存还剩多少，此时缓存不足了需要flush将缓存内容写到writer
	if len(p) > b.bw.Available() {
		err := b.bw.Flush()
		if err != nil {
			return 0, err
		}
	}
	//会先写到缓存中，如果缓存大小不够p的大小那么会直接写入到writer中
	return b.bw.Write(p)
}

//将缓存的全部内容写到writer中
func (b *BoundaryBufferedWriter) Flush() error {
	return b.bw.Flush()
}
