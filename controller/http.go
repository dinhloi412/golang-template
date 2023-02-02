package controller

import (
	"fmt"
	"net/http"

	"time"
)

func NewApiServer(address string) (*http.Server, error) {

	apiServer := &http.Server{
		Addr: address,
		// ErrorLog:     NoLogger,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}
	fmt.Println("address", address)
	return apiServer, nil
}
