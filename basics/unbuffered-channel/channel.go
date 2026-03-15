package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Object struct {
	key   string
	value int
	time time.Duration
}

func ObjectCreation(key string, value int) *Object {
	//This is sleep call is just used to stimulate some variance for each method call in the goroutine
	t := rand.N(5*time.Second);
	time.Sleep(t)
	return &Object{
		key:   key,
		value: value,
		time: t,
	}
}

func ObjectUpdate(o *Object, val int) {
	(*o).value = val
}

func ObjectDisplay(o *Object) string {
	return fmt.Sprintf("Address: %v\tKey: %s\tValue: %d\tTime: %v\n", &o, o.key, o.value, o.time);
}

func main() {
	//How a simple unbuffered channel is done in golang.
	obj := make([]*Object, 10);
	channel := make(chan string, 0);

	for i := 0; i < 10; i++ {
		go func (index int) {
			obj[index] = ObjectCreation(fmt.Sprintf("obj%d",index), index);
			channel <- ObjectDisplay(obj[index]);
		}(i);
		
	}

	for i := 0; i < 10; i++ {
		fmt.Print(<-channel);
	}

	/*
		fmt.Println(<-channel);
		channel <- "Noice";
		will cause a deadlock error to happen because that the sender and receiver calls
		are happening at the same thread(routine). 
		When this happens the receive call, <-channel, is the first call the main thread is waiting for
		a sender call to happen in another thread, so the main thread is blocked until then. 
		Because of this, the main thread never reaches the sender call giving the deadlock error. 
		
		Below is the correct way of doing it.
	*/
	go func() {
		fmt.Print(<-channel);
	}()
	
	channel <- "Noice";
	
}