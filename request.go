package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	requestUrl        = flag.String("url", "http://127.0.0.1", "server url")
	requestKFrequency = flag.Int("count", 10, "request k frequency")
)

func main() {
	flag.Parse()
	w := sync.WaitGroup{}
	for k := 0; k < *requestKFrequency; k++ {
		go func() {
			for i := 0; i < 1000; i++ {
				get(*requestUrl)
			}
			w.Done()
		}()
		w.Add(1)
	}
	w.Wait()
}

func get(url string) error {
	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println("err:", err)
		return err
	} else {
		var buffer [512]byte
		result := bytes.NewBuffer(nil)
		for {
			n, err := res.Body.Read(buffer[0:])
			result.Write(buffer[0:n])
			if err != nil && err == io.EOF {
				break
			} else if err != nil {
				return err
			}
		}
		fmt.Println("res:", result.String())
	}
	return nil
}
