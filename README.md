# API

## 系统要求

Sniper 仅支持 UNIX 环境。Windows 用户需要在 WSL 下使用。

环境准备好之后，需要安装以下工具的最新版本：

- go
- git
- make
- [protoc](https://github.com/google/protobuf)

## 快速入门

安装 sniper 脚手架：

```bash
go install github.com/learninto/sniper-api/cmd/sniper@latest
go install github.com/learninto/sniper-api/cmd/sniper@feat-cli
```

创建一个新项目：

```bash
sniper new --pkg helloworld
```

切换到 helloworld 目录。

运行服务：

```bash
CONF_PATH=`pwd` go run main.go http
```

使用 [httpie](https://httpie.io) 调用示例接口：

```bash
http :8080/api/foo.v1.Bar/Echo msg=hello
```

应该会收到如下响应内容：

```
HTTP/1.1 200 OK
Content-Length: 15
Content-Type: application/json
Date: Thu, 14 Oct 2021 09:49:16 GMT
X-Trace-Id: 08c408b0a4cd12c0

{
    "msg": "hello"
}
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
├── svc         # 业务逻辑层
└── sniper.toml # 配置文件
```

## 批量修改文件内容
grep -rl "检索内容" --include="*" ./ | xargs sed -i "" "s/检索内容/修改后内容/g"
