// lab1
package main

import (
	"encoding/json"
	"net/http"
)

type countAdd struct {
	Count int `json:"count"`
}

var n = 0
var c = make(chan int, 10)

func f() {
	<-c //确保n++正确执行
	n++
}

func count_add(w http.ResponseWriter, req *http.Request) {
	c <- 0
	go f()
	output := countAdd{n}
	var outputJson []byte
	outputJson, _ = json.Marshal(output)
	w.Write(outputJson)
}

func main() {
	http.HandleFunc("/add", count_add)
	http.ListenAndServe(":9090", nil)
}
