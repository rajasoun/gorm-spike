package load

import (
	"math/rand"
	"runtime"
	"time"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	// Create a channel to receive random characters
	charChan := make(chan byte, length)

	// Create multiple goroutines to generate random characters
	numGoroutines := runtime.NumCPU() // Use the number of available CPUs
	for i := 0; i < numGoroutines; i++ {
		go func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for j := 0; j < length/numGoroutines; j++ {
				charChan <- charset[r.Intn(len(charset))]
			}
		}()
	}

	// Collect the random characters from the channel
	for i := 0; i < length; i++ {
		result[i] = <-charChan
	}

	return string(result)
}
