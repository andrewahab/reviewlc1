package main

import "fmt"

func main() {
	var number int

	fmt.Println("Please enter a number:")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	if number%2 == 0 {
		fmt.Println("The number is EVEN.")
	} else {
		fmt.Println("The number is ODD.")
	}

	fmt.Println("Even numbers upto your input are:")
	for i := 1; i < number; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	if isPrime(number) {
		fmt.Println("The number is a prime number.")
	} else {
		fmt.Println("The number is not a prime number.")
	}

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}