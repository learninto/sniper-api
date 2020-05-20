# cmd/job

注册定时任务请参考 [demo.go](./demo.go)。

## 查看所有定时任务
```bash
go main.go job list
```

## 执行一次某个任务
```bash
go main.go job once foo
```

## 调度所有定时任务
```bash
go main.go job
```

## 案例
```go
package job

import (
	"context"
	"fmt"
	"time"
)

// 定时任务示例，开源专用
// 业务相关任务请使用 cron.go
func init() {
	manual("foo", func(ctx context.Context) error {
		fmt.Printf("manual run foo with args: %+v\n", onceArgs)
		return nil
	})

	cron("bar", "@every 1m", func(ctx context.Context) error {
		fmt.Printf("run bar @%v\n", time.Now())
		return nil
	})

	http("baz", "0 18-23 * * *", func(ctx context.Context) error {
		fmt.Printf("run http task @%v\n", time.Now())
		return nil
	})
}
```
