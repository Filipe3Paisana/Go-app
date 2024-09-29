package main

import "fmt"

func main() {
    var nome string
    fmt.Println("Digite seu nome:")
    fmt.Scanln(&nome)
    fmt.Printf("Hello, %s!\n", nome)
}
