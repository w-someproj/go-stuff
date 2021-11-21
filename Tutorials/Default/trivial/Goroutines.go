package trivial

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0

func Goroutines() {
	WaitGroup()
	//Mutex()

	//results := make(map[int]int)
	//structCh := make(chan struct{})
	//go factorialSync(20, structCh, results)
	//<-structCh       // wait for close
	//
	//for i, v := range results{
	//	fmt.Println(i, " - ", v)
	//}
	//fmt.Println(`End`)

	//intCh := make(chan int)
	//go factorial(20, intCh)
	//fmt.Println(<-intCh)
	//fmt.Println("The End")
	//for i:= 1; i < 7; i++{
	//	go factorial(i, intCh)
	//}
	//fmt.Println(<-intCh)
	//fmt.Println("The End")

	//intCh := make(chan int)
	//go func(){
	//	for i:= 0; i < 100; i ++{
	//		fmt.Println(i)
	//		if i == 50 {
	//			intCh <- i
	//		}}
	//
	//}()
	//fmt.Println(`Some`)
	//fmt.Println(`From channel`,<-intCh)
	//fmt.Println("The End")
}

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Goroutine %d starts \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Goroutine %d ends \n", id)
	}

	go work(1)
	go work(2)

	wg.Wait()
	fmt.Println("All end")
}

func Mutex() {
	ch := make(chan bool)
	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The End")
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock()
	ch <- true
}

func factorial(n int, intCh chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
	intCh <- result
}

func factorialSync(n int, ch chan struct{}, results map[int]int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}
