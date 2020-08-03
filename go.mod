module sniper-api

go 1.14

require (
	github.com/bilibili/twirp v0.0.0-20200513112506-b854eb103b5b
	github.com/dave/dst v0.25.5
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/spf13/cobra v0.0.5
	google.golang.org/protobuf v1.24.0 // indirect
	kingstar-go/sniper v0.0.0
)

replace kingstar-go/sniper => ./../kingstar-go/sniper
