package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("How many digits would you like calculate to (up to 20)?")

	for {
		fmt.Printf("\n->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		i, err := strconv.Atoi(text)
		if err != nil || i == 0 {
			fmt.Printf("Please enter an integer to continue.")
		}
		if i > 20 {
			fmt.Printf("Please enter an integer of 20 or less.")
		}
		if i <= 20 && i != 0 {
			calculate(i)
			break
		}
	}
}

func calculate(i int) {
	//outer loop checks how many digits there are in loop runs Nilakantha Series https://www.mathscareers.org.uk/calculating-pi/

	var pi float64 = 3
	var is_addition bool = true
	var first_in_series int = 2

	for i := 0; i < 999999999; i++ {
		if is_addition {
			pi += (4 / float64(getDenominator(first_in_series)))
		} else {
			pi -= (4 / float64(getDenominator(first_in_series)))
		}
		first_in_series += 2
		is_addition = !is_addition
	}
	var round_value float64 = math.Pow(10, float64(i))
	pi = math.Round(pi*round_value) / round_value
	fmt.Println(pi)
}

func getDenominator(start int) int {
	var denominator int = (start * (start + 1) * (start + 2))
	return denominator
}
