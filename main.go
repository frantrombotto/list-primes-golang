package main

import (
	"fmt"
	"math"
)

/***
Given a positive integer, get the list of all the previous prime numbers
***/
func main() {
	var n = 1571

	primes := getPrimeNumbers(n)

	fmt.Println("Primes list:")
	fmt.Println(primes)
}

func getPrimeNumbers(n int) []int {
	var primes = make([]int, 0)

	var flags = make([]bool, n+1)

	var maxPrime = int(math.Sqrt(float64(n)))
	var prime = 2

	for prime <= maxPrime {
		discardMultiples(prime, flags)
		prime = getNextPrime(prime, flags)
	}

	for i, discarded := range flags {
		if !discarded {
			primes = append(primes, i)
		}
	}

	return primes
}

// Set flag position in true when discarded
func discardMultiples(prime int, flags []bool) {
	for i := prime * prime; i < len(flags); i += prime {
		flags[i] = true
	}
}

// Iterate numbers by one checking if flag position was not discarded
func getNextPrime(prime int, flags []bool) int {
	next := prime + 1
	for next < len(flags) && flags[next] {
		next++
	}
	return next
}
