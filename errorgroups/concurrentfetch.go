package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	urls := []string{"https://google.com", "https://linkedin.com", "https://x.com"}
	g, ctx := errgroup.WithContext(context.Background())
	client := &http.Client{}

	for _, url := range urls {
		g.Go(func() error {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				return err
			}

			resp, err := client.Do(req)
			if err != nil {
				return err
			}

			defer resp.Body.Close()

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
