package pkg

import (
	"context"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func Worker(ctx context.Context, wg *sync.WaitGroup, m *sync.Map, rateLimitChan chan struct{}, site string) {
	defer wg.Done()
	counter := 0

	htmlPage, status, err := GetHtmlPage(site)
	if err != nil {
		m.Store(site, status)
		return
	}

	token := html.NewTokenizer(strings.NewReader(htmlPage))

	for {
		select {
		case <-ctx.Done():
			m.Store(site, "cancel")
			return
		case rateLimitChan <- struct{}{}:
			for {
				select {
				case <-ctx.Done():
					m.Store(site, "cancel")
					return
				default:
					// задержка для возможности проверки корректного срабатывания прерывания программы
					time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
					if token.Next() == html.ErrorToken {
						m.Store(site, counter)
						<-rateLimitChan
						return
					} else {
						counter++
					}
				}
			}
		}
	}
}
