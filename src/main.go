package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	. "scf/src/controller"
)

const (
	host_ = "0.0.0.0"
	port_ = 9000
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if _, ok := HandlerMap[url]; !ok {
			res := &Res{}
			res.Code = "404"
			res.Error = errors.New(fmt.Sprintf("url %s not found", url))
			err := json.NewEncoder(rw).Encode(res)
			if err != nil {
				log.Printf("url %s not found", url)
			}
			return
		}
		controllerFunc := HandlerMap[url]
		controller := controllerFunc()
		request := controller.Request()
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			res := &Res{}
			res.Code = "504"
			res.Error = errors.New(fmt.Sprintf("request json error : %s", err.Error()))
			err = json.NewEncoder(rw).Encode(res)
			if err != nil {
				log.Printf("request json error : %s", err.Error())
			}
			return
		}
		response, err := controller.Process(context.Background(), request)
		if err != nil {
			res := &Res{}
			res.Code = "505"
			res.Error = errors.New(fmt.Sprintf("url %s process error : %s", url, err.Error()))
			err = json.NewEncoder(rw).Encode(res)
			if err != nil {
				log.Printf("url %s process error : %s", url, err.Error())
			}
			return
		}
		res := &Res{}
		res.Code = "200"
		res.Response = response
		err = json.NewEncoder(rw).Encode(res)
		if err != nil {
			res.Code = "506"
			res.Error = err
			err = json.NewEncoder(rw).Encode(res)
			if err != nil {
				log.Printf("url %s response error : %s", url, err.Error())
			}
			return
		}
	})
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host_, port_), nil)
	if err != nil {
		log.Fatal("setup server fatal:", err)
	}
}
