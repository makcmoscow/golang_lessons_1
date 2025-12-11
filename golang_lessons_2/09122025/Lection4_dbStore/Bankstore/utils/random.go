package utils

import (
	"github.com/go-faker/faker/v4"
	"time"
	"math/rand"
	"log"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

type RandomAccountParams struct {
	Owner string `faker:"last_name"`
	Balance int64
	Currency string `faker:"oneof: USD, EUR"`
}

func RandomAccount() RandomAccountParams {
	rap := RandomAccountParams {}
	err := faker.FakeData(&rap)
	if err != nil {
		log.Fatal(err)
	}
return rap
}