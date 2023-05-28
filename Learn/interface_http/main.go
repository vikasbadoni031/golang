package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.google.com")
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
