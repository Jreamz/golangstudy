package main

import "fmt"

func printPrimes(max int) {
	for num := 2; num <= max; num++ {
		if num == 2 {
			fmt.Println(num)
			continue
		}

		if num%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i*i <= num; i += 2 {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(num)
		}
	}
}

// don't edit below this line

func main() {
	printPrimes(10)
}
