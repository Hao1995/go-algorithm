package decodestring

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrorIsEmpty = errors.New("data is empty")
)

type Stack struct {
	data []byte
}

func (s *Stack) Push(item byte) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (byte, error) {
	var res byte

	if len(s.data) == 0 {
		return res, ErrorIsEmpty
	}

	res, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]

	return res, nil
}

func (s *Stack) Show() (byte, error) {
	if len(s.data) == 0 {
		var res byte
		return res, ErrorIsEmpty
	}

	return s.data[1], nil
}

func decodeStringV1(s string) string {
	stash := &Stack{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ']' {
			// pop data until get a '['
			strList := getRepeatedByteFromStack(stash)

			// get repeated number: k
			k := getKFromStack(stash)

			// repeate encoded string
			var decodedList []byte
			for i := 0; i < k; i++ {
				decodedList = append(decodedList, strList...)
			}

			// push back to stash
			for _, b := range decodedList {
				stash.Push(b)
			}
		} else {
			// push to stash
			stash.Push(c)
		}
	}

	// convert stash to string
	var ansByte []byte
	for {
		b, err := stash.Pop()
		if err != nil {
			break
		}
		ansByte = append([]byte{b}, ansByte...)
	}

	return string(ansByte)
}

func getRepeatedByteFromStack(stack *Stack) []byte {
	var strList []byte
	for {
		ob, err := stack.Pop()
		if err != nil {
			return strList
		}

		if ob == '[' {
			break
		}

		strList = append([]byte{ob}, strList...)
	}
	return strList
}

func getKFromStack(stack *Stack) int {
	var numByte []byte
	for {
		ob, err := stack.Pop()
		if err != nil {
			break
		}

		if num := ob - '0'; !(num >= 0 && num <= 9) {
			// not number
			stack.Push(ob)
			break
		}
		numByte = append([]byte{ob}, numByte...)
	}

	k, err := strconv.Atoi(string(numByte))
	if err != nil {
		fmt.Printf("Can't convert %v to int", string(numByte))
		return 0
	}

	return k
}
