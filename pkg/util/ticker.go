package util

import (
	"time"
)

type Fn func() error

type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
}

// NewTicker 实例化定时器
func NewTicker(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

// Start 启动定时器需要执行的任务
func (t *MyTicker) Start() {
	for {
		select {
		case <-t.MyTick.C:
			t.Runner()
		}
	}
}
