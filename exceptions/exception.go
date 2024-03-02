package exceptions

import (
	"encoding/json"
	"github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/honestbank/tech-assignment-backend-engineer/model"
	"log"
	"net"
	"net/http"
)

func HandleError(err error, statusCode int) (string, int, error) {
	if err != nil {
		return err.Error(), statusCode, err
	}
	return "", http.StatusOK, nil
}

func HandleStatusCodeNot200Error(resp *http.Response, err error) (string, int, error) {
	if resp.StatusCode != http.StatusOK {
		return "unexpected error code", resp.StatusCode, err
	}
	return "", http.StatusOK, nil
}

func HandleRedisServerError(err error) {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		log.Println(constants.LOG_LEVEL_ERROR + "Timeout connecting to Redis server")
	} else {
		log.Println(constants.LOG_LEVEL_WARN + err.Error())
	}
	log.Print(constants.LOG_LEVEL_INFO + "Trying to fetch data from db")
}

func HandleValidationOrBadJsonInputError(err error, resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(resp).Encode(model.JsonError{Success: false, Message: err.Error()})
}

func HandleDecisionServiceError(err error, statusCode int, resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	json.NewEncoder(resp).Encode(model.JsonError{Success: false, Message: err.Error()})
}

func HandleOSOpenFileError(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
		return
	}
}

func HandleFilePathError(exist bool, msg string) {
	if !exist {
		log.Fatal(msg)
	}
}

func HandleFileReadingError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
