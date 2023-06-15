package main

import (
	"fmt"
	"math"
)

func isPrime(n uint64) bool {
	for _, p := range createPrimes(n) {
		if n%p == 0 {
			return false
		}
	}
	return true
}

func createPrimes(n uint64) []uint64 {
	var primes []uint64
	n = uint64(math.Sqrt(float64(n)))
	for i := uint64(2); i < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func Test() {
	fmt.Println(isPrime(7427323))     // true
	fmt.Println(isPrime(34781637149)) // false
	fmt.Println(isPrime(153259))      // true
}

func main() {
	Test()
}
