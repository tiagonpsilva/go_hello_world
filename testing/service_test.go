package calc

import (
	"errors"
	"testing"
)

// Mock do repositório para testes
type MockUserRepository struct {
	userScores       map[int]int
	getScoreError    error
	saveScoreError   error
	getScoreCalled   bool
	saveScoreCalled  bool
	lastUserID       int
	lastSavedScore   int
}

// Implementação do método GetUserScore para o mock
func (m *MockUserRepository) GetUserScore(userID int) (int, error) {
	m.getScoreCalled = true
	m.lastUserID = userID
	
	if m.getScoreError != nil {
		return 0, m.getScoreError
	}
	
	score, exists := m.userScores[userID]
	if !exists {
		return 0, errors.New("usuário não encontrado")
	}
	
	return score, nil
}

// Implementação do método SaveUserScore para o mock
func (m *MockUserRepository) SaveUserScore(userID, score int) error {
	m.saveScoreCalled = true
	m.lastUserID = userID
	m.lastSavedScore = score
	
	if m.saveScoreError != nil {
		return m.saveScoreError
	}
	
	m.userScores[userID] = score
	return nil
}

// Cria um novo mock para testes
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		userScores: make(map[int]int),
	}
}

// Teste do ScoreService com mock
func TestScoreService_AddPoints(t *testing.T) {
	// Caso de teste: adição de pontos bem-sucedida
	t.Run("Adiciona pontos com sucesso", func(t *testing.T) {
		// Preparação
		mockRepo := NewMockUserRepository()
		mockRepo.userScores[123] = 100  // Usuário 123 tem 100 pontos inicialmente
		
		service := NewScoreService(mockRepo)
		
		// Execução
		newScore, err := service.AddPoints(123, 50)
		
		// Verificação
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		
		if newScore != 150 {
			t.Errorf("Pontuação esperada: 150, obtida: %d", newScore)
		}
		
		// Verifica se os métodos foram chamados
		if !mockRepo.getScoreCalled {
			t.Error("GetUserScore não foi chamado")
		}
		
		if !mockRepo.saveScoreCalled {
			t.Error("SaveUserScore não foi chamado")
		}
		
		// Verifica se o valor correto foi salvo
		if mockRepo.lastSavedScore != 150 {
			t.Errorf("Valor salvo esperado: 150, obtido: %d", mockRepo.lastSavedScore)
		}
	})
	
	// Caso de teste: erro ao obter pontuação
	t.Run("Erro ao obter pontuação", func(t *testing.T) {
		// Preparação
		mockRepo := NewMockUserRepository()
		mockRepo.getScoreError = errors.New("erro de banco de dados")
		
		service := NewScoreService(mockRepo)
		
		// Execução
		_, err := service.AddPoints(123, 50)
		
		// Verificação
		if err == nil {
			t.Error("Esperava um erro, mas não ocorreu")
		}
		
		// Verifica se SaveUserScore não foi chamado após o erro
		if mockRepo.saveScoreCalled {
			t.Error("SaveUserScore não deveria ter sido chamado após erro em GetUserScore")
		}
	})
	
	// Caso de teste: erro ao salvar pontuação
	t.Run("Erro ao salvar pontuação", func(t *testing.T) {
		// Preparação
		mockRepo := NewMockUserRepository()
		mockRepo.userScores[123] = 100
		mockRepo.saveScoreError = errors.New("erro ao salvar")
		
		service := NewScoreService(mockRepo)
		
		// Execução
		_, err := service.AddPoints(123, 50)
		
		// Verificação
		if err == nil {
			t.Error("Esperava um erro, mas não ocorreu")
		}
		
		// Mesmo com erro no save, a pontuação não deve ter sido atualizada no repositório
		if mockRepo.userScores[123] != 100 {
			t.Errorf("A pontuação não deveria ter sido alterada no repositório, valor: %d", mockRepo.userScores[123])
		}
	})
}

/*
Este exemplo demonstra como criar mocks para testes unitários em Go.
Em projetos reais, considere usar bibliotecas de mocking como:

1. gomock (github.com/golang/mock)
   - Geração automática de mocks a partir de interfaces
   - API para configurar comportamentos e expectativas

2. testify (github.com/stretchr/testify)
   - Suíte de ferramentas para testes em Go
   - Mock objects, assertions, suites, etc.

3. go-sqlmock (github.com/DATA-DOG/go-sqlmock)
   - Mock específico para SQL/banco de dados
   - Ideal para testar código que utiliza database/sql

Para executar os testes:
go test -v
*/
