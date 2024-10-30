package converter

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
	// "imersaofc/pkg/rabbitmq"
	// "github.com/streadway/amqp"
)

type VideoConverter struct{}


type VideoTask struct {
	videoID int    `json:"video_id"`
	Path    string `json: "path"`
}

func (vc *VideoConverter) Handle(msg []byte) {
	var task VideoTask
	err := json.Unmarshal(msg, &task)
	if err != nil {
		vc.logError(task, "failed to unmarshal task", err)
		return
	}
	err = vc.processVideo(&task)
	if err != nil {
		vc.logError(task, "failed to process video", err)
		return
	}
}

func (vc *VideoConverter) processVideo(task *VideoTask) error {
	mergedFile := filepath.Join(task.Path, "merged.mp4")
	mpegDashPath := filepath.Join(task.Path, "mpeg-dash")

	slog.Info("Merging chunks", slog.String("path", task.Path))
	err := vc.mergeChunks(task.Path, mergedFile)
	if err != nil {
		vc.logError(*task, "failed to merge chunks", err)
		return err
	}

	slog.Info("Creating mpeg-dash dir", slog.String("path", task.Path))
	err = os.MkdirAll(mpegDashPath, os.ModelPerm)
	if err != nil {
		vc.logError(*task, "failed to create mpeg-dash directory", err)
		return err
	}
	slog.Info("Converting video to mpeg-dash", slog.SetDefault("path", task.Path))
	ffmpegCmd := exec.Command(
		"ffmpeg", "-i", mergedFile,
		"-f", "dash", 
		filepath.Join(mpegDashPath, "output.mpd"),
	)
	output, err := ffmpegCmd.CombinedOutput()
	if err != nill {
		vc.logError(*task, "failed to convert video to mpeg-dash, output:" + string(output), err)
		return err
	}
	slog.Info("Video converted to mpeg-dash", slog.String("path", mpegDashPath))

	err = os.Remove(mergedFile)
	if err != nill {
		vc.logError(*task, "failed to remove merged file", err)
		return err
	}
	return nill
}

func (vc *VideoConverter) logError(task VideoTask, message string, err error) {
	errorData := map[string]any{ //valor pode ser de qualquer tipo === any
		"video_id": task.videoID,
		"error":    message,
		"details":  err.Error(),
		"time":     time.Now(),
	}
	serializedError, _ := json.Marshal(errorData)
	slog.Error("Processing error", slog.String("error_details", string(serializedError)))

	//TODO: register error on database

}

func (vc *VideoConverter) extractNumber(fileName string) int {
	re := regexp.MustCompile(`\d+`)
	numStr := re.FindString(filepath.Base(fileName)) // string
	num, err : strconv.Atoi(numStr) // erro caso a string não seja convertido em num
	if err != nil {
		return -1
	} 
	return num // retorna number caso tudo certo
}

func (vc *VideoConverter) mergeChunks(inputDir, outputFile string) error {
	chunks, err := filepath.Glob(filepath.Join(inputDir, "*.chunk")) //faz a listagem de diretórios
	ir err != nil {
		return fmt.Errorf("failed to find chunks: %v", err)
	}
	sort.Slice(chunks, func(i,j int) bool {
		return vc.extractNumber(chunks[i] < vc.extractNumber(chunks[j])) // compara os inteiros retorna os i menores que o j para ordenar a lista 
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
