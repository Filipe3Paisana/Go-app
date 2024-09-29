package main

import "fmt"

func main() {
	var nome string
	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)
	fmt.Println("Hello, " + nome + "!\n")
}