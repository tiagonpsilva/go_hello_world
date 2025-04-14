package main

import (
	"fmt"
	"log"
	"math"
)

// Definindo uma interface
type Forma interface {
	Area() float64
	Perimetro() float64
}

// Implementação para Círculo
type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

// Implementação para Retângulo
type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

// Implementação para Triângulo
type Triangulo struct {
	LadoA, LadoB, LadoC float64
	Base, Altura         float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangulo) Perimetro() float64 {
	return t.LadoA + t.LadoB + t.LadoC
}

// Função que recebe uma interface
func exibirInfo(f Forma) {
	fmt.Printf("Área: %.2f\n", f.Area())
	fmt.Printf("Perímetro: %.2f\n", f.Perimetro())
}

// Interface vazia - pode receber qualquer tipo
func descreverTipo(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
}

// Interface para String personalizada
type Stringer interface {
	String() string
}

func (c Circulo) String() string {
	return fmt.Sprintf("Círculo com raio %.2f", c.Raio)
}

func (r Retangulo) String() string {
	return fmt.Sprintf("Retângulo com dimensões %.2f x %.2f", r.Largura, r.Altura)
}

// Interface para log personalizado
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
	Prefix string
}

func (l ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.Prefix, message)
}

type FileLogger struct {
	FileName string
}

func (l FileLogger) Log(message string) {
	// Simplificado: em um caso real, escreveria em um arquivo
	fmt.Printf("[%s] %s\n", l.FileName, message)
}

// Função que usa a interface Logger
func processarDados(logger Logger, dados string) {
	logger.Log("Iniciando processamento")
	// Processamento simulado
	logger.Log("Dados processados: " + dados)
	logger.Log("Processamento concluído")
}

func main() {
	// Configurando logger para o exemplo
	log.SetPrefix("[INTERFACES] ")
	log.Println("Iniciando exemplo de interfaces...")

	// Criando instâncias das formas
	circulo := Circulo{Raio: 5}
	retangulo := Retangulo{Largura: 4, Altura: 3}
	triangulo := Triangulo{
		LadoA:  3,
		LadoB:  4,
		LadoC:  5,
		Base:   3,
		Altura: 4,
	}

	// Usando a interface Forma
	fmt.Println("\n=== Usando interface Forma ===")
	fmt.Println("Informações do círculo:")
	exibirInfo(circulo)

	fmt.Println("\nInformações do retângulo:")
	exibirInfo(retangulo)

	fmt.Println("\nInformações do triângulo:")
	exibirInfo(triangulo)

	// Slice de interfaces
	formas := []Forma{circulo, retangulo, triangulo}

	fmt.Println("\n=== Iterando sobre slice de Formas ===")
	for i, forma := range formas {
		fmt.Printf("Forma %d: \n", i+1)
		exibirInfo(forma)
		fmt.Println()
	}

	// Interface vazia (any)
	fmt.Println("=== Interface vazia (any) ===")
	descreverTipo(42)
	descreverTipo("Olá")
	descreverTipo(true)
	descreverTipo(3.14)
	descreverTipo(circulo)

	// Type assertions
	fmt.Println("\n=== Type assertions ===")
	var forma Forma = Circulo{Raio: 3}
	
	// Verificando se é um círculo
	if c, ok := forma.(Circulo); ok {
		fmt.Printf("É um círculo com raio %.2f\n", c.Raio)
	} else {
		fmt.Println("Não é um círculo")
	}

	// Type switch
	fmt.Println("\n=== Type switch ===")
	var items []interface{} = []interface{}{
		42, "texto", true, 3.14, circulo, retangulo,
	}

	for _, item := range items {
		switch v := item.(type) {
		case int:
			fmt.Printf("Inteiro: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %v\n", v)
		case float64:
			fmt.Printf("Float: %.2f\n", v)
		case Forma:
			fmt.Printf("Forma com área: %.2f\n", v.Area())
		default:
			fmt.Printf("Tipo desconhecido: %T\n", v)
		}
	}

	// Interface Stringer (String personalizado)
	fmt.Println("\n=== Interface Stringer ===")
	fmt.Println(circulo)      // Usa o método String()
	fmt.Println(retangulo)    // Usa o método String()

	// Interface Logger
	fmt.Println("\n=== Interface Logger ===")
	consoleLogger := ConsoleLogger{Prefix: "CONSOLE"}
	fileLogger := FileLogger{FileName: "log.txt"}

	processarDados(consoleLogger, "dados do usuário")
	fmt.Println()
	processarDados(fileLogger, "backup de arquivos")

	log.Println("Exemplo de interfaces concluído com sucesso!")
}
