package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 测试，开发一个关于gin的一个中间件
func main() {
	engin := gin.Default()
	engin.Use(hui)
	engin.GET("/", func(context *gin.Context) {
		fmt.Println("/")
	})
	engin.GET("/test", func(context *gin.Context) {
		fmt.Println("/test")
	})
	engin.Run(":8080")
}


func hui(ctx *gin.Context){
	fmt.Println("我来测试一下，所有的输出全部都会经过这一层，然后输出我这一句话：😝")
}



//我来测试一下，所有的输出全部都会经过这一层，然后输出我这一句话：😝
///
//[GIN] 2019/03/16 - 10:19:37 | 200 |      30.598µs |             ::1 | GET      /
//我来测试一下，所有的输出全部都会经过这一层，然后输出我这一句话：😝
//[GIN] 2019/03/16 - 10:19:37 | 404 |      38.884µs |             ::1 | GET      /favicon.ico
//我来测试一下，所有的输出全部都会经过这一层，然后输出我这一句话：😝
///test
//[GIN] 2019/03/16 - 10:19:50 | 200 |      33.662µs |             ::1 | GET      /test
//我来测试一下，所有的输出全部都会经过这一层，然后输出我这一句话：😝
//[GIN] 2019/03/16 - 10:20:02 | 404 |      29.198µs |             ::1 | GET      /d




