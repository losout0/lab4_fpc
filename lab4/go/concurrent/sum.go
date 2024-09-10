package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Tuple struct{
	_sum int
	path string
}

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string, out chan Tuple) {
	data, err := readFile(filePath)

	if err == nil{}

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}
	
	dado := Tuple{_sum, filePath}

	out <- dado
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	info := make(chan Tuple)
	for _, path := range os.Args[1:] {
		//_sum, err := sum(path)
		go sum(path, info)

		//if err != nil {
		//	continue
		//}
		//sums[_sum] = append(sums[_sum], path)
	}

	for i := 0; i < len(os.Args[1:]); i++ {
		val := <- info
		totalSum += int64(val._sum)
		sums[val._sum] = append(sums[val._sum], val.path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

}
