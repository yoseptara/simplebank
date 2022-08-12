package util

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomWithExclusion(min, max int64, excludedNumbers ...int64) int64 {
	num := min + rand.Int63n(max-min+1-int64(len(excludedNumbers)))
	for _, ex := range excludedNumbers {
		if num < ex {
			break
		}
		num++
	}
	return num
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func MinMax(values []int64) (min int64, max int64) {
	if len(values) == 0 {
		log.Fatal("Cannot detect a minimum value in an empty slice")
	}
	min, max = values[0], values[len(values)-1]
	for _, e := range values {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	return
}

func RandomAccountId(accountIds []int64) int64 {
	min, max := MinMax(accountIds)
	randomId := RandomInt(min, max)
	return randomId
}

func RandomAccountIdWithExclusion(accountIds []int64, excludedId int64) int64 {
	min, max := MinMax(accountIds)
	randomId := RandomWithExclusion(min, max, excludedId)
	return randomId
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
