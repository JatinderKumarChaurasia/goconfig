package filelooker

import (
	"context"
	"log"
	"os"
	"strings"
)

const (
	DefaultAppConfigDir = "config"
)

func init() {
	loadEnvironmentValues()
}

func FileLooker() {}

func loadEnvironmentValues() {
	log.Println("getting default configuration directory")
	var configDir string
	if val, exist := os.LookupEnv("DEFAULT_APP_CONFIG_DIR"); exist && len(strings.TrimSpace(val)) != 0 {
		configDir = val
		log.Println("application configuration directory from environment: ", configDir)
	} else {
		workingDir, err := os.Getwd()
		if err != nil {
			panic("unable to get working directory,exiting!")
		}
		configDir = workingDir + "/" + DefaultAppConfigDir

		log.Println("application configuration directory taken from default: ", configDir)
	}
	context.WithValue(context.Background(), "CONFIG_WORKING_DIR", configDir)
}
