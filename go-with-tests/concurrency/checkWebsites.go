package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// Initial function
// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		results[url] = wc(url)
// 	}

// 	return results
// }

// Concurrent function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChanel := make(chan result)

	for _, url := range urls {
		go func() {
			// results[url] = wc(url)
			resultChanel <- result{url, wc(url)} // send statement
		}()
	}

	// time.Sleep(2 * time.Second)
	for i := 0; i < len(urls); i++ {
		r := <-resultChanel // receive expression
		results[r.string] = r.bool
	}

	return results
}
