// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"github.com/xfali/loadbalance"
	"testing"
)

func TestFactory(t *testing.T) {
	t.Run("rr", func(t *testing.T) {
		lb := loadbalance.Create(loadbalance.LBRoundRobbin)
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		for i := 0; i < 17; i++ {
			t.Log(lb.Select(nil))
		}
	})

	t.Run("rrw", func(t *testing.T) {
		lb := loadbalance.Create(loadbalance.LBRoundRobbinWeight)
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		for i := 0; i < 17; i++ {
			t.Log(lb.Select(nil))
		}
	})

	t.Run("random", func(t *testing.T) {
		lb := loadbalance.Create(loadbalance.LBRandom)
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		for i := 0; i < 17; i++ {
			t.Log(lb.Select(nil))
		}
	})

	t.Run("random weight", func(t *testing.T) {
		lb := loadbalance.Create(loadbalance.LBRandomWeight)
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		for i := 0; i < 17; i++ {
			t.Log(lb.Select(nil))
		}
	})
}
