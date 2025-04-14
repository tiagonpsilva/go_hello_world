package main

import "fmt"

// Função simples sem parâmetros e sem retorno
func saudacao() {
	fmt.Println("Olá, seja bem-vindo!")
}

// Função com parâmetros
func somar(a, b int) int {
	return a + b
}

// Função com múltiplos retornos
func operacoes(a, b int) (int, int, int, float64) {
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	var divisao float64 = 0
	if b != 0 {
		divisao = float64(a) / float64(b)
	}
	return soma, subtracao, multiplicacao, divisao
}

// Função com retornos nomeados
func calcularRetangulo(largura, altura float64) (area, perimetro float64) {
	area = largura * altura
	perimetro = 2 * (largura + altura)
	return // retorno limpo quando os retornos são nomeados
}

// Função variádica (aceita número variável de argumentos)
func somarTodos(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// Função que recebe função como parâmetro (função de ordem superior)
func aplicarOperacao(a, b int, operacao func(int, int) int) int {
	return operacao(a, b)
}

// Função anônima e closure
func gerarContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

// Função com defer (executa ao final da função)
func exemploDefer() {
	defer fmt.Println("Esta mensagem aparece por último (defer)")
	fmt.Println("Esta mensagem aparece primeiro")
}

func main() {
	// Função simples
	saudacao()

	// Função com parâmetros
	resultado := somar(5, 3)
	fmt.Println("Soma:", resultado)

	// Função com múltiplos retornos
	s, sub, mult, div := operacoes(10, 2)
	fmt.Printf("Operações com 10 e 2: soma=%d, subtração=%d, multiplicação=%d, divisão=%.2f\n",
		s, sub, mult, div)

	// Ignorando retornos com _
	soma2, _, _, _ := operacoes(8, 4)
	fmt.Println("Apenas a soma:", soma2)

	// Função com retornos nomeados
	a, p := calcularRetangulo(5.0, 3.0)
	fmt.Printf("Retângulo: área=%.2f, perímetro=%.2f\n", a, p)

	// Função variádica
	total := somarTodos(1, 2, 3, 4, 5)
	fmt.Println("Soma de 1 a 5:", total)

	// Passando slice para função variádica
	numeros := []int{10, 20, 30, 40, 50}
	totalSlice := somarTodos(numeros...)
	fmt.Println("Soma do slice:", totalSlice)

	// Função de ordem superior
	multiplicar := func(x, y int) int {
		return x * y
	}
	resultadoMult := aplicarOperacao(6, 7, multiplicar)
	fmt.Println("Multiplicação via função de ordem superior:", resultadoMult)

	// Função anônima direta
	resultadoAnon := aplicarOperacao(6, 7, func(x, y int) int {
		return x - y
	})
	fmt.Println("Subtração via função anônima:", resultadoAnon)

	// Closure
	contador := gerarContador()
	fmt.Println("Contador:", contador()) // 1
	fmt.Println("Contador:", contador()) // 2
	fmt.Println("Contador:", contador()) // 3

	// Novo contador independente
	contador2 := gerarContador()
	fmt.Println("Contador2:", contador2()) // 1

	// Exemplo de defer
	exemploDefer()
}
