package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	for {
		if rand_n := rand.Float64(); rand_n < 0.33 {
			logger.Printf("[INFO][%s] Info message, n=%.3f", time.Now().Format(time.RFC3339), rand_n)
		} else if rand_n < 0.66 {
			logger.Printf("[WARNING][%s] Warning message, n=%.3f", time.Now().Format(time.RFC3339), rand_n)
		} else {
			logger.Printf("[ERROR][%s] Error message, n=%.3f", time.Now().Format(time.RFC3339), rand_n)
		}
		time.Sleep(time.Second * 5)
	}
}
