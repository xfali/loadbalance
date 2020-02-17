// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package loadbalance

import (
	"math/rand"
	"time"
)

type RandomLoadBalance struct {
	BaseLoadBalance
	rand *rand.Rand
}

func NewRandomLoadBalance() *RandomLoadBalance {
	return &RandomLoadBalance{
		rand:            rand.New(rand.NewSource(time.Now().UnixNano())),
		BaseLoadBalance: BaseLoadBalance{Compare: DefaultCompare, lock: &DummyLocker{}},
	}
}

func (lb *RandomLoadBalance) Select() interface{} {
	lb.lock.RLock()
	defer lb.lock.RUnlock()

	size := len(lb.invokers)
	if size == 0 {
		return nil
	}

	return lb.invokers[lb.rand.Intn(size)]
}

type RandomWeightLoadBalance struct {
	invokers map[interface{}]uint
	total    uint
	lock     RWLocker
	rand     *rand.Rand
}

func NewRandomWeightLoadBalance() *RandomWeightLoadBalance {
	return &RandomWeightLoadBalance{
		invokers: map[interface{}]uint{},
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
		lock:     &DummyLocker{},
	}
}

func (lb *RandomWeightLoadBalance) WithLocker(locker RWLocker) {
	lb.lock = locker
}

func (lb *RandomWeightLoadBalance) Add(weight uint, invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	lb.invokers[invoker] = weight
	lb.total += weight
}

func (lb *RandomWeightLoadBalance) Remove(invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	if v, ok := lb.invokers[invoker]; ok {
		delete(lb.invokers, invoker)
		lb.total -= v
	}
}

func (lb *RandomWeightLoadBalance) Select() interface{} {
	lb.lock.RLock()
	defer lb.lock.RUnlock()

	if len(lb.invokers) == 0 {
		return nil
	}
	var cur uint = 0
	r := uint(lb.rand.Uint64() % uint64(lb.total))
	for k, v := range lb.invokers {
		if cur <= r && r < cur+v {
			return k
		}
		cur += v
	}
	return nil
}