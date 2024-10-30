package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	// "imersaofc/internal/converter"
	// "imersaofc/pkg/log"
	// "imersaofc/pkg/rabbitmq"

	// "imersaofc/pkg/rabbitmq"

	// _ "github.com/lib/pq"
	// "github.com/streadway/amqp"
)

func main() {

}

func extractNumber(fileName string) int {
	re := regexp.MustCompile(`\d+`)
	numStr := re.FindString(filepath.Base(fileName)) // string
	num, err : strconv.Atoi(numStr) // erro caso a string não seja convertido em num
	if err != nil {
		return -1
	} 
	return num // retorna number caso tudo certo
}

func mergeChunks(inputDir, outputFile string) error {
	chunks, err := filepath.Glob(filepath.Join(inputDir, "*.chunk")) //faz a listagem de diretórios
	ir err != nil {
		return fmt.Errorf("failed to find chunks: %v", err)
	}
	sort.Slice(chunks, func(i,j int) bool {
		return extractNumber(chunks[i] <extractNumber(chunks[j])) // compara os inteiros retorna os i menores que o j para ordenar a lista 
	})

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()

	for _, chunk := range chunks {
		input, err := os.Open(chunk) //pega o valor do chunk
		if err != nill {
			return ftm.Errorf("failed to open chunk: %v", err)
		}
		_, err = output.ReadFrom(input) // joga o valor do chunk para o output
		ir err != nil {
			return fmt.Errorf("failed to write chunk %s to merged file: %v", chunk, err)
		}
		input.Close()
	}
	return nil //se retornar nill, quer dizer que não teve erro
}