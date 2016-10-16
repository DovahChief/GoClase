package main

//#include <stdio.h>
// void taco()
// {
//	    printf("\nse de los tacos soy codigo en C");
// }
// float sum(float a, float b){
//     return a+b;
// }
import "C"
 
import "fmt"

func main()  {
    C.taco()
    x := C.sum(1.2 , 3.5)
    fmt.Println("no se nada de suma soy GOpher")
    fmt.Println(x)
}
