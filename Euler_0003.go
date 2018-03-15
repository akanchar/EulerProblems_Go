/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
 */
package main

import "fmt"

const Number = 600851475143
func main() {
	largestPrime := 0
	for i := 2; i < Number/2; i++ {
		isPrime := true;
		if(Number % i == 0) {
			for j := 2; j < i/2; j++ {
				if(i % j == 0) {
					isPrime = false
					break
				}
			}
			if(isPrime) {
				largestPrime = i;
			}
		}
	}
	fmt.Println("The largest prime factor of",Number,"is",largestPrime,".")
}