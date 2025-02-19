package prime

// GetPrime returns a slice of prime numbers up to the given limit
func GetPrime(number int) []int {
	if number < 2 {
		return nil
	}

	var primes []int
	isPrime := make([]bool, number)

	// Initialize array
	for i := 2; i < number; i++ {
		isPrime[i] = true
	}

	// Sieve of Eratosthenes
	for i := 2; i*i < number; i++ {
		if isPrime[i] {
			for j := i * i; j < number; j += i {
				isPrime[j] = false
			}
		}
	}

	// Collect prime numbers
	for i := 2; i < number; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
