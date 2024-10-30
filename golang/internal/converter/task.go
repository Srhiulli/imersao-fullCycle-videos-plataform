package converter

import (
	"encoding/json"
	"log/slog"
	"time"
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
		panic(err)
	}
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
