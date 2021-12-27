package main

import (
	"fmt"
	"time"
)

type S struct {
	Id   int
	Name string
}

//var mux sync.Mutex

func main() {
	//title := "gtrg55"
	//reg1 := regexp.MustCompile(`([^\d]*)([0-9]*)`)
	//if reg1 == nil { //解释失败，返回nil
	//	fmt.Println("regexp err")
	//	return
	//}
	//
	////根据规则提取关键信息
	//result1 := reg1.FindAllStringSubmatch(title, -1)
	//fmt.Println(result1[0][2])
	//
	//panic(3)

	ch := make(chan S, 100)

	go outs(ch, 1)
	go outs(ch, 2)

	for i := 1; i <= 100; i++ {
		ss := &S{i, string(i)}
		go ints(ch, *ss)

	}

	//close(ch)

	//time.Sleep(10 * time.Second)

	time.Sleep(time.Second * 1)

}

func ints(ch chan S, s S) {
	ch <- s

}

func outs(ch chan S, i int) {
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
