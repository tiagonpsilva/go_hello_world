package main

import (
	"fmt"
	"time"
)

// Definição de uma struct básica
type Pessoa struct {
	Nome      string
	Idade     int
	Endereco  string
	CreatedAt time.Time
}

// Struct com tags (para JSON, etc.)
type Produto struct {
	ID          int     `json:"id"`
	Nome        string  `json:"nome"`
	Preco       float64 `json:"preco"`
	Disponivel  bool    `json:"disponivel"`
	Categorias  []string `json:"categorias,omitempty"`
}

// Método de uma struct (receiver function)
func (p Pessoa) Saudacao() string {
	return fmt.Sprintf("Olá, meu nome é %s e tenho %d anos", p.Nome, p.Idade)
}

// Método com ponteiro receiver (modifica a struct original)
func (p *Pessoa) FazerAniversario() {
	p.Idade++
}

// Struct aninhada
type Endereco struct {
	Rua     string
	Numero  int
	Cidade  string
	Estado  string
	CEP     string
}

type Cliente struct {
	ID       int
	Nome     string
	Email    string
	Endereco Endereco  // Struct aninhada
	Desde    time.Time
}

// Struct com composição (embedded struct)
type Funcionario struct {
	Pessoa    // Embedding da struct Pessoa
	Cargo     string
	Salario   float64
}

func main() {
	// Inicialização de struct
	p1 := Pessoa{
		Nome:      "João Silva",
		Idade:     30,
		Endereco:  "Rua A, 123",
		CreatedAt: time.Now(),
	}
	fmt.Println("Pessoa 1:", p1)

	// Inicialização pela ordem dos campos (não recomendado)
	p2 := Pessoa{"Maria Souza", 25, "Av B, 456", time.Now()}
	fmt.Println("Pessoa 2:", p2)

	// Struct parcialmente inicializada
	var p3 Pessoa
	p3.Nome = "Pedro Santos"
	p3.Idade = 40
	fmt.Println("Pessoa 3:", p3)

	// Usando métodos de struct
	fmt.Println(p1.Saudacao())
	
	// Usando método que modifica a struct (com ponteiro)
	fmt.Printf("Idade antes: %d\n", p1.Idade)
	p1.FazerAniversario()
	fmt.Printf("Idade depois: %d\n", p1.Idade)

	// Struct aninhada
	cliente := Cliente{
		ID:    1,
		Nome:  "Ana Costa",
		Email: "ana@exemplo.com",
		Endereco: Endereco{
			Rua:    "Rua das Flores",
			Numero: 42,
			Cidade: "São Paulo",
			Estado: "SP",
			CEP:    "01234-567",
		},
		Desde: time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
	}
	fmt.Println("\nCliente:", cliente)
	fmt.Println("Endereço:", cliente.Endereco)
	fmt.Println("CEP:", cliente.Endereco.CEP)

	// Usando composição (embedded struct)
	funcionario := Funcionario{
		Pessoa: Pessoa{
			Nome:  "Carlos Gomes",
			Idade: 35,
		},
		Cargo:   "Desenvolvedor",
		Salario: 5000.0,
	}
	fmt.Println("\nFuncionário:", funcionario)
	fmt.Println("Nome do funcionário:", funcionario.Nome)  // Acesso direto por causa do embedding
	fmt.Println("Idade do funcionário:", funcionario.Idade)

	// Chamando método da struct embedded
	fmt.Println(funcionario.Saudacao())
	
	// Struct como um mapa de campos
	type Dinamico struct {
		Campos map[string]interface{}
	}
	
	d := Dinamico{
		Campos: map[string]interface{}{
			"nome":   "Teste",
			"valor":  123,
			"ativo":  true,
			"tags":   []string{"a", "b", "c"},
		},
	}
	fmt.Println("\nStruct dinâmica:", d)
	fmt.Println("Campo nome:", d.Campos["nome"])
	fmt.Println("Campo valor:", d.Campos["valor"])
	
	// Slice de structs
	pessoas := []Pessoa{
		{Nome: "Alice", Idade: 32},
		{Nome: "Bob", Idade: 28},
		{Nome: "Carol", Idade: 45},
	}
	
	fmt.Println("\nLista de pessoas:")
	for i, p := range pessoas {
		fmt.Printf("%d: %s tem %d anos\n", i+1, p.Nome, p.Idade)
	}
}
