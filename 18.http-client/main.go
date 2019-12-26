package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// https://golang.org/pkg/net/http/
var sum int
var mu sync.Mutex

func newHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
		},
	}
	tr := client.Transport.(*http.Transport)
	tr.MaxIdleConnsPerHost = 100
	return client
}

func newHTTPClient2() *http.Client {
	// http://tleyden.github.io/blog/2016/11/21/tuning-the-go-http-client-library-for-load-testing/
	tr, _ := http.DefaultTransport.(*http.Transport)
	tr.MaxIdleConnsPerHost = 100
	client := &http.Client{Transport: tr}
	return client
}

func getNcVersion(client *http.Client, request *http.Request) int {
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("err: %v; NumGoroutine: %d\n", err, runtime.NumGoroutine())
		return 0
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Printf("reponse code %d != 200\n", response.StatusCode)
		return 0
	}

	io.Copy(ioutil.Discard, response.Body)

	// cookie := response.Header.Get("set-cookie")
	// fmt.Println(cookie)

	version, ok := strconv.Atoi(response.Header.Get("nc-version"))
	if ok != nil {
		fmt.Printf("strconv.Atoi(%s) had failed\n", response.Header.Get("nc-version"))
		return 0
	}

	mu.Lock()
	sum++
	mu.Unlock()

	return version
}

func taskManager(wd *sync.WaitGroup) {
	url := "http://k8s.navicore.cn/"
	request, _ := http.NewRequest(http.MethodGet, url, nil)

	client := newHTTPClient2()
	referVersion := getNcVersion(client, request)
	if referVersion <= 0 {
		fmt.Println("get Nc-Version had error")
	} else {
		fmt.Println("refer version is", referVersion)
	}

	start := time.Now()
	for {
		v := getNcVersion(client, request)
		if v < referVersion {
			fmt.Printf("v(%d) < referVersion(%d)\n", v, referVersion)
		}

		if time.Now().Sub(start) > 10*time.Second {
			break
		}
	}
	fmt.Println("task finish...")
	wd.Done()
}

func main() {
	maxThread := 20
	var wg sync.WaitGroup
	for i := 0; i < maxThread; i++ {
		wg.Add(1)
		go taskManager(&wg)
	}
	wg.Wait()
	fmt.Printf("\nwait go thread finish, sum is %d\n", sum)
}
