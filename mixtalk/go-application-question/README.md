### GO语言面试知识考点总汇
> 答案参考1：https://github.com/googege/blog/tree/master/go/go/concurrency/README.md

> 答案参考2 https://github.com/googege/blog/tree/master/go/tool/pprof/README.md
- goroutine 基于线程池的P:M:G协程模型
- channel 基于生产者消费者模型的无锁队列
- net.conn 基于epoll的异步io同步阻塞模型
- syscall 基于操作系统的原生syscall能力
- gosched 基于阻塞的协程调度
- gc 基于三色标记法的并发gc模型
- io.reader/writer unix文件哲学升级版
- net/http 基于goroutine的http服务器
- 开箱即用error基于c风格的if(erron ! = 0)错误处理机制
- panic 传统的exception异常机制可配合coredumprecover可用于恢复异常的堆栈，
- 以进行排错map传统的hashmap
- 并发安全的hash map slice
- 基于内存复用和读优化设计的数据结构
- defer函数返回前清理各种垃圾，防止内存泄露
- go tool asm go专用汇编，
- 可选性能优化手段cgo非并发安全的c调用能力，
- 可选性能优化手段unsafe非并发安全的指针调用，
- 可选性能优化手段reflect提供反射能力，
- 以实现有限的动态性atomic基于cpu原子操作的包装，
- 可实现cas context基于channel的goroutine流程控制能力
- interface提供高级语言的抽象和多态能力闭包提供主流编程语言的闭包设计
- 逃逸分析提供主流编程语言的逃逸优化能力指针提供并发安全的指针
- 非并发安全的指针
- pprof自带的性能分析工具用于调优和查错

***

后端工程师 面试重点算法和数据结构重点

算法与数据结构是面试考察的重中之重，也是大家日后学习时需要着重训练的部分。简单的总结一下，大约有这些内容：


- 算法 - Algorithms

排序算法：快速排序、归并排序、计数排序

搜索算法：回溯、递归、剪枝技巧

图论：最短路、最小生成树、网络流建模

动态规划：背包问题、最长子序列、计数问题

基础技巧：分治、倍增、二分、贪心

- 数据结构 - Data Structures

数组与链表：单 / 双向链表、跳舞链

栈与队列

树与图：最近公共祖先、并查集

哈希表

堆：大 / 小根堆、可并堆

字符串：字典树、后缀树

