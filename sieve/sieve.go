package sieve

import "time"

// return list of primes less than N
func Eratosthenes(N int) (primes []int, dur float64) {
	start := time.Now()
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	dur = float64(time.Since(start).Nanoseconds()) / 1000
	return
}
