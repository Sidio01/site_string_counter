package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/Sidio01/site_string_counter/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	go application.Start(ctx)
	<-ctx.Done()
	application.Stop(ctx)
}

// go run cmd/main.go "ya.ru, google.com, mts.ru, stepik.org, mail.ru, dfhdfhrtdhrth.com"
