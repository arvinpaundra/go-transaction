package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func GetEnv(key string, fallback string) string {
	// Godotenv read the .env file on the root folder
	a, _ := godotenv.Read()
	var (
		val     string
		isExist bool
	)
	// Check the key of the env using Hashmap
	// if exist return the actual value, if !exist return the fallback value
	val, isExist = a[key]
	if !isExist {
		val = fallback
	}
	return val
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func FormatValidationError(err error) []string {
	var dataErrror []string
	var foo *json.UnmarshalTypeError
	if errors.As(err, &foo) {
		dataErrror = append(dataErrror, err.Error())
		return dataErrror
	}
	for _, e := range err.(validator.ValidationErrors) {
		dataErrror = append(dataErrror, e.Error())
	}

	return dataErrror
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse

}

func CreateErrorLog(errMessage error) {
	fileName := fmt.Sprintf("./storage/error_logs/error-%s.log", time.Now().Format("2006-01-02"))

	// open log file
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	defer logFile.Close()

	// set log out put
	log.SetOutput(logFile)

	log.SetFlags(log.LstdFlags)

	_, fileName, line, _ := runtime.Caller(1)
	log.Printf("[Error] in [%s:%d] %v", fileName, line, errMessage.Error())
}
