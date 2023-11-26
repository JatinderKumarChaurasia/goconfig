package filelooker

import (
	"log"
	"os"
	"strings"
)

const (
	DEFAULT_APP_CONFIG_DIR = "config"
)

func init() {
	lookFileBasedOnEnvironment()
}
func lookFileBasedOnEnvironment() {
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
		configDir = workingDir + "/" + DEFAULT_APP_CONFIG_DIR
		log.Println("application configuration directory taken from default: ", configDir)
	}
}
