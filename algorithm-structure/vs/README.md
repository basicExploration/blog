## go ints 使用的算法和各大经典排序算法的对比

> n 表示 使用go的rand 的 perm随机生成的数据的len

> 前 go 后 其它算法

- go ints vs 冒泡法

```bash
n = 1 5 10 15 20 25 30 35

56.2 ns/op
5.13 ns/op

84.8 ns/op
25.1 ns/op

138 ns/op
76.7 ns/op

248 ns/op
151 ns/op

352 ns/op
257 ns/op

395 ns/op
402 ns/op

599 ns/op
580 ns/op

691 ns/op
810 ns/op

```

基本上 30是分界线，越往后 go越快

- go ints vs 选择排序

```bash
n = 1 5 10 20 21 100 1000 
56.3 ns/op
5.13 ns/op

85.4 ns/op
26.3 ns/op

127 ns/op
74.6 ns/op

323 ns/op
290 ns/op

338 ns/op
329 ns/op

2936 ns/op
7364 ns/op

51202 ns/op
673224 ns/op
```

可以看到 在go vs 选择排序的对比中 21是一个分界线。
