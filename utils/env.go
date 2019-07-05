package utils

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
)

// LoadEnv variables from .env.* file
func LoadEnv(env string) {
	envFile := fmt.Sprintf(".env.%s", env)

	WriteInfoLog("Try to load " + envFile)

	err := godotenv.Load(envFile)

	if err != nil {
		CheckErr(err)
	}
}

// FillStructFromConfig takes environment variable (json config file) by key,
// load config and fill struct.
func FillStructFromConfig(key string, target interface{}) {
	file, err := os.Open(os.Getenv(key))
	defer file.Close()

	CheckErr(err)

	jsonData, err := ioutil.ReadAll(file)

	CheckErr(err)

	err = json.Unmarshal(jsonData, target)

	if err != nil {
		panic(err)
	}
}
