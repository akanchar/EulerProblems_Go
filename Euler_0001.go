/// <summary>
/// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
/// The sum of these multiples is 23.
/// Find the sum of all the multiples of 3 or 5 below 1000.
/// </summary>
package main
import "fmt"
// Global variables
const  (FinalNumber = 1000
		FirstFactor = 3
		SecondFactor = 5)

func main() {
	var (
	sum = 0
	i = 0)
	for ; i < FinalNumber; i++ {
		if (i%FirstFactor == 0 && i%SecondFactor == 0) {
			sum += i
		}
	}
	fmt.Println("The sum of all the natural numbers below", FinalNumber, "that are multiples of", FirstFactor, "and", SecondFactor, "is", sum)
}