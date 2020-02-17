// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"github.com/xfali/loadbalance"
	"testing"
)

func TestRRLB(t *testing.T) {
	t.Run("rr", func(t *testing.T) {
		lb := loadbalance.NewRoundRobbinLoadBalance()
		lb.Add("a")
		lb.Add("b")
		lb.Add("c")
		for i := 0; i < 10; i++ {
			t.Log(lb.Select())
		}
	})

	t.Run("rr remove", func(t *testing.T) {
		lb := loadbalance.NewRoundRobbinLoadBalance()
		lb.Add("a")
		lb.Add("b")
		lb.Add("c")
		lb.Remove("b")
		for i := 0; i < 10; i++ {
			t.Log(lb.Select())
		}
	})
}

func TestRRWeightLB(t *testing.T) {
	t.Run("rr weight", func(t *testing.T) {
		lb := loadbalance.NewRoundRobbinWeightLoadBalance()
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		a, b, c := 0, 0, 0
		for i := 0; i < 17; i++ {
			v := lb.Select()
			t.Log(v)
			if v.(string) == "a" {
				a++
			} else if v.(string) == "b" {
				b++
			} else if v.(string) == "c" {
				c++
			}
		}
		t.Logf("a: %d b: %d c: %d\n", a, b, c)
	})

	t.Run("rr weight remove", func(t *testing.T) {
		lb := loadbalance.NewRoundRobbinWeightLoadBalance()
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		lb.Remove("b")
		a, b, c := 0, 0, 0
		for i := 0; i < 12; i++ {
			v := lb.Select()
			t.Log(v)
			if v.(string) == "a" {
				a++
			} else if v.(string) == "b" {
				b++
			} else if v.(string) == "c" {
				c++
			}
		}
		t.Logf("a: %d b: %d c: %d\n", a, b, c)
	})
}

func TestRandomLB(t *testing.T) {
	t.Run("random", func(t *testing.T) {
		lb := loadbalance.NewRandomLoadBalance()
		lb.Add("a")
		lb.Add("b")
		lb.Add("c")
		a, b, c := 0, 0, 0
		for i := 0; i < 3000; i++ {
			v := lb.Select()
			//t.Log(v)
			if v.(string) == "a" {
				a++
			} else if v.(string) == "b" {
				b++
			} else if v.(string) == "c" {
				c++
			}
		}

		t.Logf("a: %d b: %d c: %d\n", a, b, c)
	})

	t.Run("random remove", func(t *testing.T) {
		lb := loadbalance.NewRandomLoadBalance()
		lb.Add("a")
		lb.Add("b")
		lb.Add("c")
		lb.Remove("b")
		a, b, c := 0, 0, 0
		for i := 0; i < 2000; i++ {
			v := lb.Select()
			//t.Log(v)
			if v.(string) == "a" {
				a++
			} else if v.(string) == "b" {
				b++
			} else if v.(string) == "c" {
				c++
			}
		}

		t.Logf("a: %d b: %d c: %d\n", a, b, c)
	})
}

func TestRandomWeightLB(t *testing.T) {
	t.Run("random weight", func(t *testing.T) {
		lb := loadbalance.NewRandomWeightLoadBalance()
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		a, b, c := 0, 0, 0
		for i := 0; i < 17000; i++ {
			v := lb.Select()
			//t.Log(v)
			if v.(string) == "a" {
				a++
			} else if v.(string) == "b" {
				b++
			} else if v.(string) == "c" {
				c++
			}
		}

		t.Logf("a: %d b: %d c: %d\n", a, b, c)
	})

	t.Run("random weight remove", func(t *testing.T) {
		lb := loadbalance.NewRandomWeightLoadBalance()
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		lb.Remove("b")
		a, b, c := 0, 0, 0
		for i := 0; i < 12000; i++ {
			v := lb.Select()
			//t.Log(v)
			if v.(string) == "a" {
				a++
			} else if v.(string) == "b" {
				b++
			} else if v.(string) == "c" {
				c++
			}
		}

		t.Logf("a: %d b: %d c: %d\n", a, b, c)
	})
}
