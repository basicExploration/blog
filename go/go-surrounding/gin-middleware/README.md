## 如果开发一个gin中间件

[具体看代码](./ginMiddlewareDemo.go)

所以说 gin的中间件原理就是，每一次的handle动作，都会通过中间件（你设置中间件的前提下）,

要想设置中间件，只需要符合 func(*gin.Context) 即可。