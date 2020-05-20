default: rpc
	go build

clean:
	git clean -x -f -d

rename:
	go run cmd/sniper/main.go rename  --package $(name)

rpc:
	find rpc -name '*.proto' -exec protoc --twirp_out=. --go_out=. {} \;

doc:
	find rpc -name '*.proto' -exec protoc --markdown_out=. --go_out=. {} \;

run-public:
	go mod vendor;
	export APP_ID=tuodianapi;	export DEPLOY_ENV=uat;	go run main.go server --port=8080;

run-private:
	go mod vendor;
	export APP_ID=tuodianapi; go run main.go server --port=8080 --internal;

run-job:
	export APP_ID=CruiseJob; go run main.go job --port=8081;

.PHONY: test rpc
