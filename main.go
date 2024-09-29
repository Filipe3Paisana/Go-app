package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    
    rand.Seed(time.Now().UnixNano())
    
    
    var random int = rand.Intn(10) + 1
    
    var numero int

    
    for numero != random {
        fmt.Println("Digite um número:")
        fmt.Scan(&numero)
        
        if numero < random {
            fmt.Println("O número é maior.")
        } else if numero > random {
            fmt.Println("O número é menor.")
        }
    }

    fmt.Println("Parabéns, você acertou!")
}
