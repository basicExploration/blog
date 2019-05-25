# 复习大纲
 > http://naotu.baidu.com/file/be0f61015fccab0c0b29f3e042174cb1?token=65092c0c03d17a8a
<!--Note-->
## 资料

- [CS-Notes](https://cyc2018.github.io/CS-Notes/#/)
- [后端面试进阶指南](https://xiaozhuanlan.com/topic/2167809435)

## 协议

分享或者修改演绎时请保留本协议，并署名  [@CyC2018](https://dwz.cn/ZGWCOICD)。

[CC BY - Creative Commons Attribution](http://creativecommons.org/licenses/by-nc-sa/4.0)

![](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)
<!--/Note-->

## README

<!--Note-->
### 协议

分享或者修改演绎时请保留本协议，并署名  [@CyC2018](https://dwz.cn/ZGWCOICD)。

[CC BY - Creative Commons Attribution](http://creativecommons.org/licenses/by-nc-sa/4.0)

![](https://i.creativecommons.org/l/by-nc-sa/4.0/88x31.png)

### 目的

方便大家系统梳理知识点，并且针对每个知识点可以在本脑图中写 Markdown 笔记。

每个知识点也有相应的完成度和优先级，对于不同重要程度的知识点应该采取不同的复习方法，从而提高学习效率。

你应该把这个脑图当做最基本的复习材料，每天都要大概地过一遍，保持短期记忆，一定要知道，短期记忆对面试来说至关重要。

也可以将收集的资料整理在本脑图中，从而方便复习。

### 来源

[知识总结方法](https://xiaozhuanlan.com/topic/4150387926)

### 关于我

https://dwz.cn/ZGWCOICD

### 样式修改

百度脑图自带的样式效果不佳，建议安装以下样式脚本：[百度脑图](https://userstyles.org/styles/163774/theme)。

### 保存方法

点击左上角菜单，然后另存为“我的文档”。




<!--/Note-->

## 数据结构与算法

<!--Note-->
### 资料

- [剑指 Offer 题解](https://cyc2018.github.io/CS-Notes/#/notes/剑指%20offer%20题解)
- [Leetcode 题解](https://cyc2018.github.io/CS-Notes/#/notes/Leetcode%20题解)
- [算法](https://cyc2018.github.io/CS-Notes/#/notes/算法)
- 《算法》
- 《剑指 Offer》
- 《程序员代码面试指南》
- 《挑战程序设计竞赛》
- [Leetcode](https://leetcode.com/problemset/algorithms/)
- [玩转算法面试 从真题到思维全面提升算法思维](https://coding.imooc.com/class/82.html)
<!--/Note-->

### 算法思想

#### 排序

##### 选择排序

##### 冒泡排序

##### 插入排序

##### 希尔排序

##### 归并排序

##### 堆排序

#### 字符串

##### 指纹

##### KMP

##### AC 自动机

##### 排序

##### Trie

#### 树

##### 红黑树

<!--Note-->
###### 回答

- JDK 中 TreeMap 和 TreeSet，1.8 之后的 HashMap 和 ConcurrentHashMap
- 介绍二叉查找树、23查找树，再介绍红黑树原理
- 与 B+ 树进行比较

###### 资料

- [红黑树 - 维基百科](https://zh.wikipedia.org/zh-hans/%E7%BA%A2%E9%BB%91%E6%A0%91)
<!--/Note-->

##### B+ 树

##### LSM

##### AVL

#### 图

##### 最短路径

##### 最小生成树

##### 拓扑排序

##### 并查集

##### 网络流

#### 散列表

##### 拉链法

##### 线性探测法

#### 其它

##### 汉诺塔

##### 哈夫曼编码

### 海量数据处理

#### TOP-K

#### 海量数据判重

#### 海量数据排序

#### MapReduce

### 数学与逻辑

#### 概率题

##### 抢红包

##### 洗牌

##### 蓄水池抽样

##### Rand7

#### 智力题

## 操作系统

<!--Note-->
### 资料

- 《现代操作系统》
- 《深入理解计算机系统》
- 《UNIX 环境高级编程》
- 《Unix/Linux 编程实践教程》
- 《鸟哥的 Linux 私房菜》
- 《The Linux Command Line》
<!--/Note-->

### 基础

#### 进程与线程

#### 进程状态

#### 进程调度算法

#### 线程实现方式

#### 协程

#### 进程同步问题

#### 进程通信

#### 死锁

- 死锁必要条件、解决死锁策略，能写出和分析死锁的代码，能说明在数据库管理系统或者 Java 中如何解决死锁。

#### 虚拟内存

#### 页面置换算法

特别是 LRU 的实现原理，最好能手写，再说明它在 Redis 等作为缓存置换算法。

#### 分页与分段

#### 静态链接与动态链接

### Linux

#### 文件系统

- 从文件系统的角度分析数据恢复原理

#### 硬链接与软链接

#### 常用命令

- 能够使用常用的命令，比如 cat 文件内容查看、find 搜索文件，以及 cut、sort 等管线命令。了解 grep 和 awk 的作用。

#### 僵尸进程与孤儿进程

- 僵尸进程与孤儿进程的区别，从 SIGCHLD 分析产生僵尸进程的原因。

## 网络

<!--Note-->
### 资料

- 《计算机网络 自顶向下方法》
- 《计算机网络》
- 《TCP/IP 详解 卷 1：协议》
- 《UNIX 网络编程 卷 1：套接字联网 API》
- 《Linux 多线程服务端编程》
- 《图解 HTTP》
<!--/Note-->

### 基础

- [计算机网络](https://cyc2018.github.io/CS-Notes/#/notes/计算机网络)

#### 体系结构

#### 以太网

#### 网络硬件设备

- 集线器、交换机、路由器的作用，以及所属的网络层。

#### IP 数据报

#### ARP 协议

#### ICMP 协议

#### UDP 与 TCP

#### TCP 连接

- 理解三次握手以及四次挥手具体过程，三次握手的原因、四次挥手原因、TIME_WAIT 的作用。

#### TCP 可靠传输

- 设计可靠 UDP 协议

#### TCP 拥塞控制

#### DNS

-

### HTTP

- [HTTP](https://cyc2018.github.io/CS-Notes/#/notes/HTTP)

#### GET 与 POST

#### 状态码

#### Cookie

#### 缓存

- [Expires 和 max-age 的区别](https://www.cnblogs.com/yinhaiming/articles/1490811.html)
- [Expires vs max-age, which one takes priority if both are declared in a HTTP response?
](https://stackoverflow.com/questions/7549177/expires-vs-max-age-which-one-takes-priority-if-both-are-declared-in-a-http-resp)

#### 连接管理

#### HTTPs

#### HTTP/2

#### 版本比较

#### HTTP 与 FTP

### Socket

- [Socket](https://cyc2018.github.io/CS-Notes/#/notes/Socket)

#### I/O 模型

#### 多路复用

#### Java NIO

## 数据库

<!--Note-->
### 资料

- 《MySQL 必知必会》
- [Leetcode](https://leetcode.com/problemset/database/)
- 《高性能 MySQL》
- 《MySQL 技术内幕》
- 《Redis 设计与实现》
- 《Redis 实战》
- 《大规模分布式存储系统》
<!--/Note-->

### SQL

#### 手写分组查询

#### 手写连接查询

#### 连接与子查询

#### drop、delete、truncate

#### 视图

- 视图的作用，以及何时能更新视图。

#### 存储过程

#### 触发器

### 系统原理

#### ACID

#### 隔离级别

四大隔离级别，以及不可重复读和幻影读的出现原因。

#### 封锁

封锁的类型以及粒度，两段锁协议，隐式和显示锁定。

#### 乐观锁与悲观锁

#### MVCC 

#### 范式

#### SQL 与 NoSQL

### MySQL

#### B+ Tree

#### 索引以及优化

#### 查询优化

#### InnoDB 与 MyISAM

#### 水平切分与垂直切分

#### 主从复制

#### 日志

### Redis

#### 字典和跳跃表

#### 使用场景

#### 与 Memchached 的比较

#### RDB 和 AOF 持久化机制

#### 数据淘汰机制

#### 事件驱动模型

#### 主从复制

#### 集群与分布式

#### 事务

#### 线程安全问题

## 面向对象

<!--Note-->
### 资料

- 《Head First 设计模式》
<!--/Note-->

### 思想

#### 三大特性

#### 设计原则

### 设计模式

#### 单例模式

手写单例模式，特别是双重检验锁以及静态内部类。

#### 工厂模式

手写工厂模式。


#### MVC

理解 MVC，结合 SpringMVC 回答。

#### 代理模式

结合 Spring 中的 AOP 回答。

#### JDK 中常用的设计模式

例如装饰者模式、适配器模式、迭代器模式等。

## 系统设计

<!--Note-->
#### 资料

- 《大型网站技术架构》
- 《从 Paxos 到 Zookeeper》
- 《微服务设计》
<!--/Note-->

### 基础

- [系统设计基础](https://cyc2018.github.io/CS-Notes/#/notes/系统设计基础)

#### 性能

#### 伸缩性

#### 扩展性

#### 可用性

#### 安全性

### 分布式

- [分布式](https://cyc2018.github.io/CS-Notes/#/notes/分布式)

#### 分布式事务

#### CAP

#### BASE

#### Paxos

#### Raft

#### 分布式锁

#### 分布式 ID

### 集群

- [集群](https://cyc2018.github.io/CS-Notes/#/notes/集群)

#### 负载均衡

#### Session 管理

### 缓存

- [缓存](https://cyc2018.github.io/CS-Notes/#/notes/缓存)

#### 缓存特征

#### LRU

#### 缓存位置

#### CDN

#### 缓存问题

#### 一致性哈希

### 攻击技术

- [攻击技术](https://cyc2018.github.io/CS-Notes/#/notes/攻击技术)

#### XSS

#### CSRF

#### SQL 注入

#### DDoS

### 消息队列

- [消息队列](https://cyc2018.github.io/CS-Notes/#/notes/消息队列)

#### 消息模型

#### 使用场景

#### 可靠性

### 高并发系统

#### 秒杀系统

#### 限流算法

#### 服务熔断与服务降级

### 服务拆分

#### 幂等性

#### 远程服务访问方法

#### 微服务

#### SOA

### 系统设计

#### Web 页面请求过程

#### 二维码登录

#### TinyURL

#### KV 存储系统

#### 搜索引擎

## Java

<!--Note-->
### 资料

- 《JAVA 核心技术》
- 《Java 编程思想》
- 《Effective java 中文版》
- 《深入理解 Java 虚拟机》
- 《Java 并发编程实战》
- 《精通 Spring 4.x》
- 《Spring 揭秘》
<!--/Note-->

### 基础

- [Java 基础](https://cyc2018.github.io/CS-Notes/#/notes/Java%20基础)

### 虚拟机

- [Java 虚拟机](https://cyc2018.github.io/CS-Notes/#/notes/Java%20虚拟机)

### 并发

- [Java 并发](https://cyc2018.github.io/CS-Notes/#/notes/Java%20并发)

### 容器

- [Java 容器](https://cyc2018.github.io/CS-Notes/#/notes/Java%20容器)

### I/O

- [Java IO](https://cyc2018.github.io/CS-Notes/#/notes/Java%20IO)

### Web

- [69 道 Spring 面试题和答案](http://ifeve.com/spring-interview-questions-and-answers/)
- [Spring 面试题](https://github.com/Homiss/Java-interview-questions/blob/master/%E6%A1%86%E6%9E%B6/Spring%20%E9%9D%A2%E8%AF%95%E9%A2%98.md)
- [Spring 面试问答 Top 25](http://www.importnew.com/15851.html)
- [Spring 总结以及在面试中的一些问题.](https://www.cnblogs.com/wang-meng/p/5701982.html)


## C++

## 中间件

### RabbitMQ

### ZooKeeper

### Dubbo

### Nginx

## 其它知识

### 新技术

### 开源项目

## 项目

## 面试相关

### 简历

### 投递
