package main

import "fmt"

func reverse(p *[]int) { //*在此表示这里传递的是一个指针
	for i, j := 0, len(*p)-1; i < len(*p)/2; i, j = i+1, j-1 { //*在此表示获取指针所对应的基本值
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(&v) //&用来获取指针
	fmt.Println(v)

	a := "123"
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(a == *b)
}
