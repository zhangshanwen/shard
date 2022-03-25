### 基于gin后台权限管理系统

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