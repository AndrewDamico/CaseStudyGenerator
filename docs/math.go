package scenariogenerator

import (
	"math/rand"
	"time"
)

// generateRandomNumber generates a random number between min and max value, based on the current time.
func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())  // Initialize random seed based on the current time
	return rand.Intn(max-min+1) + min // Generate a random number within the specified range
}
