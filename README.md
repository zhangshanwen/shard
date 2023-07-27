### 基于gin后台权限管理系统

TODO 集成一些小工具库
1.cookie 字符串在线转 json(纯前端功能)
2.crontab 在线校验生成等...
3.图片裁剪(上传图片,返回裁剪后的图片)
4.微信登录,聊天.....
5.可添加hook
....
#### 功能

    1.用户管理
    2.管理员管理
    3.角色管理
    4.后端路由查看
    5.按钮级权限管理
    6.侧边栏动态管理
    7.操作日志管理
    8.主机管理(支持web终端连接)
    9.任务管理(可定时向主机发送操作指令)
    10.文件管理(支持在线编写,并可运行shell/python文件)

***
数据库: mysql + redis 创建mysql database ```create database shard charset=utf8mb4```
导入sql/shard.sql数据
***

1. 生成公私钥

```
生成私钥
openssl genrsa -out rsa/shard.rsa
根据私钥生成公钥
openssl rsa -in rsa/shard.rsa -pubout > rsa/shard.rsa.pub
```

2.目录树

```
├── Makefile
├── README.md
├── admin_api                                
├── api                      
├── bin                                         
├── cmd                                       
├── code                                       
├── common                                   
├── env                                        
├── initialize
├── internal
├── middleware
├── model
├── router
├── rsa
├── sql
└── tools

```

3.运行

```shell script
make run 
```

4.编译

```shell script
make build
```