package main

import (
	"math/rand"
	"time"
	"fmt"
	"net/http"
)

var r *rand.Rand // Rand for this package.

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString - Generate a random string
func RandomString(strlen int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "[FLAG]"+RandomString(5)+"-"+RandomString(5)+"-"+RandomString(5)+"-"+RandomString(5), r.URL.Path[1:])
}


func main() {
	fmt.Println("[FLAG]"+RandomString(5)+"-"+RandomString(5)+"-"+RandomString(5)+"-"+RandomString(5))
}