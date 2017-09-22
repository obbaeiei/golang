package main

import "net/http"

func main() {
	for i := 0; i < 1000000; i++ {
		http.Get("http://localhost:1323/")
	}
}

//lsof -i -P -n | grep 1323 | wc -l
