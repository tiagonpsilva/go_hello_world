package calc

// Interface para o repositório de usuários
type UserRepository interface {
	GetUserScore(userID int) (int, error)
	SaveUserScore(userID, score int) error
}

// Serviço de pontuação
type ScoreService struct {
	repo UserRepository
}

// Construtor para o serviço
func NewScoreService(repo UserRepository) *ScoreService {
	return &ScoreService{repo: repo}
}

// Método para adicionar pontos à pontuação de um usuário
func (s *ScoreService) AddPoints(userID, points int) (int, error) {
	// Obter a pontuação atual
	currentScore, err := s.repo.GetUserScore(userID)
	if err != nil {
		return 0, err
	}

	// Calcular a nova pontuação
	newScore := Soma(currentScore, points)

	// Salvar a nova pontuação
	err = s.repo.SaveUserScore(userID, newScore)
	if err != nil {
		return 0, err
	}

	return newScore, nil
}
