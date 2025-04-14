package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	Estoque     int     `json:"estoque"`
	Disponivel  bool    `json:"disponivel"`
	Descricao   string  `json:"descricao,omitempty"` // omitempty: campo será omitido se vazio
	Categorias  []string `json:"categorias"`
}

// Struct aninhada (composição)
type Endereco struct {
	Rua        string
	Numero     int
	Cidade     string
	Estado     string
	CEP        string
}

type Cliente struct {
	ID         int
	Nome       string
	Email      string
	Telefone   string
	Endereco   Endereco // Struct aninhada
}

// Método de uma struct
func (p Pessoa) Saudacao() string {
	return fmt.Sprintf("Olá, meu nome é %s e tenho %d anos", p.Nome, p.Idade)
}

// Método com ponteiro (permite modificar a struct)
func (p *Pessoa) FazerAniversario() {
	p.Idade++
	log.Printf("Aniversário de %s comemorado! Nova idade: %d", p.Nome, p.Idade)
}

// Struct com campo anônimo (embedded struct)
type Funcionario struct {
	Pessoa     // Embedded struct - herda campos e métodos
	Cargo      string
	Salario    float64
}

func main() {
	// Inicializando uma struct básica
	log.Println("Iniciando exemplo de structs...")
	
	pessoa1 := Pessoa{
		Nome:      "João Silva",
		Idade:     30,
		Endereco:  "Rua ABC, 123",
		CreatedAt: time.Now(),
	}
	
	fmt.Println("Pessoa criada:", pessoa1)
	
	// Acessando campos
	fmt.Println("Nome:", pessoa1.Nome)
	fmt.Println("Idade:", pessoa1.Idade)
	
	// Inicialização parcial
	pessoa2 := Pessoa{Nome: "Maria Souza"}
	fmt.Println("\nPessoa com inicialização parcial:", pessoa2)
	
	// Inicialização por ordem (não recomendado - depende da ordem dos campos)
	pessoa3 := Pessoa{"Pedro Santos", 25, "Rua XYZ, 456", time.Now()}
	fmt.Println("\nPessoa com inicialização por ordem:", pessoa3)
	
	// Usando métodos
	saudacao := pessoa1.Saudacao()
	fmt.Println("\nSaudação:", saudacao)
	
	// Método que modifica a struct (usando ponteiro)
	fmt.Printf("\nIdade antes do aniversário: %d\n", pessoa1.Idade)
	pessoa1.FazerAniversario()
	fmt.Printf("Idade após o aniversário: %d\n", pessoa1.Idade)
	
	// Struct aninhada
	cliente := Cliente{
		ID:       1,
		Nome:     "Ana Costa",
		Email:    "ana@email.com",
		Telefone: "(11) 99999-9999",
		Endereco: Endereco{
			Rua:    "Avenida Paulista",
			Numero: 1000,
			Cidade: "São Paulo",
			Estado: "SP",
			CEP:    "01310-100",
		},
	}
	fmt.Printf("\nCliente: %+v\n", cliente)  // %+v mostra nomes dos campos
	fmt.Println("Endereço do cliente:", cliente.Endereco.Rua, cliente.Endereco.Numero)
	
	// Embedded struct (herança de campo/método)
	funcionario := Funcionario{
		Pessoa: Pessoa{
			Nome:     "Carlos Oliveira",
			Idade:    35,
			Endereco: "Rua DEF, 789",
		},
		Cargo:   "Desenvolvedor",
		Salario: 5000.00,
	}
	
	fmt.Printf("\nFuncionário: %+v\n", funcionario)
	
	// Acesso direto a campos embedados
	fmt.Println("Nome do funcionário:", funcionario.Nome)  // Acesso direto ao campo da struct embedada
	fmt.Println("Saudação do funcionário:", funcionario.Saudacao())
	
	// Convertendo struct para JSON
	produto := Produto{
		ID:         1,
		Nome:       "Notebook",
		Preco:      3500.99,
		Estoque:    10,
		Disponivel: true,
		Categorias: []string{"Eletrônicos", "Computadores"},
	}
	
	jsonBytes, err := json.MarshalIndent(produto, "", "  ")
	if err != nil {
		log.Fatalf("Erro ao converter para JSON: %v", err)
	}
	fmt.Println("\nProduto em JSON:")
	fmt.Println(string(jsonBytes))
	
	// JSON para struct
	jsonStr := `{
	  "id": 2,
	  "nome": "Smartphone",
	  "preco": 1999.99,
	  "estoque": 15,
	  "disponivel": true,
	  "categorias": ["Eletrônicos", "Celulares"]
	}`
	
	var novoProduto Produto
	if err := json.Unmarshal([]byte(jsonStr), &novoProduto); err != nil {
		log.Fatalf("Erro ao converter JSON para struct: %v", err)
	}
	
	fmt.Printf("\nProduto deserializado do JSON: %+v\n", novoProduto)
	
	log.Println("Exemplo de structs concluído com sucesso!")
}
