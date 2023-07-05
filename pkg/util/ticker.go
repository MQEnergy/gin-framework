package util

import (
	"time"
)

type Fn func() error

type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
	Done   chan bool // 添加一个Done通道
}

// NewTicker 实例化定时器
func NewTicker(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
		Done:   make(chan bool), // 初始化Done通道
	}
}

// Start 启动定时器需要执行的任务
func (t *MyTicker) Start() {
	go func() {
		for {
			select {
			case <-t.MyTick.C:
				t.Runner()
			case <-t.Done: // 添加Done通道的处理逻辑
				t.MyTick.Stop()
			}
		}
	}()
}

// Stop 停止定时器
func (t *MyTicker) Stop() {
	t.Done <- true
}
