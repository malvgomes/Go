package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i < int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func listPrimes(n int) []int {
	primes := []int{}
	for i := 0; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func factorize(n int, primes []int) []int {
	i := 0
	factors := []int{}
	for n != 1 {
		if n%primes[i] == 0 {
			n /= primes[i]
			factors = append(factors, primes[i])
		} else {
			i++
		}
	}

	return factors
}

func countDivs(factors []int) int {
	if len(factors) == 0 {
		return 1
	}
	if len(factors) == 1 {
		return 2
	}

	i := 0
	total := 1
	count := 1
	cur := factors[0]
	for i < len(factors)-1 {
		if factors[i+1] != cur {
			total *= count + 1
			count = 1
			cur = factors[i+1]
		} else {
			count++
		}
		i++
	}

	if factors[i] != cur {
		total *= 2
	} else {
		total *= count + 1
	}

	return total
}

func main() {
	primes := listPrimes(1000000)
	j := 0
	for i := 1; ; i++ {
		j += i
		factors := factorize(j, primes)
		if countDivs(factors) >= 1000 {
			break
		}
	}

	fmt.Println(j)
}
