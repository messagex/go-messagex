package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env struct to be used as the logger object
type Env struct {
}

//CreateEnv - Loads .env file
func CreateEnv() (*Env, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	return &Env{}, nil
}

//CreateEnvFromFile - Loads a specified file
func CreateEnvFromFile(filename string) (*Env, error) {

	err := godotenv.Load(filename)
	if err != nil {
		log.Printf("Error loading file %s: %v", filename, err)
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	return &Env{}, nil
}

// Get - Return a value from the environment
func (e *Env) Get(key string) string {
	val := os.Getenv(key)
	return val
}

// Set - Sets an environment variable value
func (e *Env) Set(key, val string) {
	os.Setenv(key, val)
}

// GetMissing - returns a list of missing env keys
func (e *Env) GetMissing(keys []string) []string {

	var missingValues []string

	for _, k := range keys {
		v := os.Getenv(k)
		if v == "" {
			missingValues = append(missingValues, k)
		}
	}

	return missingValues
}

// CheckEnv - accepts a list of keys to check for existense
func (e *Env) CheckEnv(keys []string) error {

	for _, k := range keys {
		v := os.Getenv(k)
		if v == "" {
			return fmt.Errorf("There %s key is missing from the environment", k)
		}
	}

	return nil
}
