# https://stackoverflow.com/a/12959694
rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

RPC_PROTOS := $(call rwildcard,rpc/,*.proto)
LIB_PROTOS := $(call rwildcard,util/,*.proto)

RPC_PBGENS := $(RPC_PROTOS:.proto=.twirp.go)
LIB_PBGENS := $(LIB_PROTOS:.proto=.pb.go)

.PRECIOUS: $(RPC_PBGENS) $(LIB_PBGENS)

# 参数 Mfoo.proto=bar/foo 表示 foo.proto 生成的 go 文件所对应的包名是 bar/foo。
#
# 如是在 proto 中引用了其他 proto，生成的 go 文件需要导入对应的包。
# 但 protoc 和 proto-gen-go 无法单独从 proto 文件获取当前项目的包名，
# 最好的办法就是通过 go_package 手工指定，但这样写起来太丑了，所以改用 M 参数。
#
# 如果你自己写了包供别人导入使用，则一定要在 proto 中设置 go_package 选项。
#
# 更多讨论请参考
# https://github.com/golang/protobuf/issues/1158#issuecomment-650694184
#
# $(...) 中的神奇代码是为实现以下替换
# util/kv/taishan/taishan.proto => sniper/util/taishan
%.pb.go: %.proto
	protoc --go_out=M$<=$(patsubst %/,%,$(dir $<)):. $<

# $(...) 中的神奇代码是为实现以下替换
# rpc/util/v0/kv.proto => rpc/util/v0;util_v0
%.twirp.go: %.proto
	$(eval m=$<=$(join \
			$(patsubst %/,%\;,\
				$(dir $<)\
			),\
			$(subst /v,_v,\
				$(patsubst rpc/%,%,\
					$(patsubst %/,%,$(dir $<))\
				)\
			)\
		))
	protoc --twirp_out=root_package=sniper,validate_enable=true,M$m:. \
		--go_out=M$m:. \
		$<

default: rpc util
	go build -trimpath -mod=readonly

rpc: $(RPC_PBGENS)
	@exit

util: $(LIB_PBGENS)
	@exit

cmd:
	go install ./cmd/protoc-gen-twirp

clean:
	git clean -x -f -d

rename:
	go run cmd/sniper/main.go rename  --package $(name)

doc:
	find rpc -name '*.proto' -exec protoc --markdown_out=. --go_out=. {} \;

run-public:
	export APP_ID=SniperApi;	export DEPLOY_ENV=uat;	go run main.go server --port=8080;

run-private:
	export APP_ID=SniperInternalApi; go run main.go server --port=8080 --internal;

run-job:
	export APP_ID=SniperJob; go run main.go job --port=8081;

.PHONY: clean rpc util cmd
