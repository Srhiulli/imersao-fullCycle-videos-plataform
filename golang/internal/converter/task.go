package converter

import "encoding/json"

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
