package main

import "fmt"

func main() {
	// Declaração de um mapa vazio
	var notas map[string]float64
	fmt.Println("Mapa não inicializado:", notas)

	// Inicializando um mapa usando make
	notas = make(map[string]float64)
	fmt.Println("Mapa vazio inicializado:", notas)

	// Adicionando elementos ao mapa
	notas["João"] = 8.5
	notas["Maria"] = 9.2
	notas["Pedro"] = 7.3
	fmt.Println("Mapa após inserção:", notas)

	// Inicialização direta
	idades := map[string]int{
		"Ana":    25,
		"Carlos": 30,
		"Lucia":  28,
	}
	fmt.Println("Mapa de idades:", idades)

	// Acessando valores
	notaJoao := notas["João"]
	fmt.Println("Nota do João:", notaJoao)

	// Tentando acessar uma chave que não existe
	notaFernando, existe := notas["Fernando"]
	if existe {
		fmt.Println("Nota do Fernando:", notaFernando)
	} else {
		fmt.Println("Fernando não está no mapa")
	}

	// Alterando um valor
	notas["Pedro"] = 8.0
	fmt.Println("Mapa após alteração da nota de Pedro:", notas)

	// Excluindo um elemento
	delete(notas, "Maria")
	fmt.Println("Mapa após excluir Maria:", notas)

	// Verificando se uma chave existe
	_, existeJoao := notas["João"]
	fmt.Println("João existe no mapa?", existeJoao)

	// Iterando sobre um mapa (ordem não garantida)
	fmt.Println("\nIterando sobre o mapa de idades:")
	for nome, idade := range idades {
		fmt.Printf("%s tem %d anos\n", nome, idade)
	}

	// Iterando apenas pelas chaves
	fmt.Println("\nNomes no mapa de idades:")
	for nome := range idades {
		fmt.Println("-", nome)
	}

	// Iterando apenas pelos valores
	fmt.Println("\nIdades no mapa:")
	for _, idade := range idades {
		fmt.Println("-", idade)
	}

	// Tamanho do mapa
	fmt.Println("\nNúmero de elementos no mapa de notas:", len(notas))

	// Mapas aninhados
	usuarios := map[string]map[string]string{
		"user1": {
			"nome":  "João Silva",
			"email": "joao@exemplo.com",
			"cargo": "Desenvolvedor",
		},
		"user2": {
			"nome":  "Maria Souza",
			"email": "maria@exemplo.com",
			"cargo": "Gerente",
		},
	}

	fmt.Println("\nDados dos usuários:")
	for id, info := range usuarios {
		fmt.Printf("ID: %s\n", id)
		fmt.Printf("  Nome: %s\n", info["nome"])
		fmt.Printf("  Email: %s\n", info["email"])
		fmt.Printf("  Cargo: %s\n", info["cargo"])
	}

	// Adicionando um novo campo a um usuário existente
	if user, ok := usuarios["user1"]; ok {
		user["telefone"] = "(11) 99999-9999"
	}

	// Verificando o campo adicionado
	fmt.Println("\nTelefone do user1:", usuarios["user1"]["telefone"])

	// Mapas como conjuntos (sets)
	fmt.Println("\nUsando mapa como conjunto (set):")
	vogais := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	letra := 'e'
	if vogais[letra] {
		fmt.Printf("'%c' é uma vogal\n", letra)
	} else {
		fmt.Printf("'%c' não é uma vogal\n", letra)
	}

	letra = 'b'
	if vogais[letra] {
		fmt.Printf("'%c' é uma vogal\n", letra)
	} else {
		fmt.Printf("'%c' não é uma vogal\n", letra)
	}
}
