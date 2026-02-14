package main

import "fmt"

func main() {
    var a, b int

    fmt.Print("Enter first number: ")
    fmt.Scan(&a)

    fmt.Print("Enter second number: ")
    fmt.Scan(&b)

    fmt.Println("\nResults:")
    fmt.Println("Sum:", a+b)
    fmt.Println("Difference:", a-b)
    fmt.Println("Product:", a*b)

    if b != 0 {
        fmt.Println("Division:", a/b)
        fmt.Println("Modulus:", a%b)
    } else {
        fmt.Println("Division: Cannot divide by zero")
        fmt.Println("Modulus: Cannot divide by zero")
    }
}
