package main

import "fmt"
import "os"
import "os/exec"

//comentario

func main()  {
    fmt.Println("Ejecutando en GO")
    fmt.Println(os.Getwd())
    _, out := exec.Command("python3 ej.py").Output()
    fmt.Println(out)
} // primera funcion
