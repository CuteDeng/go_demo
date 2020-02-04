package main

import (
	"fmt"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/test", f1)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Println("http server listen err:", err)
		return
	}
}
