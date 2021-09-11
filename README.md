# API

## 系统要求

1. 类 UNIX 系统
2. go v1.12+
3. [protoc](https://github.com/google/protobuf)
4. [protoc-gen-go](https://github.com/golang/protobuf/tree/master/protoc-gen-go)

## 初始化环境

```shell script
brew install protobuf
go mod download
go mod vendor
go get -u github.com/golang/protobuf/protoc-gen-go
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/learninto/sniper-api/cmd/protoc-gen-twirp
```

## 文档生成

```
第一步：安装文档生成器：
go get -u github.com/learninto/protoc-gen-markdown

第二步：在项目根目录下执行：
find rpc -name '*.proto' -exec protoc --markdown_out=. --go_out=. {} \;
```

## 目录结构

```
├── cmd         # 服务子命令
├── dao         # 数据访问层
├── main.go     # 项目总入口
├── rpc         # 接口描述文件
├── server      # 控制器层
├── service     # 业务逻辑层
├── conf.toml # 配置文件
└── util        # 业务工具库
```

## 快速入门

- [定义接口](./rpc/README.md)
- [实现接口](./rpc/README.md)
- [注册服务](./cmd/server/README.md)
- [启动服务](./cmd/server/README.md)
- [配置文件](https://github.com/learninto/goutil/conf/README.md)
- [日志系统](https://github.com/learninto/goutil/log/README.md)
- [数据库](https://github.com/learninto/goutil/db/README.md)
- [redis](https://github.com/learninto/goutil/redis/README.md)
- [指标监控](https://github.com/learninto/goutil/metrics/README.md)
- [链路追踪](https://github.com/learninto/goutil/trace/README.md)
- [高性能JSON处理](https://github.com/learninto/goutil/blob/master/xjson/json_test.go)

## 框架依赖项

- [框架依赖](./go.mod)

## 批量修改文件内容
grep -rl "检索内容" --include="*" ./ | xargs sed -i "" "s/检索内容/修改后内容/g"
