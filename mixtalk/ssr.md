# 如何科学上网

- 选择一家服务器厂商 例如谷歌云 或者[cloudcone](https://app.cloudcone.com/?ref=2525)

-  购买服务器（vps，或者叫做虚拟机都有）

- 服务器请安装Linux，版本不限

- 安装[go语言](https://golang.google.cn/dl/)

- 按照说明安装好go语言后 安装ssr-server：go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-server

- 在/home/go/bin 中找到shadowsocks-server 在同一个路径下新建一个config.json的文件文件内容[为](https://github.com/shadowsocks/shadowsocks-go/blob/master/config.json)

- 使用root去运行shadowsockets-server即可 `sudo nohup ./shadowsocks-server &`

- 本地下载好一个ssr的客户端配置信息后即可科学上网
