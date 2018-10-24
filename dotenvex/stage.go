package dotenvex

import (
	"os"

	"github.com/joho/godotenv"
)

var stageName = "dev"

var stageNames = []string{"alpha", "sit", "beta", "staging", "sandbox", "prod"}

// SetStage set stage
func SetStage(stage string) {
	if len(stage) > 0 {
		stageName = stage
	}
}

// SetStages set stage
func SetStages(stages []string) {
	stageNames = stages
}

// Stage return stage
func Stage() string {
	return stageName
}

// Stages return stages
func Stages() []string {
	return stageNames
}

func init() {
	SetStage(os.Getenv("STAGE"))
}

// LoadByStage load by stage
// 1. stage file exists
// 2. stage file does not exist, stage encrypted file exists
func LoadByStage() {
	// Check stage file exists or not
	stageFile := filePrefix + "." + stageName
	if _, err := os.Stat(stageFile); !os.IsNotExist(err) {
		godotenv.Load(stageFile)
		return
	}

	// Check stage encrypted file exists or not
	stageEncryptedFile := stageFile + encryptedFileExt
	if _, err := os.Stat(stageEncryptedFile); os.IsNotExist(err) {
		return
	}
	// Get password from env
	password := os.Getenv("SECRETS_PASSWORD")
	if len(password) == 0 {
		return
	}
	// Decrypt stage encrypted file
	if err := DecryptFile(stageName, []byte(password)); err == nil {
		godotenv.Load(stageFile)
		return
	}
}
