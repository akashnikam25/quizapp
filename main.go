package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")
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
	res := 0
	for i, qa := range d {
		var ans string

		fmt.Printf("Q.%d what is %s = ? \nPlease Enter your ans \n",i+1,qa[0])
		fmt.Scanf("%s", &ans)

		if qa[1] == ans {
			res++
		}
	}

	fmt.Println("Result : ", res, "out of", len(d))
	fmt.Println("Thank you")
}
