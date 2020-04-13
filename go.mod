module sniper-api

go 1.14

require (
	github.com/bilibili/twirp v0.0.0-20200305140827-a09be7e42ab8
	github.com/golang/protobuf v1.3.5
	github.com/spf13/cobra v0.0.5
	kingstar-go/sniper v0.0.0
)

replace kingstar-go/sniper => ./../kingstar-go/sniper
