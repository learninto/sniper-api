# API
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flearninto%2Fsniper-api.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flearninto%2Fsniper-api?ref=badge_shield)


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
- [实现接口](./server/README.md)
- [注册服务](./cmd/server/README.md)
- [启动服务](./cmd/server/README.md)
- [配置文件](util/conf/README.md)
- [日志系统](./util/log/README.md)
- [数据库](./util/db/README.md)
- [redis](./util/redis/README.md)
- [指标监控](./util/metrics/README.md)
- [链路追踪](./util/trace/README.md)

## 框架依赖项

- [框架依赖](./go.mod)

## 批量修改文件内容
grep -rl "检索内容" --include="*" ./ | xargs sed -i "" "s/检索内容/修改后内容/g"

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flearninto%2Fsniper-api.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flearninto%2Fsniper-api?ref=badge_large)