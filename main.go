package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/***
Given a positive integer, get the list of all the previous prime numbers
***/
func main() {
	var n = 1571

	var inputErr error
	n, inputErr = askForConsoleInpunt()
	if inputErr != nil {
		fmt.Println("Error reading input:", inputErr)
		return
	}

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

func askForConsoleInpunt() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Provide a positive integer")

	text, rerr := reader.ReadString('\n')
	if rerr != nil {
		return 0, rerr
	}

	text = strings.TrimSpace(text)

	n, parseErr := strconv.Atoi(text)
	if parseErr != nil {
		return 0, parseErr
	}
	return n, nil
}
