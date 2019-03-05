## go mod go官方的module管理工具
> youtube https://www.youtube.com/watch?v=saJ2c006vp4
#### 用法：

在一个非go path的路径中新建一个项目，然后使用`go mod init` 就可以初始化一个新的包(要开启这个 `export GO111MODULE=on`写入.bash_profile即可 win的同学自己找找设置 GO111MODULE的win版本设置方法哈)，其实跟github(gitlab都行)用在一起更好

1. 在github上新建一个项目，例如说 test
2. 在本地将这个远程包给clone过来，然后这个文件夹里面就是一个.git 隐藏的文件项目这个就是git的管理文件包
3. 将此包放在远离 go path的文件路径里，然后使用 `go mod init` 就可以创建一个名为 github.com/XXX/test的包
记住，这个包名不能随意称呼，你的文件夹的名称就是你的包名然后里面的XX.go的名字无所谓，只要`package test`用对就可以了，然后你会发现你的包内出现了两个文件 **go.sum** 和 **go.mod**  这个sum包你可以忽略，主要是go.mod包这里的包 首先开头是 module github.com/XXX/test 声明了 这个包的具体名称是 **"github.com/XXX/test"** 但是在调用的时候 包的名称是 test为什么module后面要加上所有的路径呢，原因也是很简单，就是当你 `go get github.com/xxx/test`的时候用，不然你只有一个test go get是没办法下载这个包的。
4. 将这个包发布到github即可。

## 版本
> 在`go.mod`中的包后面手动输入latest `github.com/nfnt/resize latest`，go将自动(go list)帮你生成一个版本号 `github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646`

- 你本地的项目引用别人的包的时候可以在go.mod 中指定version的版本，但是什么都不指定也可以，默认是latest，也就是你直接`go get github.com/xx/app`的时候它自动就是引入latest的版本了，要指定某个版本，你在go.mod 改了就行了。当包中未使用具体version的时候，第三个人用你的包时候然后默认下载你的包引用的包的版本，indirect的包（非直接引用的包）的版本是最新的那个版本，这里的最新就跟你的最新不是一回事了，如果引用的包的go.mod中明确指明了version而不是latest这种形式，那么你引用它的时候它的间接包的版本默认跟你的是相同的，不过你自己也可以更改的哦（能明白吧我感觉有点绕，😝）
举例子：

这是gin的go.mod摘选
```go
	github.com/gin-contrib/sse v0.0.0-20190124093953-61b50c2ef482
	github.com/golang/protobuf v1.2.0

```
这是我本地项目的go.mod摘选：
```go
github.com/gin-contrib/sse v0.0.0-20190125020943-a7658810eb74 // indirect
github.com/golang/protobuf v1.2.0 // indirect
```
你会发现同样都是v1.2.0 但是上面的那个 一个是 20190124 一个是20190125,我们看一下下载包的时候的bash的截图
```bash
finding github.com/stretchr/testify/assert latest
```
如果你用的包它用的包，它没有指定具体的版本，那么你的latest跟他的就可能不是一回事了，这个要清楚。因为go get 下载的时候是此时此刻的latest。

`require github.com/coastroad/test v0.0.0-20190216094021-0555a706efff // indirect`
这里 后面就是指定的版本 // indirect就是间接引用的意思。
- 你的包的具体version其实是 git来管理的 1 你可以使用 git push -tag 来指定某一个version.
第二你可以选择不写version
然后别人用你的包的时候就不能指定这个version了只能用 latest来引入包 通常来说是你最后编辑的日期例如`v0.0.0-20190216094-55a706efff`

- 版本 version2.X version 3.X该怎么办？go mod官方推荐的方法如下

```go
module github.com/xxx/test/v2

```
也就是说你的包的名字还是test但是因为版本是v2.X 所以要在包名后面加入大版本号 /v2(一定是v+数字 比如 v2 v3 v4 v5不能改变写法也不能用v2.3这种写法)
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

还有种写法是
```go
module github.com/xx/app.v2
```
这种写法也是类似，只是官方不推荐这种写法。
> 作为一种特殊情况，以gopkg.in/开头的模块路径继续使用,在该系统上建立的惯例：主要版本始终存在，它前面有一个点而不是斜杠：gopkg.in/yaml.v1和gopkg.in/yaml.v2，而不是gopkg.in/yaml和gopkg.in/yaml/v2。

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
### 包的存放位置

再使用了go mod后你的go path将没有用了，但是存放的包的位置还是在 老的go path 更明确来说是在 $gopath/src路径
这个路径会有两个文件 pkg bin 前面这个包是存放的非可执行的包 后面的bin放置的就是可执行的包，你可以把path指向这个bin即可。

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
### 我go get的包，但是并没有在我的项目中使用，我该怎么处理我的go.mod
举个例子，你要拉取一个在github上的包，这个包呢，是一个可执行的包（main包）然后你需要在一个有go.mod的项目中使用go get才能拉下来，然后不出意外你的这个拉取记录就会出现在go.mod上,你不想让他出现该怎么办呢？很简单使用 `go mod tidy `就可以从go.mod 中删除。

### 我无法使用goalng.org/x的包我该怎么办

例如:
你本地的包要引入 golang.org/x/net/html
但是被封了，那么你可以使用github上的镜像包 例如说是 github.com/golang/x/net/html

在你的项目的go.mod 中 加入 replace golang.org/x/net/html => github.com/golang/x/net/html
但是一般你科学上网不就行了吗。。
### 我该怎么处理我的子包和我的包的关系
- 本地
在本地 比如说你的大包要引用子包的内容你可以go.mod 中使用replace，比如github.com/app/中
要引入 github.com/app/app的东西，你可以 在go.mod 中 用 replace github.com/app/app => ./app 即可
当,你发布的时候你把这个replace删掉即可。（仅限 *Unix系统，就是改变路径而已，win的同学自己看看咋弄就ok了，反正发布的时候要删除）

### 当我的文件夹的名称跟我的package写的包名不一样怎么办

要记得 一个文件夹（包）内的package 名称必须相同，但是可以不跟文件夹的名称一样，使用的时候其实很不爽就是了
不过也可以用 用法是

比如一个包 它的文件夹是 app 但是 package中的名称是 app1 那么可以这么用

```go
imoirt "github.com/xxx/app" //  这个地方要跟文件夹的名称保持一致 其实就是跟路径保持一致
func dd(){
app1.xxx// 这里要跟真实的package保持一致 你看 得不偿失吧 所以 文件夹的名称要跟package的名称保持一致

}
```
### 当我想忽略掉某包的时候我该怎么做？

```go

module github.com/x/x
exclude github.com/test/test latest // 或者把latest换成其它的版本比如v1.23.1 都可以。

```
### 该如何配置包的路径
举个例子，你有一个域名 `please.io`，有服务器，你想让被人下载的时候不使用github.com/xx/xx 而是使用自己的路径进行下载，那么你可以这么做

- 首先将你的go mod 中的module命名为你自己设定的路径 例如 `module please.io/tt` 然后还是放在github上，
- 在你的服务器上增加一个route，`/tt`然后在打开的这个页面的html的meta标签里添加上`<meta name="go-import" content="please.io/tt git https//github.com/xx/xx">`
- 搞定了。
### 注意

你的包如果存放在github上 你的包的go.mod module后面一定是github.com/xxx/xxx不能直接写成 xxx 这样的话
go mod 无法获得包 错误是`parsing go.mod: unexpected module path "test"
                 go: error loading module requirements
`
总结：
- 也就是说 go.mod 的module 要跟go get xx/xxx 保持一致 例：`module github.com/app/app` 在使用
开启 go mod后 使用go get 的时候也是 `go get github.com/app/app`

- 文件夹可以不跟package的具体包名保持一致，import的时候是用的文件夹名（路径）， 函数中时候是用的package的名字

- 有main包的时候 go.mod 要跟文件夹的名称保持一致，子包该怎么引入就怎么引入 例如 github.com/app/dd
dd是子包 app这个路径中是main包。

### ⚠️

- 文件包的名称和package的名称要保持一致(main包除外)
- 要用go mod 代替 go path 和dep 大势所趋
- go mod中的module名称一定要跟gitub上的路径（其实是git的路径，这个路径没有github上的tree/master）保持一致。不然没办法拉取
- 子包和包不能互相引用可以小引大也可以大引小但是不能互相。
- 不要把项目放到go path中了，go path要取消了。（我猜的）
- 最后 使用 `go help mod `有几个命令上面有提示。就不一一解释了。

```go
The commands are:

	download    download modules to local cache
	edit        edit go.mod from tools or scripts
	graph       print module requirement graph
	init        initialize new module in current directory
	tidy        add missing and remove unused modules
	vendor      make vendored copy of dependencies
	verify      verify dependencies have expected content
	why         explain why packages or modules are needed


```
