## pplabs 真实面试题



¶ 第一题：
下面显示了当前 PPIO 测试链上每个钱包地址所包含的余额数排行榜信息（数据为测试数据，仅供参考），如下所示：

```bash
A total of 100000 accounts found         上一页  < 1/200  >下一页

Rank  Address                                  Balance         TxCount
1     ppio1SDKRLjDALmN7kV874JLhZAv4Q9TKVhrhg   30000 PPCoin    2000
2     ppio1ZoKAKZftwkdAVXSuPnz2TLJp14pSfhvMj   29990 PPCoin    1999
3     ppio1Yd4y3cYYZZ3pZqxnmn6xydrE4VmvWiX3d   29980 PPCoin    1998

...
```

如果上面所示的排行榜数据都是放在 redis 数据库中维护的，为了实现以上功能，需要用到 redis 的哪些数据结构及相应的方法？

提示有以下几个功能点：

获取当前排行榜中的总地址（Address）数。
分页获取当前排行榜数据（根据余额数倒序排序）。包含 Rank（名次）、Address（钱包地址）、Balance（余额）、 TxCount（参与事务的总数）。
请给出详细的设计文档。

文档内容示例： 需要使用 String 数据结构来存储各个钱包地址的余额信息，键名为：PPIO:Wallet:ppio1SDKRLjDALmN7kV874JLhZAv4Q9TKVhrhg，键值为该账户的余额。 通过 GET PPIO:Wallet:ppio1SDKRLjDALmN7kV874JLhZAv4Q9TKVhrhg 来获取账户 ppio1SDKRLjDALmN7kV874JLhZAv4Q9TKVhrhg 的余额信息。

¶ 第二题：
下面给出某个文件（https://resource.testnet.pp.io/demo/release/macos/latest/ppio-demo.dmg）在下载过程中所涉及的 HTTP 请求包头及响应包头。

请求包头：
```bash
:authority: resource.testnet.pp.io
:method: GET
:path: /demo/release/macos/latest/ppio-demo.dmg
:scheme: https
accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3
accept-encoding: gzip, deflate, br
accept-language: zh-CN,zh;q=0.9,en;q=0.8
cookie: ppio_uid=testuid; ppio_token=testtoken
if-range: "2c1210c6567c9278985062f9dac53e57-10"
range: bytes=15332307-15332307
upgrade-insecure-requests: 1
user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36
响应包头：

accept-ranges: bytes
content-length: 1
content-range: bytes 15332307-15332307/78696231
content-type: application/x-apple-diskimage
date: Mon, 08 Apr 2019 16:14:42 GMT
etag: "2c1210c6567c9278985062f9dac53e57-10"
last-modified: Wed, 27 Feb 2019 11:53:53 GMT
server: AmazonS3
status: 206
via: 1.1 3245d45aedf7a8621aabe6b30d2f5a48.cloudfront.net (CloudFront)
x-amz-cf-id: EbYLgTsiNOC2U1GOqYMozH0eGMwkj94dNuD8asrjiZrE2wzOIFwd4w==
x-cache: Miss from cloudfront
```
请问：如果该文件的 HTTP 下载服务器对每个请求的下载限速为 10 KB/s，有什么办法让下载速度达到 100 KB/s？（请给出伪码以及必要的说明）。

¶ 第三题：
请实现一个简单的 Linux Shell 脚本，用于监控机器上的某个进程，如果进程退出，则需要重新拉起该进程，并输出新的进程 ID。

进程启动命令：
```bash
/usr/local/bin/ppio-demo start
```
¶ 第四题：

现在有一个需求，需要你去爬取微博用户的详细数据，其中包括用户的 nickname（昵称，具有唯一性）、follow_num（关注数）、follower_num（粉丝数）、weibo_num（微博数），以及每个用户发布的所有微博信息，其中包括微博的 content（内容，只需考虑文本内容）、post_time（发布时间）、repost_num（转发数）、comment_num（评论数）、like_num（点赞数）。爬取的数据需要通过 MySQL 数据库进行存储。

请你设计出用于存储以上数据所需要的建立的 MySQL 数据库表结构及表之间的关联关系。（用 SQL 语句表示）
在已爬取的数据中，如果要统计用户（昵称：PPlabs2019）历史微博的总点赞数，请你写出对应的 SQL 查询语句。
在已爬取的数据中，需要按用户历史微博的总点赞数进行倒序排序，选出前 50 名用户，并显示其昵称、关注数、粉丝数。请你写出对应的 SQL 查询语句。
当数据量较大时，请合理地建立索引来优化你在第 2、3 个问题中给出的 SQL 语句查询效率。
提示： 这个属于开放性试题，没有严格的答案。下面给出几个注意点。

命名规范。
表的各字段的数据类型合理。
SQL 语句没有明显的语法错误。
bs笔试题目.html



