## 一个项目在mac上开发然后发布到Linux上该怎么做？

https://nudao.xyz/

原文：https://github.com/googege/blog/blob/master/db/mysql/mysqlMacOSTOLinux/README.md

- 首先你先在mac上安装好mysql 然后安装过程 很简单了，因为mysql for mac是有可视工具的。

- 然后你在你的mac上的设置里最下面会看到一个mysql的标志你打开它然后可以更改信息，例如说 更改密码之类的

- 关于开发就是开发的过程了巴拉巴拉一堆省略了。

下面是关于如何发布的

首先我假设你使用一个工具，叫做`mysql workbench ` 这个工具是mysql官方出的，我个人用起来感觉还行，挺不错的，然后你去mysql官方下一个就ok了，用这个东西

你就不需要自己一个字一个字的敲sql了，然后这里要说的是如何导出数据，

在这个工具里，上面有个 Server 然后里面有个 EXport data 你点开，就会出现一个窗口，然后你配置好导出的路径，以及导出的数据库的表，然后导出即可。就ok了

***** 

其实在这之前都不是很难，我的意思是配置不是说开发，然后真正放到Linux上的时候有一些坑还是要注意的，我发现网上很多的坑，很多资料都是错的。然后我们来一步一步的看看到底该怎么配置

- 开始安装 mysql community 和mysql community server

```bash
sudo yum install mysql-community-server

```
- 然后开始启动mysql的服务器

```bash
sudo service mysqld start

```

- 然后查看一下mysql服务器的状态 看看是不是正在running

```bash
sudo service mysqld status
```

- 然后接下来你需要知道mysqlroot服务器的默认root密码 当然我本来也是一位不用密码就能进去 我记得有些教程就是这么写的，但是我貌似进不去总是提示错误，
所以我是这么做的 我找到了 默认的 初始密码

```bash
sudo grep 'temporary password' /var/log/mysqld.log
```
这样你将得到了密码 密码就是 : 后面的那一坨 例如 `: db?/fsdfdsfi  `

- 接下来你将使用默认密码去更改密码

```bash
mysql -uroot -p

```

进入 mysql的执行界面 然后

```bash
ALTER USER 'root'@'localhost' IDENTIFIED BY 'MyNewPass4!';
```
谨记一件事 你的密码 必须有大小写 数字和特殊符号，没有任何一项都创建失败

> mysql document : validate_password 默认安装。实现的默认密码策略validate_password要求密码包含至少一个大写字母，一个小写字母，一个数字和一个    特殊字符，并且总密码长度至少为8个字符。

- 然后这个时候你已经更改好密码了，mysql会退出然后你直接再次输入新密码即可，

- 到这一步基本上就快结束了，你只需要把导出来的 数据 也就是 数据和表头（忘了说了，你在workbench导出的数据记得要同时导出表头信息和信息）
然后倒入信息的时候不能没有数据库 举个例子 假如说你的数据库叫做 `example ` 那么你应该先在这个新的空数据库里 先 CREATE example 创建一个同样
名称的数据库 然后 退出来 control +d 退出后呢 再开始导入数据

- 最后一步 导入数据 

```
mysql -u用户名 <  example.sql  -p

例如 mysql -uroot < example.sql -p
```
就可以了，我这里要说一下 

在某网站里它是这么写的

```
mysql -u用户名    -p密码    <  要导入的数据库数据(runoob.sql) 
```
这其实会报错,原因是 mysql不允许 使用-p密码 这种 方式 只能使用-p 然后输入密码 才可以 

****
如何进行备份你的文件

```bash
mysqldump -u root mathcoolEnd > ~/Desktop/tt.sql -p

```
这会备份 你的所有文件，包括表头和表内容。

以上所有言论 均在 mysql community 8.x版本的言论，不涉及到 mariadb 以及 8.0以前的版本.

参考资料: https://dev.mysql.com/doc/refman/8.0/en/linux-installation-yum-repo.html
        http://www.runoob.com/mysql/mysql-database-import.html

