package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type History struct {
	history map[string]bool
	mux     sync.Mutex
}

func (h *History) isFetched(url string) bool {
	h.mux.Lock()
	defer h.mux.Unlock()

	_, ok := h.history[url]
	if ok {
		return ok
	} else {
		fmt.Printf("h.history : %v\n", h.history)
		fmt.Printf("map「h.history」に%vいれる\n", url)
		h.history[url] = true
	}
	return false
}

var history = History{history: make(map[string]bool)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	if depth <= 0 || history.isFetched(url) {
		fmt.Printf("if depth <= 0 || history.isFetched(url) :%v or %v\n +++++++++++ \n", depth, url)
		return
	}

	// fetcherのなかに該当するmapの値があるか
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	fmt.Println("!!!!!!!!")

	for _, u := range urls {
		//goroutineを起動する前にwgをインクリメント
		wg.Add(1)

		go Crawl(u, depth-1, fetcher, wg)
	}

	return
}

func syncCrawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
}

func main() {
	syncCrawl("https://golang.org/", 4, fetcher)
	//time.Sleep(5*time.Second)
	fmt.Println("done")
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// fetcherのなかに該当するmapの値があるか
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		fmt.Printf("fetcherのなかに該当するmapの値%vがある\n", url)
		fmt.Printf("%v, %v\n", res.body, res.urls)
		return res.body, res.urls, nil
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>")
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