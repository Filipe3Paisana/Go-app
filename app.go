package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
    "sync"
    "time"
)

// Gerar número aleatório
var random int
var mu sync.Mutex

func generateRandomNumber() {
    mu.Lock()
    random = rand.Intn(5) + 1 
    mu.Unlock()
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bem-vindo à minha aplicação web em Go! Use /hello para um cumprimento ou /random para jogar um jogo de adivinhação.")
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Olá, visitante! Você acessou a rota /hello.")
}


func randomHandler(w http.ResponseWriter, r *http.Request) {
    
    generateRandomNumber()

    fmt.Fprintf(w, "Um número aleatório foi gerado. Tente adivinhar entre 1 e 5! Acesse /guess/{n} para adivinhar.")
}


func guessHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    
    guess := r.URL.Path[len("/guess/"):]

    
    number, err := strconv.Atoi(guess)
    if err != nil {
        http.Error(w, "Por favor, forneça um número válido.", http.StatusBadRequest)
        return
    }

    
    if number == random {
        fmt.Fprintf(w, "Parabéns! Você adivinhou o número: %d", random)
    } else {
        fmt.Fprintf(w, "Errado! O número era %d. Tente novamente!", random)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano()) 

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/random", randomHandler)
    http.HandleFunc("/guess/", guessHandler)

    fmt.Println("Servidor rodando em http://localhost:8080/")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
