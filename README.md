# loadbalance

## 介绍

loadbalance实现了常用的负载均衡算法：


 标签 | 说明
:---: | :---
轮询 | 从1开始，直到N，然后重新开始循环
加权轮询 | 根据相应权值数的轮询循环选择（平滑轮询）
随机 | 从1-N随机选择一个
加权随机 | 从相应权值数中随机选择一个

## 待完成项

* 继续增加负载均衡算法

## 使用


### 1、直接创建
```cassandraql
lb := loadbalance.NewRoundRobbinLoadBalance()
		lb.Add(1, "a")
		lb.Add(1, "b")
		lb.Add(1, "c")
		for i := 0; i < 10; i++ {
			t.Log(lb.Select(nil))
		}
```

### 1、使用factory
```cassandraql
lb := loadbalance.Create(loadbalance.LBRoundRobbin)
		lb.Add(2, "a")
		lb.Add(5, "b")
		lb.Add(10, "c")
		for i := 0; i < 17; i++ {
			t.Log(lb.Select(nil))
		}
```
