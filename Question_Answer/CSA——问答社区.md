# CSA——问答社区

## 主要实现功能

1.实现用户的注册和登录

2.用户可以发表自己的问题也可以对问题进行回答

3.用户可以修改、删除、查看自己的问题和回答

## 中间件

运用了cookie和cors

### cookie

cookie 存储在客户端： cookie 是服务器发送到用户浏览器并保存在本地的一小块数据，它会在浏览器下次向同一服务器再发起请求时被携带并发送到服务器上。因此，服务端脚本就可以读、写存储在客户端的cookie的值。

cookie用于获取登录状态，确保用户在进行其他操作的时候，必须处于登录状态。避免资源的浪费。



### cors

CORS 是一种基于 [HTTP Header](https://link.juejin.cn?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FGlossary%2FHeader) 的机制，该机制通过允许服务器标示除了它自己以外的其它域。服务器端配合浏览器实现 `CORS` 机制，可以突破浏览器对跨域资源访问的限制，实现跨域资源请求。

采用cors中间件，解决跨域问题

## Gin框架

### 什么是gin框架

Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

### 框架结构

README.md：项目的说明文档

api：接口层，在里面是详细的逻辑实现以及路由。

dao：全名为 data access object，操作数据库。

service:调用dao层的一些函数从而实现api层的一些功能

model：模型层，主要放数据库实例的结构体。

tools：一些常用的工具函数，封装在这里减少代码的重复使用。

middleware:存放一些中间件

cmd:项目的入口，存放main函数

go.mod：依赖管理

![image-20221220152434971](https://github.com/TOMORROWLHY/CSA-work/blob/main/Question_Answer/photos/1.png)

## Mysql——数据库

### MySQL介绍

该项目运用MySQL作为数据库，进行数据储存和数据反馈

MySQL是一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统关系型数据库管理系统之一，在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一

### 数据表结构

使用了三个数据表，进行数据关联

用csaqa表示该项目数据库

![image-20221220153712995](C:\Users\李\AppData\Roaming\Typora\typora-user-images\image-20221220153712995.png)

user表涉及用户ID、用户名、密码

![image-20221220153926830](C:\Users\李\AppData\Roaming\Typora\typora-user-images\image-20221220153926830.png)

question表涉及用户ID、问题ID、问题内容

![image-20221220154038612](C:\Users\李\AppData\Roaming\Typora\typora-user-images\image-20221220154038612.png)

answer表，涉及用户ID、问题ID、回答ID、回答内容

![image-20221220154121882](C:\Users\李\AppData\Roaming\Typora\typora-user-images\image-20221220154121882.png)