package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Exemplo básico de logging em Go usando a biblioteca padrão log
// Em projetos reais, considere usar bibliotecas como logrus, zap ou zerolog

func main() {
	// Log básico usando a biblioteca padrão
	log.Println("Iniciando aplicação")
	log.Printf("Versão %d.%d.%d\n", 1, 0, 0)

	// Log com prefixo
	log.SetPrefix("APP: ")
	log.Println("Log com prefixo personalizado")

	// Log com carimbo de data/hora e arquivo/linha
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Log com carimbo de data/hora e informação de localização")

	// Erro fatal (chama os.Exit(1) após o log)
	// log.Fatal("Erro fatal encontrado")

	// Panic (chama panic() após o log)
	// log.Panic("Situação de pânico encontrada")

	// Logging para arquivo
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo de log:", err)
	}
	defer logFile.Close()

	// Criando um logger customizado para arquivo
	fileLogger := log.New(logFile, "FILE_LOG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	fileLogger.Println("Este log vai para o arquivo app.log")

	// Logger com múltiplos outputs
	multiWriter := os.Stdout
	multiLogger := log.New(multiWriter, "MULTI: ", log.LstdFlags)
	multiLogger.Println("Este log vai para múltiplos destinos")

	// Exemplo de níveis de log
	info("Operação concluída com sucesso")
	warning("Recurso está com capacidade reduzida")
	error("Erro ao conectar ao banco de dados")

	// Exemplo de log em estrutura
	logEvent("USER_LOGIN", map[string]interface{}{
		"user_id":   12345,
		"timestamp": time.Now().Unix(),
		"ip":        "192.168.1.1",
		"success":   true,
	})

	log.Println("Aplicação finalizada")
}

// Log info
func info(message string) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println(message)
}

// Log warning
func warning(message string) {
	logger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	logger.Println(message)
}

// Log error
func error(message string) {
	logger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	logger.Println(message)
}

// Log estruturado
func logEvent(eventType string, data map[string]interface{}) {
	logger := log.New(os.Stdout, fmt.Sprintf("EVENT [%s]: ", eventType), log.Ldate|log.Ltime)
	
	// Formata os dados como uma string
	dataStr := "{"
	for k, v := range data {
		dataStr += fmt.Sprintf("%s: %v, ", k, v)
	}
	// Remove a última vírgula e fecha a chave
	if len(data) > 0 {
		dataStr = dataStr[:len(dataStr)-2]
	}
	dataStr += "}"
	
	logger.Println(dataStr)
}

/*
NOTA: Para projetos reais, considere usar bibliotecas de logging mais avançadas:

1. Logrus (github.com/sirupsen/logrus)
   - Suporte a estruturas de dados
   - Níveis de log
   - Hooks para integração com sistemas externos

2. Zap (go.uber.org/zap)
   - Alta performance
   - Log estruturado
   - Otimizado para produção

3. Zerolog (github.com/rs/zerolog)
   - API para logs JSON
   - Baixa alocação de memória
   - Alta performance
*/
