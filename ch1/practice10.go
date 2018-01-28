// fetchall
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go twice(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func twice(url string, ch chan string) {
	secs1, nbytes1, err := fetch(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs2, nbytes2, err := fetch(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%s fetch1: %.2f %7d, fetch2: %.2f %7d", url, secs1, nbytes1, secs2, nbytes2)
}

func fetch(url string) (float64, int64, error){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return 0.0, 0, err
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0.0, 0, err
	}
	secs := time.Since(start).Seconds()
	return secs, nbytes, nil
}
