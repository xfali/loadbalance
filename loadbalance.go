// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package loadbalance

import (
	"context"
	"reflect"
	"sync"
)

type Compare func(a, b interface{}) int

type RWLocker interface {
	Lock()
	Unlock()

	RLock()
	RUnlock()
}

type LoadBalance interface {
	//增加一个invoker
	//
	//factor影响选择的因素,不同的实现该类型不一致: 如带权值负载均，weight类型为int
	Add(factor interface{}, invoker interface{})
	//从记录中删除，不会再次被选中
	Remove(invoker interface{})
	//选择invoker
	Select(ctx context.Context) interface{}
	//设置锁
	WithLocker(RWLocker)
}

type BaseLoadBalance struct {
	invokers []interface{}
	Compare  Compare
	lock     RWLocker
}

func (lb *BaseLoadBalance) WithLocker(locker RWLocker) {
	lb.lock = locker
}

func (lb *BaseLoadBalance) Add(factor interface{}, invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	lb.invokers = append(lb.invokers, invoker)
}

func (lb *BaseLoadBalance) AddInvokers(invokers ...interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	lb.invokers = append(lb.invokers, invokers...)
}

func (lb *BaseLoadBalance) Remove(invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	index := -1
	for i := range lb.invokers {
		if lb.Compare(lb.invokers[i], invoker) == 0 {
			index = i
			break
		}
	}
	if index != -1 {
		lb.invokers = append(lb.invokers[:index], lb.invokers[index+1:]...)
	}
}

type DummyLocker struct{}

func (l *DummyLocker) Lock() {}

func (l *DummyLocker) Unlock() {}

func (l *DummyLocker) RLock() {}

func (l *DummyLocker) RUnlock() {}

type RwLocker sync.RWMutex

func (l *RwLocker) Lock() {
	(*sync.RWMutex)(l).Lock()
}

func (l *RwLocker) Unlock() {
	(*sync.RWMutex)(l).Unlock()
}

func (l *RwLocker) RLock() {
	(*sync.RWMutex)(l).RLock()
}

func (l *RwLocker) RUnlock() {
	(*sync.RWMutex)(l).RUnlock()
}

func DefaultCompare(a, b interface{}) int {
	if reflect.DeepEqual(a, b) {
		return 0
	} else {
		return 1
	}
}
