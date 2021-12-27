package main

import (
	"fmt"
	"sync"
	"time"
)

type Y struct {
	Id   int
	Name string
}

func main() {

	tasks := []func(){
		func() { time.Sleep(time.Second); fmt.Println("1 sec later") },
		func() { time.Sleep(time.Second * 2); fmt.Println("2 sec later") },
	}

	var wg sync.WaitGroup // 1-1
	wg.Add(len(tasks))    // 1-2

	fmt.Println(len(tasks))
	for _, task := range tasks {
		task := task
		go func() { // 1-3-1
			defer wg.Done() // 1-3-2
			task()          // 1-3-3
		}() // 1-3-1
	}
	wg.Wait() // 1-4
	fmt.Println("exit")

	//ch := make(chan S, 100)
	//
	//go outs(ch, 1)
	//go outs(ch, 2)
	//
	//for i := 1; i <= 100; i++ {
	//	ss := &S{i, string(i)}
	//	go ints(ch, *ss)
	//
	//}

	//close(ch)

	//time.Sleep(10 * time.Second)

	//time.Sleep(time.Second * 1)

}

func intss(ch chan Y, s Y) {
	ch <- s

}

func outss(ch chan Y, i int) {
	for {
		select {
		case ss, ok := <-ch:

			if ok {
				fmt.Println(" chan ", ss, i)
			} else {
				fmt.Println("chan finish", time.Now(), i)
				break
			}

		}
	}

}
