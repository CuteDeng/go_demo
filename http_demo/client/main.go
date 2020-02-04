package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// res, err := http.Get("http://127.0.0.1:9000/test")
	// if err != nil {
	// 	fmt.Println("http get err:", err)
	// 	return
	// }
	// b := make([]byte, 128)
	// n, err := res.Body.Read(b)
	// b, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	if err != io.EOF {
	// 		fmt.Println("res body err:", err)
	// 		return
	// 	}

	// }
	// fmt.Println(string(b[:n]))

	urlObj, _ := url.Parse("http://127.0.0.1:9000/test")
	data := url.Values{}
	data.Set("name", "张三")
	data.Set("age", "19")
	urlObj.RawPath = data.Encode()
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Println("http get err:", err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http get err:", err)
		return
	}
	defer res.Body.Close()
	b := make([]byte, 128)
	n, err := res.Body.Read(b)
	if err != nil {
		if err != io.EOF {
			fmt.Println("res body err:", err)
			return
		}
	}
	fmt.Println(string(b[:n]))
}
