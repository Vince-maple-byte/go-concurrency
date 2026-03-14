package main

import "fmt"

type Object struct {
	key   string
	value int
}

func ObjectCreation(key string, value int) *Object {
	return &Object{
		key:   key,
		value: value,
	}
}

func ObjectUpdate(o *Object, val int) {
	(*o).value = val
}

func ObjectDisplay(o *Object) {
	fmt.Printf("Address: %v\tKey: %s\tValue: %s\n", &o, o.key,o.value);
}

func main() {
	obj := make([]*Object, 10);

	for i := 0; i < 10; i++ {
		obj[i] = ObjectCreation(fmt.Sprintf("obj%d",i), i);
	}

	for i := 0; i < 10; i++ {
		ObjectDisplay(obj[i]);
	}

}