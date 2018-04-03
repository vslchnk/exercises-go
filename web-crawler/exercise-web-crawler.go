package main

import (
	"fmt"
	"sync"
	"time"
)

/*type res struct {
	url string
	body string
}*/

type SafeMap struct {
	v   map[string]string
	mux sync.Mutex
	//r chan res
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (m *SafeMap) Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	if _, ok := m.v[url]; !ok {
		m.mux.Lock()
		body, urls, err := fetcher.Fetch(url)
		m.mux.Unlock()
		m.v[url] = body
		//m.r <- res{url, body}

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go m.Crawl(u, depth-1, fetcher)
		}
	}

	return
}

func main() {
	m := SafeMap{v: make(map[string]string) /*, r: make(chan res)*/}
	//Crawl("https://golang.org/", 4, fetcher)
	for i := 0; i < 10; i++ {
		go m.Crawl("https://golang.org/", 4, fetcher)
	}

	time.Sleep(time.Second)
	/*for k := range m.v {
	    fmt.Println("key:", k)
	}*/
	/*for resp := range m.r {
		fmt.Printf("found: %s %q\n", resp.url, resp.body)
	}*/
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
