package dotenvex

import "os"

var stage = "dev"

var stages = []string{"alpha", "sit", "beta", "staging", "sandbox", "prod"}

// SetStage set stage
func SetStage(s string) {
	if len(s) > 0 {
		stage = s
	}
}

// SetStages set stage
func SetStages(ss []string) {
	stages = ss
}

// Stage return stage
func Stage() string {
	return stage
}

// Stages return stages
func Stages() []string {
	return stages
}

func init() {
	SetStage(os.Getenv("STAGE"))
}
