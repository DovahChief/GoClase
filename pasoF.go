package main

   import "fmt"


   func plusX(x int) (func(v int) (int)) {
      return func(v int) (int) {
          return v+x
      }
   } //funcion que recibe una funcion como param

   func main() {
       p := plusX(5)
       fmt.Printf("3+5: %d\n", p(3))

       px := plusX(3)
       fmt.Printf("3+3: %d\n", px(3))
   } //metodo principal
