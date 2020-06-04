package practice

import (
	"fmt"
	"testing"
)

func TestIpChange(t *testing.T) {
	ip := "192.168.23.34"
	var newChars []rune
	for _, c := range ip {
		if c == '.' {
			newChars = append(newChars, '[', '.', ']')
		} else {
			newChars = append(newChars, c)
		}
	}
	fmt.Printf("new ip : %s \n", string(newChars))
}
