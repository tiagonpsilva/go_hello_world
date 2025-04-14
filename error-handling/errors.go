package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

// Erro simples usando errors.New
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}
	return a / b, nil
}

// Erro formatado usando fmt.Errorf
func calcularMedia(numeros []int) (float64, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("impossível calcular média de um slice vazio")
	}
	
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	
	return float64(soma) / float64(len(numeros)), nil
}

// Erro customizado usando struct que implementa a interface error
type ValorNegativoError struct {
	Valor   int
	Mensagem string
}

func (e *ValorNegativoError) Error() string {
	return fmt.Sprintf("%s: %d", e.Mensagem, e.Valor)
}

func calcularRaizQuadrada(n int) (float64, error) {
	if n < 0 {
		return 0, &ValorNegativoError{
			Valor:   n,
			Mensagem: "não é possível calcular raiz quadrada de número negativo",
		}
	}
	
	// Cálculo simplificado para exemplo
	return float64(n * n), nil
}

// Encadeamento (wrapping) de erros
func lerArquivoConfig(caminho string) ([]byte, error) {
	dados, err := os.ReadFile(caminho)
	if err != nil {
		// Encadeia o erro original com contexto adicional
		return nil, fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
	}
	return dados, nil
}

// Verificando tipos de erro
func verificarArquivo(caminho string) error {
	_, err := os.Stat(caminho)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("o arquivo %s não existe", caminho)
		}
		return fmt.Errorf("erro ao verificar arquivo %s: %w", caminho, err)
	}
	return nil
}

// Verificando tipos de erro customizados com errors.As
func processarValor(entrada string) error {
	valor, err := strconv.Atoi(entrada)
	if err != nil {
		return fmt.Errorf("erro de conversão: %w", err)
	}
	
	_, err = calcularRaizQuadrada(valor)
	if err != nil {
		var valorNegativoErr *ValorNegativoError
		if errors.As(err, &valorNegativoErr) {
			return fmt.Errorf("erro de domínio: %w", err)
		}
		return err
	}
	
	return nil
}

// Exemplo com múltiplos erros possíveis
var (
	ErrVazio      = errors.New("string vazia")
	ErrMuitoCurto = errors.New("string muito curta")
	ErrMuitoLongo = errors.New("string muito longa")
)

func validarUsername(username string) error {
	if len(username) == 0 {
		return ErrVazio
	}
	if len(username) < 4 {
		return ErrMuitoCurto
	}
	if len(username) > 20 {
		return ErrMuitoLongo
	}
	return nil
}

func main() {
	fmt.Println("=== Exemplo 1: Erro básico ===")
	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
	
	fmt.Println("\n=== Exemplo 2: Erro formatado ===")
	media, err := calcularMedia([]int{})
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Média:", media)
	}
	
	fmt.Println("\n=== Exemplo 3: Erro customizado ===")
	_, err = calcularRaizQuadrada(-5)
	if err != nil {
		fmt.Println("Erro:", err)
		
		// Verificando tipo do erro usando type assertion
		if valorNegativoErr, ok := err.(*ValorNegativoError); ok {
			fmt.Printf("Valor inválido detectado: %d\n", valorNegativoErr.Valor)
		}
	}
	
	fmt.Println("\n=== Exemplo 4: Encadeamento de erros ===")
	_, err = lerArquivoConfig("config.json")
	if err != nil {
		fmt.Println("Erro:", err)
		
		// Verificando erro original encapsulado
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Printf("Erro de caminho: %v\n", pathErr.Path)
		}
	}
	
	fmt.Println("\n=== Exemplo 5: Verificando tipos de erro ===")
	err = verificarArquivo("arquivo_inexistente.txt")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	
	fmt.Println("\n=== Exemplo 6: Múltiplos erros possíveis ===")
	usernames := []string{"", "abc", "usuario_com_nome_muito_longo_demais"}
	for _, nome := range usernames {
		err := validarUsername(nome)
		fmt.Printf("Validando '%s': ", nome)
		switch err {
		case nil:
			fmt.Println("Username válido!")
		case ErrVazio:
			fmt.Println("Username não pode ser vazio!")
		case ErrMuitoCurto:
			fmt.Println("Username muito curto (mínimo 4 caracteres)!")
		case ErrMuitoLongo:
			fmt.Println("Username muito longo (máximo 20 caracteres)!")
		default:
			fmt.Println("Erro desconhecido:", err)
		}
	}
}
