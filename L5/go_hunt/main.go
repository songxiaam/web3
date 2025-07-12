package main

import (
	"fmt"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	printSymbol()
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
