package main
//"github.com/gorilla/mux"
// "encoding/json"
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v\n", config)
	timeout := 3 * time.Second

	client, err := docker.NewDefaultClient(timeout)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
	}
	//Pulling image before start
	log.Printf("Start pull: " + config.Image)
	err = client.PullImage(config.Image, &docker.AuthConfig{})
	log.Printf("End pull: " + config.Image)
  if err != nil {
      fmt.Println("ERROR pulling:", err)
  }

	log.Printf("Launching image: " + config.Image)
	cid, err := client.CreateContainer(&config, config.Hostname)
	if err != nil {
	 	//json.NewEncoder(w).Encode(Articles{Error: err})
		w.WriteHeader(http.StatusInternalServerError)
	}

	// start the container
	err = client.StartContainer(cid, &docker.HostConfig{})
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
	}
  message := &Response{Message: "Success for deploy "+config.Image}


	json.NewEncoder(w).Encode(message)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	}

// func deploytest(w http.ResponseWriter, r *http.Request) {
// 	client, _ := docker.NewClientFromEnv()
// }
