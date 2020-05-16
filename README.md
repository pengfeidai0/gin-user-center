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

## 配置文件

```yaml
server:
  port: 8000
  mode: 'release'
  limitNum: 20
  #是否使用mongo
  useMongo: false
  # 是否使用redis
  useRedis: false

# redis配置文件
redis:
  Addr: '127.0.0.1:6379'
  password: ''
  db: 0

# mysql配置文件
mysql:
  user: ''
  password: ''
  path: '127.0.0.1:3306'
  database: 'user-center'
  config: 'charset=utf8&parseTime=True&loc=Local'
  driver: 'mysql'
  maxIdleConns: 10
  maxOpenConns: 100
  log: false

mongo:
  database: ''
  url: ''

session:
  key: ''
  size: 10
  # 7 * 86400 7天
  maxAge: 604800
  path: '/'
  domain: ''
  httpOnly: true

# 日志文件
log:
  debug: true
  maxAge: 7
  fileName: 'server.log'
  dirName: '/opt/data/gin-user-center/logs'

file:
  # dirName: '/opt/data/gin-user-center/file/'
  dirName: '/Users/zl/workspace/go/gin-user-center/public/file/'
  urlPrefix: 'http://127.0.0.1:8000/api/v1/gin-user-center/file/'  

url:
  # 路由前缀
  prefix: '/api/v1/gin-user-center'

oss:
  endpoint: ''
  accessKeyId: ''
  accessKeySecret: ''
  bucket: ''
```

### 启动容器

```shell
# docker build
make docker

# 启动容器（线上使用）
docker run --name gin-user-center --network=host \
  -v /opt/conf/gin-user-center:/opt/conf \
  -v /data/gin-user-center:/data/gin-user-center \
  -d gin-user-center
```