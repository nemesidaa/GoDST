package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Вам предстоит обратная задача: по указанной строке восстановить множество чисел,
из которого указанная строка была построена.
Заметьте, что выводить множество нужно в отсортированном порядке.
1-6,8-9,11 => 1 2 3 4 5 6 8 9 11
*/

const (
	Sep   = ','
	Defis = '-'
)

func restore(from string, to string) []byte {
	resp := make([]byte, 0)
	fromVal, err := strconv.Atoi(from)
	if err != nil {
		panic(err)
	}

	toVal, err := strconv.Atoi(to)
	if err != nil {
		panic(err)
	}

	for fromVal < toVal+1 {
		resp = append(resp, []byte(fmt.Sprintf("%d ", fromVal))...)
		fromVal++
	}
	return resp
}

func First() {
	reader := bufio.NewReader(os.Stdin)
	unparsed, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	resp := make([]byte, 0)
	instArr := strings.FieldsFunc(unparsed, func(r rune) bool {
		return r == Sep
	})

	for _, v := range instArr {
		if strings.Contains(v, string(Defis)) {
			split := strings.Split(v, string(Defis))
			if len(split) != 2 {
				panic("Can't split")
			}

			resp = append(resp, restore(split[0], split[1])...)
		} else {
			resp = append(resp, []byte(fmt.Sprintf("%s ", v))...)
		}
	}

	fmt.Println(string(resp))
}
