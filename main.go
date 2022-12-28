package main

import (
	"fmt"
)

func greet(c chan string, cntrol chan int) {
	fmt.Println("Init gret")
	for {
		select {
		case i := <-c:
			{
				fmt.Println("Hello " + i + "!")
				if len(c) == 0 {
					cntrol <- 0
				}
			}

		}
	}
}

func main() {
	var names = [...]string{"A", "b", "c", "d", "e","dd","323"}

	fmt.Println("main() started")
	c := make(chan string, 2)
	control := make(chan int)
	go greet(c, control)
	defer close(c)
	defer close(control)

	cnt := 0
	for idx, item := range names {

		fmt.Println("+ Add to chain item ", item)
		c <- item
		cnt++
		fmt.Printf("= add done %s. len chain=%d\n", item, cnt)
		if cnt == 2 || idx == len(names)-1 {
			select {
			case k := <-control:
				{
					fmt.Println("UnLock main gorutine", k)
					cnt = 0
				}
			}
		}
	}

	fmt.Println("main() stopped")
}
