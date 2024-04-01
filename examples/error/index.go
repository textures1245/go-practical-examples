package error

import (
	"errors"
	"fmt"
	"os"
	"sync/atomic"
)

// - Write a function that reads data from a file and handles various errors that may occur during the process.
type cError struct {
	msg string
	e   error
}

// passing error to custom error
func ErrorMiddleware(err error) error {
	if err != nil {
		return &cError{e: err}
	}
	return nil
}

func (cE *cError) FileHandler(err error) {

	if err != nil {
		cE.e = err
		if errors.Is(err, os.ErrNotExist) {
			cE.msg = "File not found"
		} else if errors.Is(err, os.ErrPermission) {
			cE.msg = "Permission denied"
		} else {
			cE.msg = "Unknown error"
		}
	}
}

func (cE *cError) Error() string {
	return fmt.Sprintf("Error: %s, %v", cE.msg, cE.e)
}

type File struct {
}

func (f *File) Read(fPath string) {
	dat, err := os.ReadFile(fPath)
	// extends error with custom error, so we can handle it inside error.As logic
	fileErr := ErrorMiddleware(err)
	var cE *cError

	// handle here
	if errors.As(fileErr, &cE) {
		cE.FileHandler(err)
		fmt.Println(cE.Error())
		return
	}

	fmt.Println(string(dat))

}

// - Implement a program that uses pointers to modify the values of variables passed to a function.
type History[T any] struct {
	pos *[]T
}

type List[T any] struct {
	elems []T
	pos   []History[T]
}

func (lst *List[T]) ChangeTo(arr []T) {
	lst.pos = append(lst.pos, History[T]{pos: &arr})
	lst.elems = arr
}

func (lst *List[T]) ChangeToHistory(index int) {
	lst.elems = *lst.pos[index].pos
}

func (lst *List[T]) GetElems() string {
	return fmt.Sprintf("current elems %v \n pointer history %v ", lst.elems, lst.pos)
}

// - Create a program that uses goroutines and channels to perform a parallel calculation, such as finding the sum of elements in a large array.
type Counter struct {
	C int32
}

func Worker[T any](id int, job <-chan func() int32, res chan<- int32) {
	for j := range job {

		doJob := j
		go func() {
			result := doJob()
			res <- result
		}()
	}
}

func (p *Counter) Sum(arr []int32) int32 {
	fmt.Println(arr)

	for _, i := range arr {
		atomic.AddInt32(&p.C, i)
	}
	return p.C
}

func Divider(divide int, num []int32) [][]int32 {
	arr := [][]int32{}

	for i := 0; i < len(num); i += divide {
		arr = append(arr, num[i:i+divide])
	}
	return arr
}
