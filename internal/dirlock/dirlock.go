// +build !windows

package dirlock

import (
	"fmt"
	"os"
	"syscall"
)

/*提供目录锁的封装，目录被锁住时，另外一个进程尝试对目录加锁会失败*/

//封装目录锁
type DirLock struct {
	dir string    //目录
	f   *os.File  //目录句柄
}

//返回一个封装的目录锁指针
func New(dir string) *DirLock {
	return &DirLock{
		dir: dir,
	}
}
//对目录加锁
func (l *DirLock) Lock() error {
	//打开目录句柄
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)  //加一个排它锁
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}
//对目录解锁
func (l *DirLock) Unlock() error {
	defer l.f.Close()//关闭目录句柄
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
