package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

type Object struct {
	key   string
	value int
	time time.Duration
}

func ObjectCreation(key string, value int) *Object {
	//This is sleep call is just used to stimulate some variance for each method call in the goroutine
	t := rand.N(5*time.Microsecond);
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
	//Example of how a channel buffer works
	var channel chan *Object = make(chan *Object, 100);

	for i := range 10 {
		channel <- ObjectCreation(strconv.Itoa(i), i);
	}

	for range 10 {
		fmt.Printf("%v\n", <-channel)
	}
	
}