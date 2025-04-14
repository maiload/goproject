package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.New(rand.NewPCG(1234, 5678))

	r := rand.IntN(100)
	c := 1

	for {
		fmt.Print("숫자값을 입력하세요: ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("ERROR - 숫자만 입력하세요.")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("축하합니다!", c, "번만에 정답을 맟추셨습니다!")
				break
			}
			c++
		}
	}
}
