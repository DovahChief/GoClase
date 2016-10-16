package main

import ("fmt"
        "time")

func mensaje(mesg string)  {
    i := 0
    for i <= 5{
        fmt.Println("funcion mensaje:", mesg ,i )
        time.Sleep(10 * time.Millisecond)
        i++
    }//ciclo de trabajo
} // funcion imprime un mensaje
func suma(val int, msj string) (int,string) {
    return val + 100 , "origen : "+msj
}// funcion de suma 100 regresa dos valores

func main()  {
    var edad int  = 22
    var pi float64 = 3.14167
    num := 1
    numero1 := 20
    numero2 := 15
    nomf := "sum"

    fmt.Printf("edad : %d \nvalor de pi : %f \nvalor de num : %d", edad, pi, num)
    go mensaje("go") // corre el codigo en el fondo
    mensaje("normal")

    a , b := suma(numero1, nomf) //asignacion doble
    c , d := suma(numero2, nomf)

    fmt.Println(" ",a,b,"\n",c,d)


}
