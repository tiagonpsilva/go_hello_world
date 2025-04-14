package main

import (
	"fmt"
	"time"
)

func main() {
	// Exemplo de if/else
	idade := 18
	
	if idade >= 18 {
		fmt.Println("Você é maior de idade")
	} else if idade >= 16 {
		fmt.Println("Você é adolescente")
	} else {
		fmt.Println("Você é menor de idade")
	}

	// if com inicialização de variável
	if nota := 8.5; nota >= 7.0 {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Reprovado com nota:", nota)
	}

	// Exemplo de for (loop tradicional)
	fmt.Println("\nLoop tradicional:")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Exemplo de for (como while)
	fmt.Println("\nLoop estilo while:")
	contador := 1
	for contador <= 5 {
		fmt.Print(contador, " ")
		contador++
	}
	fmt.Println()

	// Exemplo de for range (para arrays/slices)
	fmt.Println("\nLoop com range:")
	numeros := []int{10, 20, 30, 40, 50}
	for indice, valor := range numeros {
		fmt.Printf("numeros[%d] = %d\n", indice, valor)
	}

	// Exemplo de switch
	fmt.Println("\nSwitch por valor:")
	diaDaSemana := time.Now().Weekday()
	switch diaDaSemana {
	case time.Saturday, time.Sunday:
		fmt.Println("É fim de semana!")
	default:
		fmt.Println("É dia útil!")
	}

	// Switch sem expressão (como múltiplos if-else)
	fmt.Println("\nSwitch como múltiplos if-else:")
	hora := time.Now().Hour()
	switch {
	case hora < 12:
		fmt.Println("Bom dia!")
	case hora < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}

	// Loop infinito com break
	fmt.Println("\nLoop com break:")
	soma := 0
	for {
		soma++
		if soma > 5 {
			break
		}
		fmt.Print(soma, " ")
	}
	fmt.Println()

	// Loop com continue
	fmt.Println("\nLoop com continue (números ímpares):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Pula números pares
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}
