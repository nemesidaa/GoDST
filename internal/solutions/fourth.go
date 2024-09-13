package solutions

import (
	"fmt"
	"math"
)

/*
Вам нужно решить очень понятную задачу: посчитать количество составных чисел от
l до r,  количество делителей которых при этом является простым числом.
*/

func countDivisors(n int) int {
	count := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			count++
			if i != n/i {
				count++
			}
		}
	}
	return count
}

func isPrime(n int) bool {
	return countDivisors(n) > 2
}

func Fourth() {
	var l, r int
	fmt.Scan(&l, &r)
	var count int8
	for i := l; i <= r; i++ {
		if isPrime(countDivisors(i)) {
			count++
		}
	}

	fmt.Println(count)
}
