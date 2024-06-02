package goroutinedemo

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func ExecuteWaitGroup() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		wg.Add(1)
		go checkUrl(link)
	}
	wg.Wait()

}

func checkUrl(link string) {
	defer wg.Done()
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")

}
