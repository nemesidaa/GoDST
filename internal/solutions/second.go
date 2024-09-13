package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solveSnowProblem(n int, a []int) ([]int, bool) {
	// Копируем массив для вывода возможных значений
	result := make([]int, n)
	copy(result, a)

	// Восстановление значений
	for i := 0; i < n; i++ {
		if a[i] != -1 {
			result[i] = a[i]
		} else {
			// Устанавливаем значение как максимум из предыдущего дня
			if i == 0 {
				result[i] = 1 // Первый день может быть любым положительным числом
			} else {
				result[i] = result[i-1] // Снег не может уменьшаться
			}
		}
	}

	// Проверка корректности
	for i := 0; i < n; i++ {
		if i > 0 && result[i] < result[i-1] {
			return nil, false // Значение не может быть меньше предыдущего
		}
		if a[i] != -1 && result[i] != a[i] {
			return nil, false // Значение не совпадает с известным
		}
	}

	// Если не найдено нарушений, возвращаем результат
	return result, true
}

func Second() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение количества дней
	dVal, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	dVal = strings.TrimSpace(dVal)
	dInt, err := strconv.Atoi(dVal)
	if err != nil {
		panic(err)
	}

	// Чтение данных о снегопадах
	daysc, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	dArr := strings.Split(daysc, " ")
	days := make([]int, len(dArr))
	for i, v := range dArr {
		v = strings.TrimSpace(v)
		if v == "-1" {
			days[i] = -1
		} else {
			days[i], err = strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
		}
	}

	// Решение задачи
	result, valid := solveSnowProblem(dInt, days)
	if !valid {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(strings.Join(intSliceToStringSlice(result), " "))
	}
}

// Преобразование слайса int в слайс string
func intSliceToStringSlice(slice []int) []string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = strconv.Itoa(v)
	}
	return strSlice
}
