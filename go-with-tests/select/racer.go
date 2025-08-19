package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Initial implementation
// func Racer(url1, url2 string) string {
// 	start1 := time.Now()
// 	http.Get(url1)
// 	duration1 := time.Since(start1)

// 	start2 := time.Now()
// 	http.Get(url2)
// 	duration2 := time.Since(start2)

// 	if duration1 < duration2 {
// 		return url1
// 	}
// 	return url2
// }

// ----------------------------------------------------------------------------------- //
// Refactor 1
// func Racer(url1, url2 string) string {
// 	duration1 := measureResponseTime(url1)
// 	duration2 := measureResponseTime(url2)

// 	if duration1 < duration2 {
// 		return url1
// 	}
// 	return url2
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// ----------------------------------------------------------------------------------- //
// Refactor 2
// func Racer(url1, url2 string) string {
// 	// Ref: https://go.dev/tour/concurrency/5
// 	select {
// 	case <-ping(url1):
// 		return url1
// 	case <-ping(url2):
// 		return url2
// 	}
// }

// ----------------------------------------------------------------------------------- //
// Refactor 3
// func Racer(url1, url2 string) (string, error) {
// 	select {
// 	case <-ping(url1):
// 		return url1, nil
// 	case <-ping(url2):
// 		return url2, nil
// 	case <-time.After(1 * time.Second):
// 		return "", fmt.Errorf("timeout waiting for %s and %s", url1, url2)
// 	}
// }

// ----------------------------------------------------------------------------------- //
// Refactor 4
const defaultTimeout = 10 * time.Second

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, defaultTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
