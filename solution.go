package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ConsoleReader(input string) []int {
	temp := strings.Split(input, " ")
	res := []int{}

	for i := 0; i < len(temp); i++ {
		tval, err := strconv.Atoi(temp[i])
		if err == nil {
			res = append(res, tval)
		} else {
			fmt.Printf("Value %s excluded\n", temp[i])
		}
	}
	return res
}

func MyReader() []int {
	var intInput []int
	var add string

	intInput = []int{}

	for add != "-" {
		fmt.Print("->")
		fmt.Fscan(os.Stdin, &add)
		tval, err := strconv.Atoi(add)
		if err == nil {
			intInput = append(intInput, tval)
		} else {
			if add != "-" {
				fmt.Printf("Value %s excluded\n", add)
			}

		}
	}
	return intInput
}

func Task178b(arr []int) int {
	i := 0
	for _, val := range arr {
		if (val%3 == 0) && (val%5 != 0) {
			i++
		}
	}
	return i
}

func IsSqrt(num int) bool {
	if num < 0 {
		return false
	}
	t := int(math.Sqrt(float64(num)))
	return t*t == num
}

func Task178v(arr []int) int {

	i := 0
	for _, val := range arr {
		if IsSqrt(val) && int(math.Sqrt(float64(val)))%2 == 0 {
			i++
		}
	}
	return i
}

func Task554(num int) []string {
	var t int
	res := []string{}
	for i := 1; i < num-1; i++ {
		for k := i; k < num; k++ {
			t = int(math.Sqrt(float64(i*i + k*k)))
			if IsSqrt(i*i+k*k) && t <= num {
				res = append(res, fmt.Sprint(i)+"^2 + "+fmt.Sprint(k)+"^2 = "+fmt.Sprint(t)+"^2 \n")
			}
		}
	}

	return res
}

func main() {
	var input string
	flag.StringVar(&input, "input", "none", "The name of a input")
	flag.Parse()
	var intInput []int
	if input == "none" {
		intInput = MyReader()
	} else {
		intInput = ConsoleReader(input)
	}

	fmt.Println(intInput, Task178b(intInput))
	fmt.Println(intInput, Task178v(intInput))
	fmt.Println(intInput[len(intInput)-1], Task554(intInput[len(intInput)-1]))

}
