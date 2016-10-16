package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
} //funcion que simula carga pesada de trabajo

func main() {
	go say("hola")
	say("mundo")
    //say ("cruel")
}// funcion principal
