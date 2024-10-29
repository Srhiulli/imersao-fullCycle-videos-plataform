package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
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
}