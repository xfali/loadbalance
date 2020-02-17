// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package loadbalance

const (
	LBUnknown     = iota
	LBRoundRobbin
	LBRoundRobbinWeight
	LBRandom
	LBRandomWeight
)

func Create(t int) LoadBalance {
	switch t {
	case LBRoundRobbin:
		return NewRoundRobbinLoadBalance()
	case LBRoundRobbinWeight:
		return NewRoundRobbinWeightLoadBalance()
	case LBRandom:
		return NewRandomLoadBalance()
	case LBRandomWeight:
		return NewRandomWeightLoadBalance()
	}
	return nil
}