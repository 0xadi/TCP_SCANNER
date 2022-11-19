package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup
var start_port int
var end_port int

func user_input() {
	fmt.Println("start_port:")
	fmt.Scan(&start_port)
	fmt.Println("end_port:")
	fmt.Scan(&end_port)
	port_range := end_port - start_port
	end_port_1 := (port_range / 10) + start_port
	end_port_2 := (port_range / 10) + end_port_1
	end_port_3 := (port_range / 10) + end_port_2
	end_port_4 := (port_range / 10) + end_port_3
	end_port_5 := (port_range / 10) + end_port_4
	end_port_6 := (port_range / 10) + end_port_5
	end_port_7 := (port_range / 10) + end_port_6
	end_port_8 := (port_range / 10) + end_port_7
	end_port_9 := (port_range / 10) + end_port_8
	end_port_10 := (port_range / 10) + end_port_9
	end_port_11 := (port_range / 10) + end_port_10
	end_port_12 := (port_range / 10) + end_port_11
	end_port_13 := (port_range / 10) + end_port_12
	end_port_14 := (port_range / 10) + end_port_13
	end_port_15 := (port_range / 10) + end_port_14
	wg.Add(15)
	go request_send(end_port_1+1, end_port_2)
	go request_send(end_port_2+1, end_port_3)
	go request_send(end_port_3+1, end_port_4)
	go request_send(end_port_4+1, end_port_5)
	go request_send(end_port_5+1, end_port_6)
	go request_send(end_port_6+1, end_port_7)
	go request_send(end_port_7+1, end_port_8)
	go request_send(end_port_8+1, end_port_9)
	go request_send(end_port_9+1, end_port_10)
	go request_send(end_port_10+1, end_port_11)
	go request_send(end_port_11+1, end_port_12)
	go request_send(end_port_12+1, end_port_13)
	go request_send(end_port_13+1, end_port_14)
	go request_send(end_port_14+1, end_port_15)
	go request_send(end_port_15+1, end_port)
	wg.Wait()

}

func request_send(t int, l int) {
	for i := t; i <= l; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		//error handling and sending tcp request if that will got error that will close the connection if not it will assign it to result
		_, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		print("open port:", i, "\n")
		break
		//close(c)
	}

	wg.Done()
}

//func loop(i int) {
//go request_send(i)
//m := <-d
// print("open port:", i, "\n")
//}
func main() {
	// hello := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// var hello int = 80
	// //p := hello
	// //hello2 := "working"
	// l := make(chan int)
	// go request_send(hello, l)
	// m := <-l
	// print("open port:", m, "\n")
	// wg.Add(1)
	// go request_send(80)
	// wg.Wait()
	user_input()
}
