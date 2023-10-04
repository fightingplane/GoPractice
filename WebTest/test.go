package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func get(url string) ([]byte, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bytes, read_err := io.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(bytes)

	return bytes, read_err
}

func main() {
	for i := 0; i < 1040; i++ {
		go get(fmt.Sprintf("http://www.httpbin.org/get?a=%d", i))
	}

	time.Sleep(1 * time.Second)
}
