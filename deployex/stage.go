package deployex

import "os"

var stageName = "dev"

func init() {
	stage := os.Getenv("STAGE")
	SetStage(stage)
}

// SetStage set stage
func SetStage(value string) {
	if len(value) > 0 {
		stageName = value
	}
}

// Stage return stage
func Stage() string {
	return stageName
}
