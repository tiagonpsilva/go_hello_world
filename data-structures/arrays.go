package main

import "fmt"

func main() {
	// Array com tamanho fixo
	var numeros [5]int
	fmt.Println("Array zerado:", numeros)

	// Atribuindo valores
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50
	fmt.Println("Array preenchido:", numeros)

	// Inicialização direta
	nomes := [3]string{"João", "Maria", "Pedro"}
	fmt.Println("Array de nomes:", nomes)

	// Inicialização com tamanho implícito
	cores := [...]string{"vermelho", "verde", "azul", "amarelo"}
	fmt.Println("Array de cores:", cores)
	fmt.Println("Tamanho do array de cores:", len(cores))

	// Acessando elementos
	fmt.Println("Primeiro elemento:", numeros[0])
	fmt.Println("Último elemento:", numeros[len(numeros)-1])

	// Iterando sobre arrays
	fmt.Println("\nIterando sobre o array de nomes:")
	for i := 0; i < len(nomes); i++ {
		fmt.Printf("nomes[%d] = %s\n", i, nomes[i])
	}

	// Iterando com range
	fmt.Println("\nIterando com range sobre o array de cores:")
	for indice, cor := range cores {
		fmt.Printf("cores[%d] = %s\n", indice, cor)
	}

	// Slices (arrays dinâmicos)
	fmt.Println("\n--- Slices ---")
	
	// Criando um slice a partir de um array
	var slice1 []int = numeros[1:4] // do índice 1 até 3
	fmt.Println("Slice de numeros[1:4]:", slice1)

	// Declaração direta de slice (sem definir tamanho)
	frutas := []string{"maçã", "banana", "laranja"}
	fmt.Println("Slice de frutas:", frutas)

	// Usando make para criar um slice
	slice2 := make([]int, 5)      // tamanho 5, capacidade 5
	fmt.Println("Slice com make:", slice2)

	slice3 := make([]int, 3, 10)  // tamanho 3, capacidade 10
	fmt.Println("Slice com tamanho 3, capacidade 10:", slice3)
	fmt.Printf("Tamanho: %d, Capacidade: %d\n", len(slice3), cap(slice3))

	// Adicionando elementos (append)
	frutas = append(frutas, "uva")
	frutas = append(frutas, "morango", "pêra")
	fmt.Println("Slice de frutas após append:", frutas)

	// Slice de slice
	algunsFrutas := frutas[1:4]
	fmt.Println("algunsFrutas (slice de frutas[1:4]):", algunsFrutas)

	// Modificar o slice afeta o slice original
	algunsFrutas[0] = "banana-prata"
	fmt.Println("Slice original após modificação:", frutas)

	// Copiando slices
	destino := make([]string, len(frutas))
	copiados := copy(destino, frutas)
	fmt.Printf("Copiados %d elementos: %v\n", copiados, destino)

	// Arrays multidimensionais
	fmt.Println("\n--- Arrays Multidimensionais ---")
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matriz 3x3:", matriz)

	// Slices multidimensionais
	tabuleiro := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	
	// Modificando o tabuleiro
	tabuleiro[0][0] = "X"
	tabuleiro[2][2] = "O"
	
	// Imprimindo o tabuleiro
	fmt.Println("\nTabuleiro de jogo da velha:")
	for i := 0; i < len(tabuleiro); i++ {
		fmt.Printf("%s\n", tabuleiro[i])
	}
}
