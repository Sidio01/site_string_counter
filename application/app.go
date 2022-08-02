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
	sites         []string
)

func Start(ctx context.Context, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	if len(os.Args) < 2 {
		fmt.Println("error: provide a valid comma-separated list of urls  like \"url1, url2, url3\" as command line argument")
		return
	}
	arg := os.Args[1]

	sites = strings.Split(arg, ",")
	for idx, site := range sites {
		sites[idx] = strings.TrimSpace(site)
	}

	fmt.Println("processing...")
	for _, site := range sites {
		wg.Add(1)
		go pkg.Worker(ctx, wg, result, rateLimitChan, site)
	}
}

func Stop(ctx context.Context) {
	wg.Wait()
	// в порядке указания в параметре команды
	for _, site := range sites {
		count, _ := result.Load(site)
		fmt.Printf("%s: %v\n", site, count)
	}

	// в случайном порядке
	// result.Range(func(key, value interface{}) bool {
	// 	fmt.Printf("%s: %v\n", key, value)
	// 	return true
	// })
}
