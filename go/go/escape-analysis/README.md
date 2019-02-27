## go语言逃逸分析

```go

package main

type user struct {
    name  string
    email string
}

func main() {
    u1 := createUserV1()
    u2 := createUserV2()

    println("u1", &u1, "u2", &u2)
}

//go:noinline
func createUserV1() user {
    u := user{
        name:  "Bill",
        email: "bill@ardanlabs.com",
    }

    println("V1", &u)
    return u
}

//go:noinline
func createUserV2() *user {
    u := user{// 发生了逃逸分析
        name:  "Bill",
        email: "bill@ardanlabs.com",
    }

    println("V2", &u)
    return &u
}
```

指针指向了栈下的无效地址空间。当 main 函数调用下一个函数，指向的内存将重新映射并将被重新初始化，这就是逃逸分析将开始保持完整性的地方

编译器将检查到，在 createUserV2 的（函数）栈中构造 user 值是不安全的，因此，替代地，会在堆中构造（相应的）值
