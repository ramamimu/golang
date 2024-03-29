package routines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("test goroutine")

	time.Sleep(1 * time.Second)
}

func TestManyRoutines(t *testing.T) {
	for i:=0; i<100000; i++ {
		go fmt.Println(i)
	}

	time.Sleep(10 * time.Second)
}

func TestChannel(t *testing.T) {
	// Create a channel
	ch := make(chan string)

	// this function called as closure
	go func() {
		time.Sleep(5 * time.Second)
		ch <- "Hello Channel"
	}()

	fmt.Println("waiting for channel")
	fmt.Println(<-ch)
	close(ch)
}

func TestChannelFunc(t *testing.T) {
	ch := make(chan string)
	go ChannelFunc(ch)
	fmt.Println("waiting for channel")
	fmt.Println(<-ch)
	close(ch)
}

// this function is an alternative to get channel data
// It's called range channel
func TestBufferChannel(t *testing.T){
	ch := make(chan string, 2)

	// this function called as closure
	go func() {
		for i:=0; i<100; i++ {
			// assign i to the string
			ch <- fmt.Sprintf("Channel (1) %d-%d", i, i+1)
			ch <- fmt.Sprintf("Channel (2) %d-%d",i, i+2)
			ch <- fmt.Sprintf("Channel (3) %d-%d",i, i+3)
			time.Sleep(5 * time.Second)
		}
	}()
	
	for data := range ch {
			fmt.Println(data)
			fmt.Println("the data?",<-ch)
			fmt.Println("=============")
	}

	close(ch)
}

func TestRaceCondition(t *testing.T) {
	var x = 0
	var mutex sync.Mutex

	// this function called as closure
	for i:=0; i<1000; i++ {
		go func() {
			mutex.Lock()
			for j:=0; j<100; j++ {
				x += 1
			}
			mutex.Unlock()
		}()
	}

	fmt.Println("waiting for channel")
	time.Sleep(1 * time.Second)
	fmt.Println("the value of x", x)
}


func TestWaitGroup(t *testing.T){
	// add wait state
	// var wg sync.WaitGroup
	wg := sync.WaitGroup{}
	// the mutex is to lock
	// to avoid race condition
	var mutex sync.Mutex

	var x = 0

	// this function called as closure
	wg.Add(1)
	for i:=0; i<1000; i++ {
		go func() {
			// defer to called variable when the function is done
			// function Donn() to decrement the counter
			defer wg.Done()
			// add the counter of waitgroup
			wg.Add(1)
			// lock the task
			mutex.Lock()
			for j:=0; j<100; j++ {
				x += 1
			}
			// unlock the task
			mutex.Unlock()
			}()
		}
	wg.Done()
	
	fmt.Println("waiting for channel")
	wg.Wait()
	fmt.Println("the value of x", x)
}