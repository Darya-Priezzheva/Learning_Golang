package main

import (
	"fmt"
	"time"
)

/*
Пока происходит вычисление числа, крутится спиннер загрузки
Запуск:
go run fib.go
*/

func main() {
	go spinner(100 * time.Millisecond)
	n := 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}

	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
