// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package loadbalance

import "context"

type RoundRobbinLoadBalance struct {
	BaseLoadBalance
	i int
}

func NewRoundRobbinLoadBalance() *RoundRobbinLoadBalance {
	return &RoundRobbinLoadBalance{
		i:               0,
		BaseLoadBalance: BaseLoadBalance{Compare: DefaultCompare, lock: &DummyLocker{}},
	}
}

func (lb *RoundRobbinLoadBalance) Select(ctx context.Context) interface{} {
	lb.lock.RLock()
	defer lb.lock.RUnlock()

	size := len(lb.invokers)
	if size == 0 {
		return nil
	}

	fac := lb.invokers[lb.i]
	lb.i = (lb.i + 1) % size
	return fac
}

type weightInvoker struct {
	weight  int
	curW    int64
	invoker interface{}
}

type RoundRobbinWeightLoadBalance struct {
	invokers []weightInvoker
	lock     RWLocker
	Compare  Compare
}

func NewRoundRobbinWeightLoadBalance() *RoundRobbinWeightLoadBalance {
	return &RoundRobbinWeightLoadBalance{
		lock:    &DummyLocker{},
		Compare: DefaultCompare,
	}
}

func (lb *RoundRobbinWeightLoadBalance) WithLocker(locker RWLocker) {
	lb.lock = locker
}

func (lb *RoundRobbinWeightLoadBalance) Add(weight interface{}, invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	lb.invokers = append(lb.invokers, weightInvoker{
		weight:  weight.(int),
		curW:    0,
		invoker: invoker,
	})
}

func (lb *RoundRobbinWeightLoadBalance) Remove(invoker interface{}) {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	index := -1
	for i := range lb.invokers {
		if lb.Compare(lb.invokers[i].invoker, invoker) == 0 {
			index = i
			break
		}
	}
	if index != -1 {
		lb.invokers = append(lb.invokers[:index], lb.invokers[index+1:]...)
	}
}

func (lb *RoundRobbinWeightLoadBalance) Select(ctx context.Context) interface{} {
	lb.lock.RLock()
	defer lb.lock.RUnlock()

	size := len(lb.invokers)
	if size == 0 {
		return nil
	}

	i, index := 0, -1
	var total int64 = 0
	for ; i < size; i++ {
		lb.invokers[i].curW += int64(lb.invokers[i].weight)
		total += int64(lb.invokers[i].weight)
		if index == -1 || lb.invokers[index].curW < lb.invokers[i].curW {
			index = i
		}
	}

	lb.invokers[index].curW -= total
	return lb.invokers[index].invoker
}
