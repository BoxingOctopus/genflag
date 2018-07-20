package main

import (
	"math/rand"
	"time"
	"fmt"
	"flag"
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


func main() {
	flagPrefix := flag.String("p", "[FLAG]", "Prefix for Flag")
	flagType := flag.String("t", "single", "Type of Flag ('single' or 'segmented')")
	segDelim := flag.String("d", "-", "If flag type is segmented, what should the delimiter be?")
	segLen := flag.Int("l", 4, "If Flag Type is segmented, how long should each segment be?")
	flagLen := flag.Int("n", 12, "Flag Suffix Length if type is 'single'")
	numFlags := flag.Int("f", 1, "Number of Flags to generate")
	flag.Parse()

	for i := 0; i < *numFlags; i++ {
		if *flagType == "segmented" {
			fmt.Println(*flagPrefix+RandomString(*segLen)+*segDelim+RandomString(*segLen)+*segDelim+RandomString(*segLen)+*segDelim+RandomString(*segLen))
		} else {
			fmt.Println(*flagPrefix+RandomString(*flagLen))
		}
	}
	
}