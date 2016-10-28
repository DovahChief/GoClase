package main

import "fmt"

// creamos un tipo con a definicion de la funcion para que sea
// mas sencillo de pasar

type convert func(int) string

//una funcion  que cumpla con el tipo
func conv(x int) string {
	return fmt.Sprintf("%v", x)
}

func creaCad(fun convert, x int) {
	fmt.Println("hey soy la funcion que recibe funciones")
	fmt.Printf("mensaje : %s ", fun(x))
}

func main() {
	creaCad(conv, 117)
}
