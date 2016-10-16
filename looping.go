package main

import "fmt"

func main()  {
    i := 1
    num := []int {23,25,32,56,78}

    for i < 5{
        fmt.Printf("ciclo 1 :%d \n", i)
        i++
    } // primer ejemplo de ciclo

    fmt.Print("\n----------\n\n")

    for n := range num{
        fmt.Printf("ciclo 2 : %d \n", n)
    }// segundo ejemplo de ciclo

    fmt.Print("\n----------\n\n")

    for x , n := range num{
        fmt.Printf("ciclo 3 : [%d] :  %d \n", x , n)
    }// segundo ejemplo de ciclo

}// funcion main
