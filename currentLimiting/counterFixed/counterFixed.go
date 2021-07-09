package counterFixed

import (
	"sync"
	"sync/atomic"
	"time"
)

// 计数固定窗口
type engine struct {
	interval time.Duration //时间间隔 例如1秒
	total    uint32        //总数 例如 1秒10个
	useTotal uint32        //已使用总数 例如 1秒10个中使用了2个
}

type counterFixed struct {
	signs map[string]*engine
}

var instance *counterFixed
var once sync.Once

func GetCfInstance() *counterFixed {
	once.Do(func() {
		instance = newCounterFixed()
	})
	return instance
}

func newCounterFixed() *counterFixed {
	cf := &counterFixed{make(map[string]*engine)}
	return cf
}

// Add 添加
func (cf *counterFixed) Add(sign string, interval time.Duration, total uint32) {
	e := &engine{interval: interval, total: total}
	go e.timer()
	cf.signs[sign] = e
}

// Check 验证
func (cf *counterFixed) Check(sign string) bool {
	e := cf.signs[sign]
	useTotal := e.useTotal
	if useTotal == e.total {
		return false
	}
	return atomic.CompareAndSwapUint32(&e.useTotal, useTotal, useTotal+1)
}

//根据时间间隔还原数据
func (e *engine) timer() {
	c := time.Tick(e.interval)
	for {
		<-c
		atomic.StoreUint32(&e.useTotal, 0)
	}
}
