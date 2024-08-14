package main

import (
	"fmt"
	"sync"
)

/*
This solution uses WaitGroup to force each gorountines to wait for its child gorountines to exit.
*/

type SafeCounter struct {
	v   map[string]bool
	mux sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var cnt SafeCounter = SafeCounter{v: make(map[string]bool)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.

	if depth <= 0 {
		wg.Done()
		return
	}

	cnt.mux.Lock()
	_, ok := cnt.v[url]
	if ok == false {
		cnt.v[url] = true
		cnt.mux.Unlock()
	} else {
		wg.Done()
		cnt.mux.Unlock()
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	var wg_ sync.WaitGroup
	for _, u := range urls {
		wg_.Add(1)
		go Crawl(u, depth-1, fetcher, &wg_)
	}

	wg_.Wait()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	go Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Add(1)
	wg.Wait()
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
