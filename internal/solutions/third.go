package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Выведите возможный пароль от ноутбука, удовлетворяющий указанным условиям.
Если вариантов пароля несколько, выберите тот, который начинается в последовательности
из первой строки правее (позже) других, а среди всех с одинаковым с ним началом — самый длинный.
==================================|
abacaba
abc        => caba
4
==================================|
*/

func CheckValidity(testing, condset string) bool {
	for _, el := range condset {
		if !strings.Contains(testing, string(el)) {
			return false
		}
	}
	return true
}

func Third() {
	reader := bufio.NewReader(os.Stdin)
	combo, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	combo = strings.TrimSpace(combo)
	symbolKit, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	symbolKit = strings.TrimSpace(symbolKit)
	symVal, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	symVal = strings.TrimSpace(symVal)
	symValint, err := strconv.Atoi(symVal)
	if err != nil {
		panic(err)
	}

	// fmt.Println(symbolKit + " " + combo + " " + symVal)
	if len(combo) < symValint {
		fmt.Println("-1")
		return
	}
	ri := len(combo)
	li := ri - symValint
	for {
		if li < 0 {
			break
		} else if CheckValidity(combo[li:ri], symbolKit) {
			fmt.Println(combo[li:ri])
			break
		}
		ri--
		li--
	}
}
