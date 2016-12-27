package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
	"github.com/yhat/go-docker"


)
type Response struct {
	Message  string `json:"Message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func deploy(w http.ResponseWriter, r *http.Request) {
	var config docker.ContainerConfig
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &config); err != nil {
		w.Header().Set("Content-Type", "application/json;")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//Debug
	//fmt.Printf("%+v\n", config)
	timeout := 3 * time.Second

	client, err := docker.NewDefaultClient(timeout)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf("Removing : " + config.Hostname)
	err = client.RemoveContainer(config.Hostname,true, true)
		if err != nil {
	      fmt.Println("ERROR remove:", err)
	  }

	//Pulling image before start
	log.Printf("Start pull: " + config.Image)
	err = client.PullImage(config.Image, &docker.AuthConfig{})
  if err != nil {
      fmt.Println("ERROR pulling:", err)
			w.WriteHeader(http.StatusInternalServerError)
			message := &Response{Message: "Failed pulling for "+config.Image}
			json.NewEncoder(w).Encode(message)
			return
  }
	log.Printf("End pull: " + config.Image)

	log.Printf("Launching image: " + config.Image)
	cid, err := client.CreateContainer(&config, config.Hostname)
	if err != nil {
		fmt.Println("ERROR CreateContainer:", err)
		w.WriteHeader(http.StatusInternalServerError)
		message := &Response{Message: "Failed CreateContainer for "+config.Image}
		json.NewEncoder(w).Encode(message)
		return
	}

	// start the container
	err = client.StartContainer(cid, &config.HostConfig)
	if err != nil {
		fmt.Println("ERROR StartContainer:", err)
		w.WriteHeader(http.StatusInternalServerError)
		message := &Response{Message: "Failed StartContainer for "+config.Image}
		json.NewEncoder(w).Encode(message)
		return
	}
	message := &Response{Message: "Success for deploy "+config.Image}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)

}
