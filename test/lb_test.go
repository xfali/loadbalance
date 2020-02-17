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
		lb.Add(nil, "a")
		lb.Add(nil, "b")
		lb.Add(nil, "c")
		for i := 0; i < 10; i++ {
			t.Log(lb.Select(nil))
		}
	})

	t.Run("rr remove", func(t *testing.T) {
		lb := loadbalance.NewRoundRobbinLoadBalance()
		lb.Add(nil, "a")
		lb.Add(nil, "b")
		lb.Add(nil, "c")
		lb.Remove("b")
		for i := 0; i < 10; i++ {
			t.Log(lb.Select(nil))
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
			v := lb.Select(nil)
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
		if a != 2 || b != 5 || c != 10 {
			t.Fatal()
		}
	})

	t.Run("rr weight remove", func(t *testing.T) {
		lb := loadbalance.NewRoundRobbinWeightLoadBalance()
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		lb.Remove("b")
		a, b, c := 0, 0, 0
		for i := 0; i < 12; i++ {
			v := lb.Select(nil)
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
		if a != 2 || b != 0 || c != 10 {
			t.Fatal()
		}
	})
}

func TestRandomLB(t *testing.T) {
	t.Run("random", func(t *testing.T) {
		lb := loadbalance.NewRandomLoadBalance()
		lb.Add(nil, "a")
		lb.Add(nil, "b")
		lb.Add(nil, "c")
		a, b, c := 0, 0, 0
		for i := 0; i < 3000; i++ {
			v := lb.Select(nil)
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
		if (900 > a || a > 1100) || (900 > b || b > 1100) || (900 > c || c > 1100) {
			t.Fatal()
		}
	})

	t.Run("random remove", func(t *testing.T) {
		lb := loadbalance.NewRandomLoadBalance()
		lb.Add(nil, "a")
		lb.Add(nil, "b")
		lb.Add(nil, "c")
		lb.Remove("b")
		a, b, c := 0, 0, 0
		for i := 0; i < 2000; i++ {
			v := lb.Select(nil)
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
		if (900 > a || a > 1100) || b != 0 || (900 > c || c > 1100) {
			t.Fatal()
		}
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
			v := lb.Select(nil)
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
		if (1900 > a || a > 2100) || (4900 > b || b > 5100) || (9900 > c || c > 11000) {
			t.Fatal()
		}
	})

	t.Run("random weight remove", func(t *testing.T) {
		lb := loadbalance.NewRandomWeightLoadBalance()
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		lb.Remove("b")
		a, b, c := 0, 0, 0
		for i := 0; i < 12000; i++ {
			v := lb.Select(nil)
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
		if (1900 > a || a > 2100) || b != 0 || (9900 > c || c > 11000) {
			t.Fatal()
		}
	})
}
