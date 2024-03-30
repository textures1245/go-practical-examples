package resource

import (
	"fmt"
	"os"
)

// -  Write a program that reads a file and prints its contents to the console.

type FileRead struct {
	FilePath string
	Len      int
}

func (file *FileRead) Reader() {
	if file.Len > 0 {
		f, err := os.Open("/tmp/dat")
		check(err)

		b := make([]byte, file.Len)
		n1, err := f.Read(b)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b))

	} else {
		data, err := os.ReadFile(file.FilePath)
		check(err)
		fmt.Println(string(data))

	}

}

func check(e error) {
	if e != nil {
		panic((e))
	}
}

// - Implement a function that takes a slice of integers and returns the sum of all even numbers.
type EvenNum struct {
	Nums []int
}

func (e *EvenNum) SumEven() int {
	sum := 0
	for _, num := range e.Nums {
		if num&2 == 0 {
			sum += num
		}
	}
	return sum
}

// - Create a program that simulates a simple bank account with deposit and withdrawal functions, ensuring thread safety using mutexes or channels.
type BackAccount struct {
	Balance int
}

func (b *BackAccount) Deposit(amount int) {
	b.Balance += amount
}

func (b *BackAccount) Withdraw(amount int) {
	if (b.Balance - amount) < 0 {
		panic("insufficient funds")
	}

	b.Balance -= amount
}
