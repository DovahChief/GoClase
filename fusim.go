package main

import "fmt"

func vet(x int) (bool , int)  {
    if x  == 1{
        return true , 1
    }else {
        return false , 0
    }
}

func main()  {

    a := 1

    if b , c := vet(a) ; c == 1{
        fmt.Println("a = 1 ", b)
    }

}
