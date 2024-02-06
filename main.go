package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	var (
		f string
		t int
	)

	flag.StringVar(&f, "fileName", "problems.csv", "default value of file is problems.csv")
	flag.IntVar(&t, "timer", 10, "default timer for each ques is 10s")
	flag.Parse()
	fmt.Println("timer	:	", t)
	o, err := os.Open(f)
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
		startQuiz(data, time.Duration(t))
	}

}
func startQuiz(d [][]string, timer time.Duration) {

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
		case <-time.After(timer * time.Second):
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
