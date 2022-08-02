package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Sidio01/site_string_counter/application"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go application.Start(ctx, wg)
	wg.Wait()

	application.Stop(ctx)
}
