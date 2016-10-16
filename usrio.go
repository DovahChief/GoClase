package main

import "fmt"
import "strconv"

func main()  {
    fmt.Println("Introduce un numero del 1 al 10 ")
    var input string
    fmt.Scanln(&input)
     in, _ := strconv.ParseInt(input , 0, 64)

   if in == 1 {
       fmt.Println("uno")
   }
    if in == 2{
      fmt.Println("dos")
    }

}
