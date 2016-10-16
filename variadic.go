package main

import "fmt"

func varsum(num ...int) {
	sum := 0
	fmt.Print(num, " ")
	for _, n := range num {
		sum += n
	}
	fmt.Println(sum)
} // funcion variadica

func print(elem ... string)  {
	for _, n := range elem{
		fmt.Print(n+" ")
	}
}// funcion que imprime


func main() {

	varsum(1, 2, 3, 4, 5, 6)
	nums := []int{1, 2, 3, 4}

	varsum(nums...) // pasar a arreglo de funcion variadica

	if len(nums) < 2{
		print("menor", "dos")
	} else {
		print("mayor","que","dos")
	}

} // metodo principal
