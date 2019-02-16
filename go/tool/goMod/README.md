## go mod go官方的module管理工具

#### 用法：

在一个非go path的路径中新建一个项目，然后使用`go mod init` 就可以初始化一个新的包，其实跟github(gitlab都行)用在一起更好

1. 在github上新建一个项目，例如说 test
2. 在本地将这个远程包给clone过来，然后这个文件夹里面就是一个.git 隐藏的文件项目这个就是git的管理文件包
3. 将此包放在远离 go path的文件路径里，然后使用 `go mod init` 就可以创建一个名为 github.com/XXX/test的包
记住，这个包名不能随意称呼，你的文件夹的名称就是你的包名然后里面的XX.go的名字无所谓，只要`package test`
用对就可以了，然后你会发现你的包内出现了两个文件 **go.sum** 和 **go.mod**  这个sum包你可以忽略，主要是go.mod包
这里的包 首先开头是 module github.com/XXX/test 声明了 这个包的具体名称是 **"github.com/XXX/test"** 但是在调用的时候 包的名称是 test
为什么module后面要加上所有的路径呢，原因也是很简单，就是当你 `go get github.com/xxx/test`的时候用，不然你只有一个test
go get是没办法下载这个包的。
4. 将这个包发布到github即可。

## 版本

- 你的包的具体version其实是 git来管理的 1 你可以使用 git push -vesion 来指定某一个大的version 第二你可以选择不写version
然后别人用你的包的时候就不能指定这个version了只能用 latest来引入包 通常来说是你最后编辑的日期

- 版本 version2.X version 3.X该怎么办？go mod官方推荐的方法如下

```go
module github.com/xxx/test/v2

```
也就是说你的包的名字还是test但是因为版本是v2.X 所以要在包名后面加入大版本号 /v2
然后 调用的时候是这样的

```go

package dd

import(
"github.com/xxx/test/v2"
)

func dd(){

test.Test()
}
```
也就是说使用imprt的时候 也要加 v2大版本号，但是在函数内调用的时候 还是test 比如这个时候你的版本是v2.45
那么 你的mod文件里最后的标注就是v2.45

> 就算不是v2.45  不通过打tag的方式来发布 比如还是默认的数字，那么只要你在mod中指定是v2 你的版本大号就是v2
在import中引入还是需要加上v2 只是在mod中 显示的信息是你的最后编辑时间。

综上：

```go
// 你的包
go.mod 文件写法 : module github.com/xxx/test/v2

别人用你的包的写法 import(
"github.com/xxx/test/v2"
)

func dd(){
test.Test()
}
```

即可。

### 包的存放位置

再使用了go mod后你的go path将没有用了，但是存放的包的位置还是在 老的go path 更明确来说是在 $gopath/src路径
这个路径会有两个文件 pkg bin 前面这个包是存放的非可执行的包 后面的bin放置的就是可执行的包，你可以把path指向这个
bin即可。

### go mod tidy

这个命令很有用 首先我们看看它的官方解释 **tidy: add missing and remove unused modules**
也就是说 你的go.mod中多引入的或者是少引入的 使用 `go mod tidy` 它可以帮你处理

### 我的包是可执行的文件，我发布到github上让人使用我的mod该怎么办
你的go.mod 可以写成 `github.com/xxx/test`但是你的文件的package 应该写成main 或者目前而言在github上的
可执行文件的包没有go.mod也可以，你只需要在本地开发的时候有go.mod即可。然后这个go.mod 中的名称不跟package一样也可以
因为别人是不会再引用你的包的了。但是如果你的包有子包，那么你还是应该把你的包go.mod文件里的 module 后面规规矩矩的写例如
```go
module github.com/XXX/add

然后别人引用的时候
import(
"gitub.com/xxx/ddd/app" // app 是子包 即可。
)
```
### 我无法使用goalng.org/x的包我该怎么办

例如:
你本地的包要引入 golang.org/x/net/html
但是被封了，那么你可以使用github上的镜像包 例如说是 github.com/golang/x/net/html

在你的项目的go.mod 中 加入 replace golang.org/x/net/html => github.com/golang/x/net/html
但是一般你科学上网不就行了吗。。
### 我该怎么处理我的子包和我的包的关系
- 本地
在本地 比如说你的大包要引用子包的内容你可以go.mod 中使用replace，比如github.com/app/中
要引入 github.com/app/app的东西，你可以 在go.mod 中 用 replace github.com/app/app => "./app" 即可
当,你发布的时候你把这个replace删掉即可。

### 当我的文件夹的名称跟我的package写的包名不一样怎么办

要记得 一个文件夹（包）内的的package 名称必须一样，但是可以不跟文件夹的名称一样，使用的时候其实很不爽就是了
不过也可以用 用法是

比如一个包 它的文件夹是 app 但是 package中的名称是 app1 那么可以这么用

```go
imoirt "github.com/xxx/app" //  这个地方要跟文件夹的名称保持一致 其实就是跟路径保持一致
func dd(){
app1.xxx// 这里要跟真实的package保持一致 你看 得不偿失吧 所以 文件夹的名称要跟package的名称保持一致

}
```

### 注意

你的包如果存放在github上 你的包的go.mod module后面一定是github.com/xxx/xxx不能直接写成 xxx 这样的话
go mod 无法获得包 错误是`parsing go.mod: unexpected module path "test"
                 go: error loading module requirements
`
总结：
- 也就是说 go.mod 的module 要跟go get xx/xxx 保持一致 例：`module github.com/app/app` 在使用
开启 go mod后 使用go get 的时候也是 `go get github.com/app/app`

- 文件夹可以不跟package的具体包名保持一致，import的时候是用的文件夹名 函数中时候是用的package的名字

- 有main包的时候 go.mod 要跟文件夹的名称保持一致，子包该怎么引入就怎么引入 例如 github.com/app/dd
dd是子包 app这个路径中是main包。

### ⚠️

- 文件包的名称和package的名称要保持一致(main包除外)
- 要用go mod 代替 go path 和dep 大势所趋
- go mod中的module名称一定要跟gitub上的路径（其实是git的路径，这个路径没有github上的tree/master）保持一致。不然没办法拉取
- 子包和包不能互相引用可以小引大也可以大引小但是不能互相。
- 不要把项目放到go path中了，go path要取消了。（我猜的）
