## go 性能测试

有以下几个工具需要知道
- go tool pprof 也就是pprof性能测试工具
   - alloc_space 选项替代默认的 -inuse_space 选项。这将会向你展示每一次分配发生在哪里，不管你分析数据时它是不是还在内存中
   (使用 go tool pprof 查看所有的帮助信息参数信息啥的 go tool trace等go tool命令也是一样，输入命令不带任何的tag就可以查看
   所有的命令，但是如果是一级命令 比如 go build 该如何查看help呢？使用go help build 这种形式即可)
   - 进入命令中 使用help查看命令中的全部信息，使用help 具体命令可以显示更详细的信息 例如 help list
- google's pprof工具
- go tool trace 工具
- go test 压力测试 -bench(压测对象 可以使用. 来测试全部的对象也可以单独制定函数) -benchtime (压测时间)  -benchmem(压力测试中的内存占用启动标签，不需要制定=谁，这个标签是开启内存测试的标志) -memprofile（将生成适用于pprof的测试文件）
- go build  -gcflags "-m -m" 这个标签可以测试何时何物逃逸分析了。（就是从栈逃到堆了）

### go tool pprof
首先pprof进行分析是有前提的那就是需要有数据，如何来提供数据？go提供了三种办法
- http/pprof包 将测试单独列为一个独立的goroutine然后单独监听，此包就可以单独的在浏览器中打开一个数据分析的html页面
- runtime/pprof包 将这包的 类似 start 和stop函数放到数据中间，就可以对这段数据进行测试，进而生成某些测试数据
- go test 使用 memprofile标签就可以得到一个册数数据
### google's pprof
使用测试数据，可将此数据更加可视化的显示在html中
### go tool trace
更pprof 类似可使用 http/httptrace包和runtime/trace包进行数据的测试得到测试数据，然后是用go tool trace + 数据即可
### go test
跟上文一致，具体的就不展开了就是那个意思。
### go build
如上

