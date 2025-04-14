package calc

import (
	"errors"
	"testing"
)

// Teste de unidade básico para a função Soma
func TestSoma(t *testing.T) {
	// Preparação
	a, b := 5, 3
	esperado := 8

	// Execução
	resultado := Soma(a, b)

	// Verificação
	if resultado != esperado {
		t.Errorf("Soma(%d, %d) = %d; esperado %d", a, b, resultado, esperado)
	}
}

// Teste de unidade com múltiplos casos para Subtrai
func TestSubtrai(t *testing.T) {
	// Tabela de casos de teste
	casos := []struct {
		nome     string
		a, b     int
		esperado int
	}{
		{"Positivos", 5, 3, 2},
		{"Negativo", 5, 8, -3},
		{"Zero", 5, 5, 0},
	}

	// Executa cada caso de teste
	for _, caso := range casos {
		t.Run(caso.nome, func(t *testing.T) {
			resultado := Subtrai(caso.a, caso.b)
			if resultado != caso.esperado {
				t.Errorf("Subtrai(%d, %d) = %d; esperado %d", 
					caso.a, caso.b, resultado, caso.esperado)
			}
		})
	}
}

// Teste para Multiplica
func TestMultiplica(t *testing.T) {
	casos := []struct {
		a, b     int
		esperado int
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{0, 3, 0},
	}

	for _, caso := range casos {
		resultado := Multiplica(caso.a, caso.b)
		if resultado != caso.esperado {
			t.Errorf("Multiplica(%d, %d) = %d; esperado %d", 
				caso.a, caso.b, resultado, caso.esperado)
		}
	}
}

// Teste para Divide
func TestDivide(t *testing.T) {
	// Caso de divisão bem-sucedida
	t.Run("Divisão válida", func(t *testing.T) {
		resultado, err := Divide(6, 3)
		if err != nil {
			t.Errorf("Não esperava erro, mas obteve: %v", err)
		}
		if resultado != 2 {
			t.Errorf("Divide(6, 3) = %d; esperado 2", resultado)
		}
	})

	// Caso de divisão por zero
	t.Run("Divisão por zero", func(t *testing.T) {
		_, err := Divide(6, 0)
		if err == nil {
			t.Error("Esperava erro de divisão por zero, mas não obteve erro")
		}
		if !errors.Is(err, ErrDivisaoPorZero) {
			t.Errorf("Esperava ErrDivisaoPorZero, mas obteve: %v", err)
		}
	})
}

// Exemplo de benchmark
func BenchmarkSoma(b *testing.B) {
	// Executa Soma b.N vezes
	for i := 0; i < b.N; i++ {
		Soma(5, 3)
	}
}

// Exemplo de benchmark para Multiplica
func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(5, 3)
	}
}

/*
Para executar os testes:
go test -v

Para executar benchmarks:
go test -bench=.

Para verificar a cobertura de testes:
go test -cover

Para gerar um relatório detalhado de cobertura:
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
*/
