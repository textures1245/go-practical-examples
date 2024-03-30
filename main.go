package main

import (
	"fmt"

	"github.com/textures1245/practical-examples/examples/resource"
)

func main() {
	resourceTask()
}

func resourceTask() {
	// -  Write a program that reads a file and prints its contents to the console.
	file := resource.FileRead{FilePath: "/tmp/dat", Len: 0}
	file.Reader()

	// - Implement a function that takes a slice of integers and returns the sum of all even numbers.
	evenNum := resource.EvenNum{Nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	sum := evenNum.SumEven()
	println(sum)

	// - Create a program that simulates a simple bank account with deposit and withdrawal functions, ensuring thread safety using mutexes or channels.
	Ba := resource.BackAccount{Balance: 0}
	Ba.Deposit(500)
	Ba.Withdraw(400)
	fmt.Println(Ba.Balance)
}
