package app

import (
	"backend-golang-evermos/src/helper"
	"github.com/joho/godotenv"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

// Get is function get key environment
func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

// New is function get environment from file .env
func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	helper.PanicIfError(err)
	return &configImpl{}
}
