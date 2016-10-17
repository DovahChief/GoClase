package main

import "fmt"

func main() {
    arreg := [5]int{1,2,3,4,5}
	slice := arreg[:]
    fmt.Println("Slice :" , slice)
    fmt.Println("Numeroo de elementos en slice : " , len(slice) )
    fmt.Println("Capacidad de slice : ", cap(slice))
    fmt.Printf("\nDireccion de arreglo = %p \nDireccion de slice = %p", &arreg, &slice[0])

    fmt.Println("\n")
    slice = append(slice, 117)

    fmt.Println("Slice :" , slice)
    fmt.Println("Numeroo de elementos en slice : " , len(slice) )
    fmt.Println("Capacidad de slice : ", cap(slice))
    fmt.Printf("\nDireccion de arreglo = %p \nDireccion de slice = %p", &arreg, &slice[0])

}
