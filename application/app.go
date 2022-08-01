package application

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Sidio01/site_string_counter/pkg"
)

var (
	wg            = new(sync.WaitGroup)
	result        = new(sync.Map)
	rateLimitChan = make(chan struct{}, 4)
)

func Start(ctx context.Context) {
	arg := os.Args[1] // TODO: проверка на корректность переданных параметров
	sites := strings.Split(arg, ", ")

	fmt.Println("processing...")
	for _, site := range sites {
		wg.Add(1)
		result.Store(site, "cancel")
		go pkg.Worker(ctx, wg, result, rateLimitChan, site)
	}

}

func Stop(ctx context.Context) {
	wg.Wait()
	result.Range(func(key, value interface{}) bool { // TODO: возможно стоит сразу выводить данные отработавшего worker, а не ждать завершения всех
		fmt.Printf("%s: %v\n", key, value)
		return true
	})
}
