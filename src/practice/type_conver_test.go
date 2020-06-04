package practice

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTypeConver1(t *testing.T) {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	f := 2222222222222222222222222222222222222222222222222345.897
	fmt.Printf("%d \n", int(f))

	s := "3.14sss"
	pf, e := strconv.ParseFloat(s, 32)
	fmt.Printf("%f , %e \n", pf, e)

	s1 := "123"
	i, e := strconv.ParseInt(s1, 0, 0)
	fmt.Printf("%d , %e \n", i, e)
}
