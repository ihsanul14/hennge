package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	test_case int
	iterate   int
	result    []int
	iteration = 1
	sum_res   = 0
)

func main() {
	fmt.Print("input test_case number : ")
	fmt.Scanln(&test_case)
	if test_case < 1 || test_case > 100 {
		fmt.Println("test case must be number and between 1 - 100")
		return
	}
	fmt.Print("input iteration number : ")
	fmt.Scanln(&iterate)
	input()
	fmt.Println(result)
}

func checkIteration(val int) {
	fmt.Println("iteration - ", iteration)
	if iteration != iterate {
		iteration++
		input()
	}
}

func input() {
	var val string
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	val = string(bytes)
	panjang := strings.Split(val, " ")
	checkInputElement(panjang, 0)
	checkIteration(len(panjang))
}

func checkInputElement(val []string, lenght int) {
	if lenght != len(val) {
		parseElem(val, lenght)
	}
	if sum_res != 0 {
		result = append(result, sum_res)
		result = result[:test_case]
	}
	sum_res = 0
}

func parseElem(val []string, lenght int) {
	arr_elem, err := strconv.Atoi(val[lenght])
	if err != nil {
		fmt.Println("invalid input, must be a number")
		os.Exit(1)
	}
	inputValidator(arr_elem)
	sumInput(arr_elem)
	lenght++
	checkInputElement(val, lenght)

}

func square(val int) int {
	return val * val
}

func sumInput(val int) {
	if val > 0 {
		sum_res += square(val)
	}
}

func inputValidator(val int) {
	if val > 100 || val < -100 {
		fmt.Println("invalid input, max number is 100")
		os.Exit(1)
	}
}
