package pkg

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func Worker(ctx context.Context, wg *sync.WaitGroup, m *sync.Map, rateLimitChan chan struct{}, site string) {
	defer wg.Done()
	counter := 0

	htmlPage, err := GetHtmlPage(site)
	if err != nil {
		m.Store(site, err) // TODO: получать statusCode
		return
	}
	token := html.NewTokenizer(strings.NewReader(htmlPage))
	for {
		select {
		case <-ctx.Done():
			log.Println("worker stopping work")
			if counter > 0 {
				m.Store(site, counter)
			}
			return
		default:
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			if token.Next() == html.ErrorToken {
				m.Store(site, counter)
				return
			} else {
				counter++
			}
		}
	}
}
