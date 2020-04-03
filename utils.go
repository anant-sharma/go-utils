package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
)

// Sha256 to generate sha256
func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateUUID - Function to generate UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateShortID - Function to Generate Short UUID
func GenerateShortID() string {
	return shortuuid.New()
}

// SetInterval - SetInterval Function
func SetInterval(someFunc func(), milliseconds int, async bool) chan bool {

	// How often to fire the passed in function
	// in milliseconds
	interval := time.Duration(milliseconds) * time.Millisecond

	// Setup the ticket and the channel to signal
	// the ending of the interval
	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	// Put the selection in a go routine
	// so that the for loop is none blocking
	go func() {
		for {

			select {
			case <-ticker.C:
				if async {
					// This won't block
					go someFunc()
				} else {
					// This will block
					someFunc()
				}
			case <-clear:
				ticker.Stop()
				return
			}

		}
	}()

	// We return the channel so we can pass in
	// a value to it to clear the interval
	return clear
}
