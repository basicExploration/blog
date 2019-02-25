## 探究go语言的并发和并行原理。
> 原文地址: https://github.com/googege/blog/tree/master/go/go/concurrency/README.md

#### goroutine基于线程池的P:M:G协程模型
首先说明一下go可以有两种并发方式
- csp并发

    也就是最常使用的go并发模式
- 共享内存

    通常意义上可以理解为通过共享了内存进而来通信的并发方式，例如加lock

首先我们谈谈关于cpu和操作系统线程，进程的那些事：
我们通常都听过这个一个词，cpu 4核8线程，这里的意思就是cpu实际内核是4核，但是在操作系统看来是8核
cpu，这里的8线程就是指的是8个虚拟内核。然后

#### channel 基于生产者消费者模型的无锁队列
首先解释一下什么是生产者消费者模式：
有三个东西，生产数据的一个任务（可以是线程，进程函数等）一个缓存区域，一个使用数据的任务
这个模式基本上可以类比：厨师做饭+把饭放到前台+你去端饭。

#### net.conn 基于epoll的异步io同步阻塞模型
#### syscall 基于操作系统的原生syscall能力
#### gosched 基于阻塞的协程调度
#### go gc基于三色标记法的并发gc模型
#### net/http基于goroutine的http服务器
#### 并发安全的hash map slice
#### 可选性能优化手段unsafe非并发安全的指针调用
#### 可实现cas context基于channel的goroutine流程控制能力
#### 非并发安全的指针
#### 以实现有限的动态性atomic基于cpu原子操作的包装，