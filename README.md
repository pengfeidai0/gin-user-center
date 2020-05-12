# gin-user-center

基于 [gin-user-start](https://github.com/pengfeidai/gin-user-start) 完成的同一的应用层用户中心，包含用户登录、会话校验、角色分配、权限管理等功能。

## 功能点
- [x] 用户管理
  - [x] 添加用户

## 快速开始

#### 代码仓库

```go
git clone git@github.com:pengfeidai/gin-user-center.git
```

#### 环境配置

- Go version >= 1.13
- Global environment configure

```go
export GO111MODULE=on
// 修改代理
export GOPROXY=https://goproxy.io
// go env -w GOPROXY=https://goproxy.cn,direct 
```

#### 服务启动

```go
cd gin-user-center

go run main.go

输出如下 `Listening and serving HTTP on Port: :8000, Pid: 15932`，表示 Http Server 启动成功。
```

#### 健康检查

```
curl -X GET http://127.0.0.1:8000/health_check?name=world
```

## 文档