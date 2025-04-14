# Go Exemplos

Coleção de exemplos de código em Go para casos de uso comuns.

## Descrição

Este repositório contém vários exemplos de código em Go demonstrando casos de uso comuns, desde o básico "Hello World" até operações mais avançadas como manipulação de arquivos, concorrência, HTTP e manipulação de JSON.

## Pré-requisitos

- [Go](https://golang.org/doc/install) instalado (versão 1.18+ recomendada)

## Como executar os exemplos

1. Clone este repositório:
   ```bash
   git clone https://github.com/tiagonpsilva/go_hello_world.git
   cd go_hello_world
   ```

2. Execute um exemplo específico:
   ```bash
   go run basic/hello.go
   ```

3. Ou você pode compilar e executar:
   ```bash
   go build -o hello basic/hello.go
   ./hello
   ```

## Estrutura do Projeto

- `basic/`: Conceitos básicos da linguagem
  - `hello.go`: Exemplo clássico Hello World
  - `variables.go`: Declaração e tipos de variáveis 
  - `control-flow.go`: Estruturas de controle (if, for, switch)
  - `functions.go`: Funções e retornos múltiplos

- `data-structures/`: Estruturas de dados comuns
  - `arrays.go`: Arrays e slices
  - `maps.go`: Mapas (dicionários)
  - `structs.go`: Estruturas (structs)
  - `interfaces.go`: Interfaces e polimorfismo

- `concurrency/`: Concorrência e paralelismo
  - `goroutines.go`: Goroutines básicas
  - `channels.go`: Comunicação via channels
  - `waitgroups.go`: Sincronização com WaitGroups
  - `mutex.go`: Exclusão mútua com Mutex

- `file-handling/`: Manipulação de arquivos
  - `read-file.go`: Leitura de arquivos
  - `write-file.go`: Escrita em arquivos
  - `json-handling.go`: Codificação e decodificação JSON

- `web/`: Programação web
  - `simple-server.go`: Servidor HTTP simples
  - `rest-client.go`: Cliente REST

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir um pull request com novos exemplos ou melhorias.

## Autor

- tiagonpsilva
