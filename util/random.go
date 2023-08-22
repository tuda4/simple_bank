package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxy"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// random Int generates a random integer from min to max

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// random String generates a random string of length n
func RandString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandObj generates a random owner name
func RandOwner() string {
	return RandString(6)
}

// randMoney generates a random amount of money
func RandMoney() int64 {
	return randInt(0, 1000)
}

func RandCurrency() string {
	currencies := []string{"USD", "VND", "EUR", "GBP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandEmail() string {
	return fmt.Sprintf("%s@email.com", RandString(6))
}
