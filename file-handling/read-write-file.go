package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Escrever texto em um arquivo
func escreverArquivo(nomeArquivo, conteudo string) error {
	log.Printf("Tentando escrever no arquivo: %s", nomeArquivo)
	
	// Cria o arquivo (ou trunca se já existir)
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		log.Printf("Erro ao criar arquivo: %v", err)
		return err
	}
	defer arquivo.Close() // Garante que o arquivo será fechado ao sair da função
	
	// Escreve o conteúdo
	_, err = arquivo.WriteString(conteudo)
	if err != nil {
		log.Printf("Erro ao escrever no arquivo: %v", err)
		return err
	}
	
	log.Printf("Arquivo escrito com sucesso: %s", nomeArquivo)
	return nil
}

// Ler todo o conteúdo de um arquivo de uma vez
func lerArquivoCompleto(nomeArquivo string) (string, error) {
	log.Printf("Tentando ler arquivo completo: %s", nomeArquivo)
	
	// Lê todo o arquivo
	conteudoBytes, err := ioutil.ReadFile(nomeArquivo)
	if err != nil {
		log.Printf("Erro ao ler arquivo: %v", err)
		return "", err
	}
	
	log.Printf("Arquivo lido com sucesso: %s (%d bytes)", nomeArquivo, len(conteudoBytes))
	return string(conteudoBytes), nil
}

// Ler arquivo linha por linha
func lerArquivoPorLinhas(nomeArquivo string) error {
	log.Printf("Tentando ler arquivo por linhas: %s", nomeArquivo)
	
	// Abre o arquivo
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		log.Printf("Erro ao abrir arquivo: %v", err)
		return err
	}
	defer arquivo.Close()
	
	// Cria um scanner para ler linha por linha
	scanner := bufio.NewScanner(arquivo)
	
	// Contador de linhas
	linha := 1
	
	// Lê linha por linha
	for scanner.Scan() {
		texto := scanner.Text()
		fmt.Printf("Linha %d: %s\n", linha, texto)
		linha++
	}
	
	// Verifica se houve erro durante a leitura
	if err := scanner.Err(); err != nil {
		log.Printf("Erro ao ler linhas: %v", err)
		return err
	}
	
	log.Printf("Arquivo lido por linhas com sucesso: %s (%d linhas)", nomeArquivo, linha-1)
	return nil
}

// Adicionar conteúdo a um arquivo existente (append)
func adicionarAoArquivo(nomeArquivo, conteudo string) error {
	log.Printf("Tentando adicionar conteúdo ao arquivo: %s", nomeArquivo)
	
	// Abre o arquivo para append (O_APPEND) e cria se não existir (O_CREATE)
	arquivo, err := os.OpenFile(nomeArquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Erro ao abrir arquivo para append: %v", err)
		return err
	}
	defer arquivo.Close()
	
	// Adiciona uma nova linha se necessário
	if conteudo != "" && !strings.HasPrefix(conteudo, "\n") {
		conteudo = "\n" + conteudo
	}
	
	// Escreve o conteúdo
	if _, err := arquivo.WriteString(conteudo); err != nil {
		log.Printf("Erro ao adicionar conteúdo: %v", err)
		return err
	}
	
	log.Printf("Conteúdo adicionado com sucesso ao arquivo: %s", nomeArquivo)
	return nil
}

// Copiar arquivo
func copiarArquivo(origem, destino string) error {
	log.Printf("Tentando copiar arquivo de %s para %s", origem, destino)
	
	// Cria diretórios de destino se não existirem
	dirDestino := filepath.Dir(destino)
	if err := os.MkdirAll(dirDestino, 0755); err != nil {
		log.Printf("Erro ao criar diretórios: %v", err)
		return err
	}
	
	// Abre o arquivo de origem
	arquivoOrigem, err := os.Open(origem)
	if err != nil {
		log.Printf("Erro ao abrir arquivo de origem: %v", err)
		return err
	}
	defer arquivoOrigem.Close()
	
	// Cria o arquivo de destino
	arquivoDestino, err := os.Create(destino)
	if err != nil {
		log.Printf("Erro ao criar arquivo de destino: %v", err)
		return err
	}
	defer arquivoDestino.Close()
	
	// Copia os dados
	bytes, err := io.Copy(arquivoDestino, arquivoOrigem)
	if err != nil {
		log.Printf("Erro ao copiar conteúdo: %v", err)
		return err
	}
	
	log.Printf("Arquivo copiado com sucesso: %s -> %s (%d bytes)", origem, destino, bytes)
	return nil
}

// Verificar se um arquivo existe
func arquivoExiste(nomeArquivo string) bool {
	_, err := os.Stat(nomeArquivo)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// Listar arquivos em um diretório
func listarArquivos(diretorio string) ([]string, error) {
	log.Printf("Listando arquivos no diretório: %s", diretorio)
	
	// Verifica se o diretório existe
	info, err := os.Stat(diretorio)
	if os.IsNotExist(err) {
		log.Printf("Diretório não existe: %s", diretorio)
		return nil, fmt.Errorf("diretório não existe: %s", diretorio)
	}
	
	// Verifica se é realmente um diretório
	if !info.IsDir() {
		log.Printf("O caminho não é um diretório: %s", diretorio)
		return nil, fmt.Errorf("o caminho não é um diretório: %s", diretorio)
	}
	
	// Lista os arquivos
	arquivos, err := ioutil.ReadDir(diretorio)
	if err != nil {
		log.Printf("Erro ao listar diretório: %v", err)
		return nil, err
	}
	
	// Cria uma lista com os nomes dos arquivos
	var nomes []string
	for _, arquivo := range arquivos {
		nomes = append(nomes, arquivo.Name())
	}
	
	log.Printf("Listagem concluída: %d arquivos encontrados", len(nomes))
	return nomes, nil
}

func main() {
	// Configurando log
	log.SetPrefix("[FILE-IO] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Iniciando exemplo de manipulação de arquivos...")
	
	// Nome do arquivo para os testes
	nomeArquivo := "exemplo.txt"
	
	// Verifica se o arquivo já existe
	if arquivoExiste(nomeArquivo) {
		log.Printf("Arquivo já existe: %s", nomeArquivo)
		// Excluir para começar do zero
		if err := os.Remove(nomeArquivo); err != nil {
			log.Fatalf("Erro ao remover arquivo existente: %v", err)
		}
		log.Printf("Arquivo removido: %s", nomeArquivo)
	}
	
	// Escrevendo em um arquivo
	conteudo := "Linha 1: Hello, World!\nLinha 2: Manipulação de arquivos em Go.\nLinha 3: Exemplo prático."
	if err := escreverArquivo(nomeArquivo, conteudo); err != nil {
		log.Fatalf("Erro ao escrever arquivo: %v", err)
	}
	
	// Lendo todo o conteúdo do arquivo
	conteudoLido, err := lerArquivoCompleto(nomeArquivo)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo: %v", err)
	}
	fmt.Println("\nConteúdo do arquivo:")
	fmt.Println(conteudoLido)
	
	// Lendo o arquivo linha por linha
	fmt.Println("\nLendo linha por linha:")
	if err := lerArquivoPorLinhas(nomeArquivo); err != nil {
		log.Fatalf("Erro ao ler arquivo por linhas: %v", err)
	}
	
	// Adicionando conteúdo ao arquivo
	if err := adicionarAoArquivo(nomeArquivo, "Linha 4: Conteúdo adicionado posteriormente."); err != nil {
		log.Fatalf("Erro ao adicionar conteúdo: %v", err)
	}
	
	// Lendo o arquivo atualizado
	conteudoAtualizado, err := lerArquivoCompleto(nomeArquivo)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo atualizado: %v", err)
	}
	fmt.Println("\nConteúdo atualizado:")
	fmt.Println(conteudoAtualizado)
	
	// Copiando o arquivo
	nomeCopia := "copia_exemplo.txt"
	if err := copiarArquivo(nomeArquivo, nomeCopia); err != nil {
		log.Fatalf("Erro ao copiar arquivo: %v", err)
	}
	
	// Verificando a cópia
	copiaCont, err := lerArquivoCompleto(nomeCopia)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo copiado: %v", err)
	}
	fmt.Println("\nConteúdo da cópia:")
	fmt.Println(copiaCont)
	
	// Criando um diretório temporário
	dirTemp := "temp_dir"
	if err := os.MkdirAll(dirTemp, 0755); err != nil {
		log.Fatalf("Erro ao criar diretório: %v", err)
	}
	log.Printf("Diretório criado: %s", dirTemp)
	
	// Criando alguns arquivos no diretório
	for i := 1; i <= 3; i++ {
		nome := filepath.Join(dirTemp, fmt.Sprintf("arquivo%d.txt", i))
		if err := escreverArquivo(nome, fmt.Sprintf("Conteúdo do arquivo %d", i)); err != nil {
			log.Fatalf("Erro ao criar arquivo %s: %v", nome, err)
		}
	}
	
	// Listando os arquivos
	arquivos, err := listarArquivos(dirTemp)
	if err != nil {
		log.Fatalf("Erro ao listar arquivos: %v", err)
	}
	fmt.Println("\nArquivos no diretório:")
	for i, arquivo := range arquivos {
		fmt.Printf("%d. %s\n", i+1, arquivo)
	}
	
	// Limpeza
	log.Println("Realizando limpeza...")
	
	// Removendo arquivos
	os.Remove(nomeArquivo)
	os.Remove(nomeCopia)
	
	// Removendo diretório e conteúdo
	os.RemoveAll(dirTemp)
	
	log.Println("Exemplo de manipulação de arquivos concluído com sucesso!")
}
