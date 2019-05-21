## 关于go语言的几个陷阱，以及我们应该注意的东西

1. 闭包

所谓闭包就是指一个函数中的函数，并且这个函数可以调用外部的变量并且无论使用多少次，
都可以一直拥有这个变量不回收，那么这个变量可以称为闭包变量。

```go
func test() []func()  {
	var funs []func()
	fmt.Println(funs == nil)
	for i:=0;i<2 ;i++  {
		//t := i  然后下面改成t就ok了。就是这么简单。因为t就一个值，每个t都是一个新的内存地址而i是一个值，它也是引用变量。
		// 闭包只是抓住某个变量不放手 刚好i是引用变量所以抓住i不放手的话 那么所有函数的i都是一眼的，但是t又不一样，抓住不放手也无所谓了。
		funs = append(funs, func() {
			println(&i,i)
		})
	}
	return funs
}

func main(){
	funs:=test()
	for _,f:=range funs{
		f()
	}
}
```
答案 xxx 2
    xxx 2
    
    原因就是 闭包是引用变量，所以 他们的最后的i是最后的那个i

2. 循环体变量

```go

package main

import (
	"fmt"
	"time"
)
func main(){
	tt()
}
func tt() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1e3)
			fmt.Println(i)
		}()
	}
	time.Sleep(1e9)
}

```
这个函数执行的时候 tt中打印出来的是10
原因也是很简单，因为go在初始化的时候先初始化参数量，全局先初始参数再看函数，在函数内部先初始参数再进行运算，所以 就造成在for执行完后 这里的i是同样的i
以为初始化的参数 i一直会变，但是都是这个变量本身，又因为for循环比内部的函数速度快很多，导致，当for循环进行完了，这些函数还没正式开始运行，然后i就取最终值 10了

3. return 和 defer 的执行顺序

在go里

```go

package main

import "fmt"

func tt() int {
	var i = 0
	defer func() {
		fmt.Println(i)
		i++
		fmt.Println(i)
	}()

	return i
}

func main() {
	tt()
}

```
在这个函数中 执行顺序是这样的 首先先初始化 i  = 0
然后 defer无法初始化（参数变量）因为它没有
然后到了return  然后 直接执行了后面的内容 没错 什么都没有就是i而已
然后 开始了return 直接将值返回了，这时候它没有结束 因为有defer
所以开始执行了defer，然后defer中i是几？嗯 是0 因为这个时候i是0了
然后 打印了0 i再次++ i等于1了，然而并没有什么机会去用这个i了，因为
已经return过了，所以这个i就被收回了，加入return后面是一个闭包，那么这个i
就有用了，它就不会被收回。

然后 这个时候函数就结束了。

看一下 这几种特殊情况


```go
// return 1 "defer tt1 0"
func tt1() int {
	var i = 0
	defer fmt.Println("defer  tt1", i)
	i++
	return i
}

// return 1 "defer tt1 0"
func tt2() int {
	var i = 0
	defer func(i int) {// 参数复制，值的复制。
		fmt.Println("defer tt2:", i)
	}(i)
	i++
	return i
}

// 1 1
func tt3()int{
	var i =0
	defer func() {
		fmt.Println("defer tt3",i)
		i++
	}()
	i++
	return i
}

//1 2
//2
func tt4() func() int {
	var i = 0
	defer func() {
		fmt.Println("defer tt4:",i)
		i++
	}()

	i++
	return func() int {

		return i// 引用变量。
	}
}
--------
package main

import "fmt"
// 0 13
func tt5() (num int) {
	defer func() {
		fmt.Println("dd", num)
		num++ //这里的num的作用于属于上层的函数体
	}()
	return 12// 返回值是函数运行最后的num值，很明显是defer中的num值。然后返回1
	// 这种形式的返回值，因为return 后面什么都没，所以它就会查找这个函数域内的最后值。因为函数的运行是
	//运行return后面的的内容 + defer + 隐藏的os.Exit()
	普通模式下是 运行return后面的内容 +返回return后面的内容 +运行defer + os.Exit()

}
func main(){
	fmt.Println(tt5())
}

func main() {
	fmt.Println(tt5())
	fmt.Println(tt6())
}
//dd 0 dd1 2 return 1
func tt6() (num int) {
	defer func(num int) {//// 使用这种形式 defer的函数内部的值已经是num的一个拷贝了，所以它里面怎么改变都不影响外部的return
		fmt.Println("dd", num)
		num++// 这里的num属于本层函数的作用域，所以它无法改变外层函数的数值。
		num++
		fmt.Println("dd1", num)
	}(num)
	num++
	return

}



```

4. 变量的调用和直接改变参数本身

```go
func tt(){
var i
dd(i)
fmt.Println(i)

}

func dd(i int){
i++
fmt.Println(i)
}

// 1 0

```
原因就是dd(i)这个是代表了 赋值，也就是将var的值赋予了dd的形参
可以看做是 d = i
dd（d） 这个叫做赋值 然后值的拷贝或者是指针的传入以及指针的获取实际值是这个地方的问题

然后还有一种是这样的

```go

func tt(){
var i = 0
{
i = 12
}
fmt.Println(i)
}
```
我之前因为跟赋值搞混所以我总是是用
指针来更更改i其实是错误的理解，因为这个地方的i = 12 压根就没有赋值
这一种说法它不过是更改自己的值而已，就像上面的那个函数即使使用指针，
那么更改指针的实际值的时候也是这么干的，所以i = 12 只是这个参数的在
调整自己的值罢了，它是改变的自己，这里就不牵涉到 是值的拷贝还是引用的拷贝了
因为它压根没有拷贝，仅仅是改变自己罢了。

5. 值的方法和指针的方法

首先 指针的方法和值的方法可以互相调用，因为go会自动帮你
比如指针的方法g（）你使用了值来调用那么go会帮你自动取地址相对的如果是
值的方法 go自动帮你取 *

第二个地方
首先 值上面可以有方法 指针上面也是有方法，我们谈的是 关于对象的方法这点先阐述
因为除了struct（对象）其它类型除了 指针和nil都可以有自己的方法
其它类型不讨论

就是指针的方法的时候，那么go会自动帮你取这个对象的*，因为指针没办法取得对象
里面的value值，只能*去得到，但是go帮你取了，所以你可以使用看似
指针直接去值。

6. 关于实现接口

这个地方go很严格，首先就是接口类型的变量不允许取指针，本来它就是引用类型了（初始化是nil）nil取不到method。
虽然slice这种也是引用类型但是go允许你取它的指针，但是接口类型不允许取（取了也没有意义）

而且对于实现一个接口来说如果你是指针类型实现的接口，那么将变量传递给指针类型时

也必须是指针类型的变量，值同样，不允许自动取&或者*
举个例子

```go
type a interface{

get()
}
type b struct {
c value
}

func(b1 *b)get(){
fmt.Print(b1.c)// go帮你自动取了 *b1这个地方不变，go帮你自动取对象的值，在这个地方也可以。
}
func ddc(a1 a){
a1.get()
}

func main(){
b1 := new(b)
var b2 b
ddc(b1)// 正确
ddc(b2)// 错误❌
}



```
7. 关于slice

slice赋值的时候可以不用指定类型

```
type t int
slice ：= make（[]t,10）

slice = [][]int{
{1}, //不需要使用 t{1}
{2},
{3},
}

```

如果slice里面还是slice或者是map等这种引用类型的话是这么处理的

```go

slice ：= make（[][]int,10）
slice[0] = make([]int,4)
// 或者
slice[1] = []int{
1,2,3,}

```
因为 引用类型不初始化的话 本身就是nil 所以会panic


8. 关于变量的初始化

关于这个地方我也出错过

```
func dd(t *int){
fmt.Println(*t)
}
func main(){
var t *int

dd(t)


}


```
这样就会出错，因为 所有的变量都会初始化（go没有声明 go会自动初始化）
但是t是个指针类型，它的初始化就是nil，所以*nil是错误的，正确的方法是


```go

func dd(t *int){
fmt.Println(*t)
}
func main(){
var t int

dd(&t)


}
```

或者优雅一点

```
func dd(t *int){
fmt.Println(*t)
}
func main(){
t := new(int)

dd(t)


}
```

8. slice 关于他的len和cap

不要超过它的len来查找数据。(而不是cap)只要是超过了len就会报错，虽然没有超过cap
但是它的out of range 错误是根据len来定的。

9. 不要获取map的值的地址

```go
t := make(map[string]string)
t["12"] = 12
fm.println(&t["12"])
```
因为map是动态的，所以它的value的值 的地址不是固定的所以go不允许取得
它的地址。

但是 slice可以。

```go

b := make([]int,12)
	b[1] = 12
	fmt.Println(&b[1])//0xc00001e128 slice 可以

```

10. recover的使用只能在 defer中使用(其它地方调用无效果)

```go

func tt(){
defer func(){
if t := recover;t {
    fmt.Print(t)
}
}
dd()//dd里有panic
}

```

11. 关于接口类型的断言
  - 接口实例.(接口类型)
  - 接口实例.(实际类型)

但是这两个的前面 无一例外都需要传入实际的类型也就是变成了
  - 实际类型的实例.(接口类型)
  - 实际类型的实例.(实例类型)

举个例子
```go
type a struct {
  value string
}

type b interface {
  get()
}
type c interface{
  post()
}

func tt(b1 b){
  // 第一种情况
  if v,ok := b1.(c);ok {// 这个实例 相当于实现了这两个interface
    fmt.println(v.post())
    fmt.Println(v.get())
  }

  // 第二种情况
  if v,ok := b1.(a);ok {
    fmt.Println(v.get())
  }
}
```
再看一个实际运用上的例子：
```go
type a struct {
	value string
}

type ber interface {
	get()
}

type cer interface {
	post()
}

func (a1 a) get() {
	fmt.Println(a1.value)
}

func (a1 a) post() {
	fmt.Println(a1.value + "p")
}

func t(b1 ber){
  // 这个内部的cer必不可少。
	type cer interface {// 这就是为了验证 已经实现了ber的变量是否也实现了cer
		post()
	}
	if v, ok := b1.(cer); ok {// 这个地方隐藏的说明了  a的实例是满足ber的，不然它这一步就会panic然后它还得满足cer不然还会panic所以这一步直接验证了两次。
		v.post()
	}
	b1.get()

}
```

12. 关于递归，
递归其实就是在执行函数里的函数，直到所有函数都结束了，然后就结束即可。举个例子


```go
func testVisit(ii int) int {
	if ii == 0 {
		return 100
	}
	fmt.Println(ii)
	ii = testVisit(ii - 1)

	return ii+1
}

```
它的执行很明显是从外层的初始栈开始往里执行，然后所有栈执行完毕即可，这里我使用了`ii = testVisit(ii -1)` 目的有两个，1 为了让每下一个的ii都少1，2 就是为了获取上一个栈的返回值，然后每次返回都+1 最后的返回值是109 这也证明了每次返回都是从最上层的栈开始往下调用然后到最下面的然后返回。

13. 关于 channel

 chan 的机制是这样的，当一个没有缓存的（有缓存也是一样只是当缓存满了就一样了）chan，显示导入一个数据，这个时候
 这个发送chan的goruntine就睡眠了（阻塞）然后直到这个chan被接受（只要被接收就行，不管是不是在同样一个goruntine)然后这个数据就被获取了，然后开始唤醒这个chan的发送者的那个goruntine。如果没有后续的数据那么这个chan就应该被关闭了可以人工关闭(close)也可能被系统收回。

看一个例子 这是一个有缓存的，并且利用缓存来限制 http请求数量的操作

```go
var st = make(chan struct{},20)// 将访问的数据限制在20
var sy synv.WaitGroup
func main(){
  dd := []string{"htps://...",",,,,,",",,,"}
  sy.Add(len(dd))
  for _,v := range dd {

      go read(dd)//
  }

sy.Wait()
fmt.Println("执行完毕")
}
func read(st){
  defer sy.Done()
  st <- struct{}// 因为是有缓存的chan所以可以保证一直有20个gorutine是不阻塞的。
  // 只要有一个goruntine不是阻塞的就不会造成死锁
  rea(st)
  <- st
}
func rea(st string){
  res,err := http.Get(st)
}
```

只要有一个goruntine不是阻塞的就不会造成死锁，死锁是程序想退出，但是chan内还有东西，没办法退出，但是又没办法运行，造成了无法结束的窘迫，最终就是各个goruntine都是阻塞然而又不能退出的局面。总之 死锁问题有必要再开一个文件来讨论一下。

14. 关于 type

alias的类型和底层可以转化但是不是隐式是显式。

这里分几个内容
- 一就是

```go
type hand func(http....,http.....)

// 例如

httprouter.handle("/",httprouter.handle)
// 这个时候就是
httprouter.Handle("/",func(http....,http....))
// 即可。
```

这种类型的尤其是在函数的调用的时候  要满足 一个hand类型也是很简单 就是函数满足后面那个样式即可

- type hand string

这种情况也是 函数满足后面的那个 type即可 也就是 是string即可 。

15. 关于引用类型

举个 slice说明一下

```go

func main(){
  t := make([]int,0,10)
  t = []int{
    "12",
  }
  visit(t)
  fmt.Println(t)
}

func visit(t []string){
  t = append(s,"1221")
}
```
猜一下 输出的是什么？
是 ["12","1221"]吗？
我本来一直以为是，后来我发现其实不是，我们要先证实一个问题，引用类型并不是指针，它是一个数据结构通常是 一个cap 一个len和一个指针对象。所以它本身也是一个实际的值。当这个地方把t传入visit后，其实是值的复制，然后在visit中，t等于了一个append返回的一个新的slice，那么它就不是指向了原来的那个底层数组了，(换言之，这样的话就不是改变底层数组了，是重新分配了一个数组，那么原来的那个slice自然就跟这个新的底层数组没有关系了)那么什么时候会改变呢，也很简单

```go
func visit(t []string){
  t [1] = "112"
}

```

这就叫做赋值，它是直接操控底层的数组进行了值的改变，这并没有去进行值的拷贝或者是指针传递。到这里我么可以说一下了
在go里所有的类型只要是传递数据只有两种模式，1 值的拷贝（包括引用类型它的拷贝只不过是拷贝的它的数据组织）2 指针的拷贝，说白了，指针的拷贝也是值的拷贝，因为它本身也是一种值只不过象征了一种钥匙罢了，所以，除了对数据本身进行直接改变，改变他的数据本身，这种行为可以改变它自己，值的传递的话 统统是有拷贝行为。所以我们以后不能把引用类型看成指针，很不一样。如果是指针的话 那么肯定必须要获取值才能去改变，但是引用类型是go的编译器自动的行为。(例如扩容啊，自动获取底层数据的值啊这种)


16. 关于带（bare return）形参及不带形参的返回值和defer的故事

带bare的那种 ，最后的栈尾是 包含了defer了的，不带bare的不包含defer 
举个例子

```go

func a(i int)(t int){
defer func(){
t+=i
}
return 1
}
a(12)
那么结局是 13 过程是  在return1的时候这个时候t就是等于1
的但是因为defer栈还没结束，那么开始执行了最后的defer了，然后return出去了的值，那这个值就是 13

func a(i int)int{
defer func(){
t+=i
}
return 1 // 因为是不带bare的那么 return的栈就直接出来了，defer是不经过这个return栈的 defer之后调用了os.exit() 所以 1 就是最后的返回值。
}


```
看两个例子：

> 之前的例子说过了，defer只是执行滞后但是参数记住是参数也就是将形参传入实参的过程其实是同步的并没有什么区别。

```go
//  这个例子中返回值是1
func tt()(t int){
	defer func() {
		t ++
	}()
	return
}

// 这个例子中返回值是0
func tt1()int{
	t := 0
	defer func() {
		t ++
	}()
	return t
}
// 0 这个例子证明了 参数顺序执行化。
func tt1()int{
	t := 0
	defer func(t int) {
		t ++
	}(t)
	return t
}
```

原因也是很简单，首先如果是没有形参的返回值，都是在return后面直接返回的，然后再执行defer然后再执行 os.exit() 但是有形参的就不一样了，它必须返回它形参定义的参数之歌例子中就是t，那么t在哪最后一个出现呢？就是在defer中，所以它的执行过程就变成了，找寻最后出现的t（这里出现在defer中）然后直接执行os.exit() 因为它return后面没有东西，所以它和没有形参的return XXX 很不一样。
那么如果是这样的呢？
```go
func age()(n int,err error){
  return
}
```
它会有什么返回结果呢？答案就是`0 nil` ---- 如果只有return 但是却没有出现n和err那么简单 返回值里不是已经初始化了嘛，那么久返回初始化的结果不就好了嘛所以是 `0 nil`(他们的初始化值)
17. 关于buffered

我们在go的执行中经常使用的一种技巧就是限制go并发的速度，那么这个时候buffered变量就可以实现了它的实现是这样的 ` make(chan xxx,number)` 在get请求中一般我都会这么使用`make(chan struct{},20)` 我们定义了一个新的类型就是 struct{} 这个类型是代表了空，当然你也可以使用bool 都可以 struct{} 类型使用的时候 用 struct{}{} 即可。这就代表了这个chan中最多可以暂存number个数据，这就是所谓的缓存技术，也叫做 buffered数据

18. 关于 recover和并发（多goruntine）

如果是在go的多协程中的panic一定要在这个协程中recover否则在主协程的recover根本无法获取这个panic

```go
go func(i int) {
			defer sy.Done()
			defer func() { // 如果是在外部获取recover可以说压根获取不了，想想也是知道的因为你并不知道主协程和这个协程到底哪个运行到哪了，所以要在这个协程中搞定这个panic
				if e := recover();e != nil {
					fmt.Println(e)
				}
			}()
			start := time.Now()
			resp, err := http.Get(url[i])

			if err != nil {
				fmt.Println(err)
			}
			n, err := html.Parse(resp.Body)
			if err != nil {
				fmt.Println("err",err)
				return
			}
			defer resp.Body.Close()
			if err != nil {
				fmt.Println(err)
			}
			wor, im := countWordsAndImagesAsync(nums, ch, n)
			ma.Store(url[i]+"   num", wor)
			ma.Store(url[i]+"   image", im)
			end := time.Now()
			timeS := end.Sub(start)
			ma.Store(url[i]+"花费的时间是：",timeS.String())
		}(i)
```
19. 关于递归的出栈和进栈
递归都有一个进出栈的过程，
```go
func a(){
  visit(start,end)
}
func visit(start,end func()){
start()//在进栈时执行的函数
  for {
    visit()
    if XXXXXX 然后退出这个栈开始出栈
  }
  end()// 在出栈时执行的函数。
}
```

20. 关于 函数内部的函数

```go

func t(){
  var d func()int // 使用这种方式一般都是函数内部有递归，如果不实现 声明一下 函数内部的递归函数将无法运行

  d = func()int{
    //fdffd
  }
  d()

  // 或者
var d = func(){

}
d()

总之，不能使用
func()int{

}
int() 在go语言中这种行为不允许

}
```
关于函数内部声明类型 倒是很随意
```go

func t(t1 inter){
  type t struct{
    get()
  }
  if d,ok := t1.(t);ok {
    d.get()
  }
  t1.post()
}



```
21. 只有接口和nil不能拥有方法。

```go
invalid receiver type io.Writer (io.Writer is an interface type) // 这是使用了接口的报错。

nil is not a type// 这是使用了nil的报错
```

> ps: 永远不要去取接口的指针，没有丝毫的意义。如果取 slice的指针还有些许的意义(比如在append的时候)但是接口的指针有什么意义？
接口本来就没有实际的意义它本来就是一个抽象的东西。而且它本来也就是引用对象。


23. time.After 的用法

它的作用是 当这个系统没有东西了，然后在设置后几秒后运行，如果其他的case一直有东西，那么它是不会被执行的。
因为在select选择的时候只有在其他的case都没有反应了的时候才会去选择time.After 所以它可以用在 比如什么东西都没有数据了以后
然后按照某时间后去取消这个东西，当然还有一种应用场景就是，无论如何就是要5分钟后取消，打死都要
那么可以 使用两个select，有一个select就放一个after和一个default即可。
```go
select{
case time.After():
return // 这样就强制 退出了
default:

}

select {
// 这个select就是干正事的。
}
```
或则，使用 context包的withcanceltimeout 这个函数厉害 无论如何 只要设置的分钟数到了，就能立马取消。
因为cancle withcancle那个函数 如果 不执行cancle函数 那么ctx.done 就无法运行，这个时候 cancle的关闭就要在执行这个有ctx的函数之前了。就不能使用defer函数来关闭这个，因为一直有东西运行。

24. 关于 string字符串 []byte 以及[]byte的十六进制表示（以string形式储存）

```go
	// 将string字符串，以unicode编码的形式，找到所有的字符的unicode表示，然后返回位一个数组。
	// [72 101 108 108 111] 就是这个数组（slice）
	src := []byte("Hello")
	// 这个encodeStr 是什么呢？它其实就是把这个数组的所有的数字用16进制表示并且没有加[]而已，而是将这个串变成了字符串的形式储存
	// 就是这个“48656c6c6f” 这个字符串其实还是unicode编码只是 用的16进制并且没有[]罢了，一定不要认为它就是"HELLO"
	encodedStr := hex.EncodeToString(src)
	fmt.Println(src)
	// 48656c6c6f -> 48(4*16+8=72) 65(6*16+5=101) 6c 6c 6f
	fmt.Println(encodedStr)
	byteValue,_ := hex.DecodeString(encodeStr)
	string(byteValue) == "Hello"
```
25. 关于json的一个解析的问题

```json
{"code":0,"data":{"ip":"173.82.115.125","country":"美国","area":"","region":"加利福尼亚","city":"洛杉矶","county":"XX","isp":"XX","country_id":"US","area_id":"","region_id":"US_104","city_id":"US_1018","county_id":"xx","isp_id":"xx"}}

```
```go
type Data struct {
	Data Values `json:"data"`
}
type Values struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

```

定义的时候可以缺少字段，但是，不能跟json字段的格式不符合，举个例子 这里的数据是 在json整个文件下的data对象中，那么你需要两个struct 一个是代表整个的json的数据，第二个struct是代表那个data，你看 那个code和其它字段没有定义吧，没有定义无所谓，但是字段的格式一定要遵守 如果直接把Values传进去就是错误的行为，是不会解析的。

26. 关于go template

第一点 如果你使用`template.Execute` 那么你在那个最外边的layout那个文件里不能使用`{{define "layout"}}{{end}}`
如果你想使用`{{define "layout"}}{{end}}` 那么 你需要`	tem.ExecuteTemplate(w,"layout",nil)`那个中间的变量要用最外边的那个模块
所以最好的就是最外边的那个不用模块，然后使用那个没有模块的就ok了

第二点 如果如果你的子模块就是小的模块很多人称作是母模块 我当他们是小模块子模块，他们里面有变量，那么你肯定是最后使用的是layout这个最外面的文件或者说是模块
那么你就要`{{template "son".}}`   `.` 看到了吗 这个点没有这个点 你在最外面的模块也就是最终使用的时候你发现你的变量压根没有导入，这个就是 变量导入的标志
也就是是 你的子有了 如果不导入 那么这个数据就消失了，我被这个地方坑了几个小时。我的天~~~~。

27. 关于iota

这个东西是有几个特点的其中
- 从0开始自动计数
- 中间可以有间隔，然后重启计数，并且仍然是按照原先的顺序进行计数
```go
package main

import "fmt"

const(
	a = iota // 从0开始
	b // 按照iota的特性可以继续往下计数，并且继承它的类型
	c
	d
	e
	f = "1" //  可以中断
	tew // 在没有重启之前一直是按照中断时候的定量来进行赋值
	weew
	we
	r
	t
	h
	cd = iota // 直到继续重启，数字是重启了， 但是值类型变成了系统默认值 int
	gt
	ut
	yyt
	tyy
)


func main() {
	fmt.Println(a,b,c,d,e,f,tew,weew,we,r,t,h,cd ,gt,ut,yyt,tyy)
	fmt.Println(reflect.TypeOf(a),reflect.TypeOf(b),reflect.TypeOf(f),reflect.TypeOf(gt))

}
// output:
//0 1 2 3 4 1 1 1 1 1 1 1 12 13 14 15 16
//int64 int64 string int

```

### 关于range使用的是复制值的 陷阱

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	for k, v := range a { // range后面的a 是对于 上面的a的复制值 ，但是由于slice底层array连带变了,所以v值就会产生变化
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	fmt.Print(a)

	a1 := [3]int{1, 2, 3}
	for k, v := range a1 { // 这里的a1 是上面a1的复制值，所以v永远就跟之前的那个样子一样，纵然 真实的a1发生了变化但是复制值是不会变化的
		if k == 0 {
			a1[0], a1[1] = 100, 200
			fmt.Print(a1)
		}
		a1[k] = 100 + v
	}
}

//[100 200 3][101 300 103][100 200 3]


```

### 关于range时的指针问题  -- 动态数据问题

```go
func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    // 错误写法
    for _, stu := range stus {
        m[stu.Name] = &stu  问题出在哪个地方了？ 很简单 所有的 m[xxx] = "同样称呼的动态地址" 也就是说 大家都是等于 &stu 除非每次等于的东西不一样否则 那最后取得的值肯定一样。 除非是 等 &stu+i i++ 就是这么个问题，为什么 stu.name 没问题原因人不是动态类型啊 直接取得到真实值了呗。
    }
 
    for k,v:=range m{
        println(k,"=>",v.Name)
    }
 
    // 正确
    for i:=0;i<len(stus);i++  {
        m[stus[i].Name] = &stus[i]
    }	
	for k,v:=range m{
        println(k,"=>",v.Name)
    }
}


```
### 关于defer和参数函数的问题

```go
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) //我们还记得 不管是不是defer 他们的参数初始化是顺序执行的，那么 里面的calc 作为参数就会先执行所以“10” 就会先执行 然后再是 “20” 然后 defer的真实执行 又开始是 “2” “1”  所以 他们的顺序是 10 20 2 1 反正记住一条就哦了 就是 参数的初始化是顺序执行的。
	
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

```
10 1 2 3 
20 0 2 2
2 0 2 2
1 1 3 4

### 关于copy的地方

```go

s := []int{1,2,3,4}
copy(s,[]int{10,332,22})

结果就是 [10,332,22,4]被覆盖了

copy(s,[]int{11,21,31,41,51,61,7,8})
[11,21,31,41] 按照小的那个的len 所以就被丢弃了。

所以一般前面的那个slice就应该是空的。
```

### append的容易错的地方

```go
 append 是添加到前面的那个slcie的后面
 
 举个例子 如果是
 d := make([]int,5)
 t := append(d,[]int{1,2})
 o: [0,0,0,0,0,1,2] 就是这样了。 跟copy不同
```
### 结构体的比较

进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与**属性类型**以及**个数**有关，还与**属性顺序**相关。
还有一点需要注意的是结构体是相同的，但是结构体属性中有**不可以比较的类型，如map,slice**。那么也不能比较， 如果该**结构属性都是可以比较的**，那么就可以使用`“==”`进行比较操作。

### go 的定量const 的值是无法取得到地址的。因为它的存在是在编译期间就获得了，不是在运行期间。

### goto不能跳转到*其他函数*或者内层代码

```go
func main(){

for {
loop: // 内层
}
goto loop
}
```
