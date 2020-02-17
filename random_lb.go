// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package loadbalance

import (
	"context"
	"math/rand"
	"time"
)

type RandomLoadBalance struct {
	BaseLoadBalance
	rand *rand.Rand
}

func NewRandomLoadBalance() *RandomLoadBalance {
	return &RandomLoadBalance{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
		BaseLoadBalance: BaseLoadBalance{
			Compare:         DefaultCompare,
			LockLoadBalance: LockLoadBalance{&DummyLocker{}},
		},
	}
}

func (lb *RandomLoadBalance) Select(ctx context.Context) interface{} {
	lb.lock.RLock()
	defer lb.lock.RUnlock()

	size := len(lb.invokers)
	if size == 0 {
		return nil
	}

	return lb.invokers[lb.rand.Intn(size)]
}

type RandomWeightLoadBalance struct {
	invokers map[interface{}]int
	total    int
	rand     *rand.Rand

	LockLoadBalance
}

func NewRandomWeightLoadBalance() *RandomWeightLoadBalance {
	return &RandomWeightLoadBalance{
		invokers:        map[interface{}]int{},
		rand:            rand.New(rand.NewSource(time.Now().UnixNano())),
		LockLoadBalance: LockLoadBalance{&DummyLocker{}},
	}
}

func (lb *RandomWeightLoadBalance) Add(weight interface{}, invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	lb.invokers[invoker] = weight.(int)
	lb.total += weight.(int)
}

func (lb *RandomWeightLoadBalance) Remove(invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	if v, ok := lb.invokers[invoker]; ok {
		delete(lb.invokers, invoker)
		lb.total -= v
	}
}

func (lb *RandomWeightLoadBalance) Select(ctx context.Context) interface{} {
	lb.lock.RLock()
	defer lb.lock.RUnlock()

	if len(lb.invokers) == 0 {
		return nil
	}
	var cur int = 0
	r := int(lb.rand.Uint64() % uint64(lb.total))
	for k, v := range lb.invokers {
		if cur <= r && r < cur+v {
			return k
		}
		cur += v
	}
	return nil
}
