# go-zero 管理消息队列



## serviceGroup

在 `go-zero` 中使用 `serviceGroup` 管理消息队列

```go
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()

	fmt.Printf("Starting mq server at %s...\n", c.ListenOn)

}
```



要使用 **serviceGroup** 管理自己的服务，需要实现 `start` 和 `stop` 两个接口

> [进程内优雅管理多个服务](https://talkgo.org/t/topic/3720)

```go
type Service interface {
    Starter
    Stopper
}
// Service is the interface that groups Start and Stop methods.

func (service.Starter) Start()
func (service.Stopper) Stop()
```



## asynq

quick start

```go
package asynqredis

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hibiken/asynq"
)

type ExamplePlayload struct {
	ID   int
	Name string
}

// Client
func TestClient(t *testing.T) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	b, err := json.Marshal(ExamplePlayload{ID: 1, Name: "hello"})
	if err != nil {
		t.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(asynq.NewTask("example:client", b))
	if err != nil {
		t.Fatalf("could not enqueue task: %v", err)
	}

	t.Log(info.Payload)
}

// Server
func TestServer(t *testing.T) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Concurrency: 10},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc("example", HandlerExampleFunc)

	if err := srv.Run(mux); err != nil {
		t.Fatalf("could not run server: %v", err)
	}
}

func HandlerExampleFunc(ctx context.Context, t *asynq.Task) error {
	var p ExamplePlayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	fmt.Printf("ID: %d, Name: %s\n", p.ID, p.Name)
	return nil
}

```

