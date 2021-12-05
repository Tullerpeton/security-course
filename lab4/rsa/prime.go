package rsa

import (
	"crypto/rand"
	"math/big"
)

func isPrime(n *big.Int) bool {
	return n.ProbablyPrime(MillerRabinN)
}

func generatePrimeNumber(min, max *big.Int) (n *big.Int) {
	if min == nil {
		min = big.NewInt(Min)
	}
	if max == nil {
		max = big.NewInt(Max)
	}
	check := false
	for check != true {
		n, _ = rand.Int(rand.Reader, max)
		if n.Cmp(min) > 0 {
			check = isPrime(n)
		}
	}
	return
}
