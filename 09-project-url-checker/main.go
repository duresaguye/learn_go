package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string, c chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		c <- fmt.Sprintf("[DOWN] %s (Error: %s)", url, err)
		return
	}

	c <- fmt.Sprintf("[UP]   %s (Status: %s, Time: %s)", url, resp.Status, elapsed)
}

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	fmt.Println("Checking URLs concurrently...")
	start := time.Now()

	// Launch a goroutine for each URL
	for _, url := range urls {
		go checkUrl(url, c)
	}

	// Wait for all results
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

	fmt.Printf("\nTotal time taken: %s\n", time.Since(start))
}
