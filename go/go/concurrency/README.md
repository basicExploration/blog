## 探究go语言的并发和并行原理。
> 原文地址: https://github.com/googege/blog/tree/master/go/go/concurrency/README.md

#### goroutine基于线程池的P:M:G协程模型
#### channel 基于生产者消费者模型的无锁队列
#### net.conn 基于epoll的异步io同步阻塞模型
#### syscall 基于操作系统的原生syscall能力
#### gosched 基于阻塞的协程调度
#### go gc基于三色标记法的并发gc模型
#### net/http基于goroutine的http服务器
#### 并发安全的hash map slice
#### 可选性能优化手段unsafe非并发安全的指针调用
#### 可实现cas context基于channel的goroutine流程控制能力
#### 非并发安全的指针