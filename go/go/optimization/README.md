### go语言中的常见优化

- 如果你有一个数量很大的struct请记住要传递指针而不是值
- go中slice和map都可以自动的进行扩容，但是要记得扩容就要浪费时间，甚至slice还是要重新指向一个新的底层array，也就是重新分配内存空间，这都要消耗大量的时间。所以能判断多少空间的，自己给定一个len和cap就好了,map中只能给定一个len不能设定cap因为map可以自动扩容。而且它也不能使用cap方法来测定他的cap，只能用len()方法
- 为什么不不让你乱用指针？原因很简单，如果发生了[逃逸机制](https://github.com/googege/blog/tree/master/go/go/escape-analysis/README.md)以后，那么你就面临了一个问题,大量的数据被存储在了堆里，而系统是无法管控你的堆空间的，你需要使用gc进行控制，而gc势必要发生时间上的消耗，和"stop the world"，所以不能所有数据统统使用指针，得不偿失！























