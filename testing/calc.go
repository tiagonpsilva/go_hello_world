package calc

// Função para somar dois números
func Soma(a, b int) int {
	return a + b
}

// Função para subtrair
func Subtrai(a, b int) int {
	return a - b
}

// Função para multiplicar
func Multiplica(a, b int) int {
	return a * b
}

// Função para dividir
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisaoPorZero
	}
	return a / b, nil
}

// Erro customizado para divisão por zero
type ErrDiv0 struct{}

func (e ErrDiv0) Error() string {
	return "divisão por zero não é permitida"
}

// Variável para o erro
var ErrDivisaoPorZero error = ErrDiv0{}
