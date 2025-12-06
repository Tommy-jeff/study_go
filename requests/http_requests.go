package requests

import (
	"log"
	"net/http"
)


func GoogleRequest() (http.Response, error) {
	res, err := http.Get("http://google.com.br")
	if err != nil {
		log.Fatal("Error during execute http.Get: ", err.Error())
		return *res.Request.Response, err
	}

	return *res.Request.Response, nil
}