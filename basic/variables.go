package main

import "fmt"

func main() {
	// Declaração com tipo explícito
	var nome string = "John"
	var idade int = 30
	var altura float64 = 1.75
	var isAdulto bool = true

	// Declaração com inferência de tipo
	cidade := "São Paulo"
	temperatura := 25.5

	// Constantes
	const PI = 3.14159

	// Múltiplas declarações
	var (
		x, y int    = 10, 20
		a, b string = "Hello", "World"
	)

	// Exibindo os valores
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura, "metros")
	fmt.Println("É adulto?", isAdulto)
	fmt.Println("Cidade:", cidade)
	fmt.Println("Temperatura:", temperatura, "°C")
	fmt.Println("Valor de PI:", PI)
	fmt.Println("Valores x e y:", x, y)
	fmt.Println("Valores a e b:", a, b)

	// Formatação de strings
	fmt.Printf("Nome: %s, Idade: %d, Altura: %.2f\n", nome, idade, altura)
}
