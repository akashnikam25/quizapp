package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	o, err := os.Open("quiz.csv")
	if err != nil {
		log.Println(err)
	}
	r := csv.NewReader(o)
	data, _ := r.ReadAll()

	var c int

	fmt.Println("Start Quiz	:	1.'Yes', 2.'No'")
	fmt.Println("Enter Your Choice:	")
	fmt.Scanf("%d", &c)

	if c == 1 {
		startQuiz(data)
	}

}
func startQuiz(d [][]string) {
	ch := make(chan int)
	var userRes int
	res := 0
	for i, qa := range d {
		fmt.Printf("Q.%d what is %s = ? \nPlease Enter your ans \n", i+1, qa[0])
		go getInput(ch)
		expectedAns, _ := strconv.Atoi(qa[1])
		select {
		case userRes = <-ch:
			fmt.Println("userRes :	", userRes)
			if userRes == expectedAns {
				res++
			}
		case <-time.After(5 * time.Second):
			fmt.Println("Times up")
		}
	}
	fmt.Println("Result : ", res, "out of", len(d))
	fmt.Println("Thank you")

}

func getInput(ch chan int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	res, _ := strconv.Atoi(scanner.Text())
	ch <- res
}
