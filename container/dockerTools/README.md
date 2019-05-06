## docker命令检索

Usage:	docker [OPTIONS] COMMAND

A self-sufficient runtime for containers

Options:

  选项：
 ```bash 
--config string客户端配置文件的位置（默认值“/Users/ThomasHuke/.docker”）

-D， -  debug启用调试模式

-H， - 主机列表要连接的守护程序套接字

-l， -  log-level string设置日志记录级别（ “调试” | “信息” | “警告” | “错误” | “致命”）（默认“信息”）

--tls使用TLS;由--tlsverify暗示

--tlscacert字符串仅由此CA签名的信任证书（默认值“/Users/[登录用户]/.docker/ca.pem”）

--tlscert string TLS证书文件的路径（默认值“/Users/xx/.docker/cert.pem”）

--tlskey string TLS密钥文件的路径（默认值“/Users/xx/.docker/key.pem”）

--tlsverify使用TLS并验证远程

-v， -  version打印版本信息并退出
```
管理命令

  - [checkpoint](#checkpoint)  管理检查站
  - config      管理Docker配置
  - container   管理容器
  - image      管理镜像
  - network     管理网络
  - node       管理Swarm节点
  - plugin     管理插件
  - secret      管理Docker的秘密
  - service     管理服务
  - stack       管理Docker堆栈
  - swarm      管理Swarm
  - system      管理Docker
  - trust       管理对Docker镜像的信任
  - volume      管理数量
  

Commands:

  - [attach](#attach)      Attach local standard input, output, and error streams to a running container
  - [build](#build)       Build an image from a Dockerfile
  - commit      Create a new image from a container's changes
  - cp          Copy files/folders between a container and the local filesystem
  - create      Create a new container
  - deploy      Deploy a new stack or update an existing stack
  - diff        Inspect changes to files or directories on a container's filesystem
  - events      Get real time events from the server
  - exec        Run a command in a running container
  - export      Export a container's filesystem as a tar archive
  - history     Show the history of an image
  - images      List images
  - import      Import the contents from a tarball to create a filesystem image
  - info        Display system-wide information
  - inspect     Return low-level information on Docker objects
  - kill        Kill one or more running containers
  - load        Load an image from a tar archive or STDIN
  - login       Log in to a Docker registry
  - logout      Log out from a Docker registry
  - logs        Fetch the logs of a container
  - pause       Pause all processes within one or more containers
  - port        List port mappings or a specific mapping for the container
  - ps          List containers
  - pull        Pull an image or a repository from a registry
  - push        Push an image or a repository to a registry
  - rename      Rename a container
  - restart     Restart one or more containers
  - rm          Remove one or more containers
  - rmi         Remove one or more images
  - run         Run a command in a new container
  - save        Save one or more images to a tar archive (streamed to STDOUT by default)
  - search      Search the Docker Hub for images
  - start       Start one or more stopped containers
  - stats       Display a live stream of container(s) resource usage statistics
  - stop        Stop one or more running containers
  - tag         Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
  - top         Display the running processes of a container
  - unpause     Unpause all processes within one or more containers
  - update      Update configuration of one or more containers
  - version     Show the Docker version information
  - wait        Block until one or more containers stop, then print their exit codes

Run 'docker COMMAND --help' for more information on a command.
