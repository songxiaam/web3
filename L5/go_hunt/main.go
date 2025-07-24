package main

import (
	"fmt"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//printSymbol()
	//testSlice()
	//testSlice2()
	//testSlice3()
	//testMap()
	testString()
}

func testString() {
	var a string = "ABCDE"
	fmt.Println(a)
	b := []byte(a)
	fmt.Println(b)
	b[0] = 'B'
	fmt.Printf(string(b))

}

func testMap() {
	m := map[string]string{"a": "1", "b": "2"}

	fmt.Println(m)
	a := m["a"]
	fmt.Println(a)

	b, ok := m["c"]
	if !ok {
		fmt.Println(ok)
	}
	fmt.Println(b)
}

func testSlice3() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
}

func testSlice2() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	//for _, i := range s {
	//	fmt.Println(i)
	//	i++
	//}

	for i := range s {
		s[i] += 1
	}
}

func testSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := slice[2:5]
	s2 := slice[2:6:6]
	fmt.Println(cap(s2))
	s2 = append(s2, 10)
	fmt.Println(s2)
	fmt.Println(slice)

}

func printSymbol() {
	letter, number := make(chan bool), make(chan bool)

	wait := sync.WaitGroup{}
	maxNum := 28

	wait.Add(1)
	go func() {
		defer wait.Done()
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d", i)
				i++
				fmt.Printf("%d", i)
				i++
				if i >= maxNum {
					return
				}
				letter <- true
			}
		}
	}()

	go func() {
		l := 'A'
		for {
			select {
			case <-letter:
				fmt.Printf("%c", l)
				l++
				fmt.Printf("%c", l)
				l++
				number <- true
				if l >= 'Z' {
					return
				}
			}
		}
	}()
	number <- true
	wait.Wait()

}
