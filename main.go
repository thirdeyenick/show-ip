package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultPort     = 8080
	portEnvVariable = "PORT"
)

func main() {
	port := selectPort()
	http.HandleFunc("/", showIP)
	log.Printf("starting server on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func showIP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf("error on gathering public IP: %v", err))
		return
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("public IP: %s", string(body)))
}

func selectPort() string {
	defPort := fmt.Sprintf(":%d", defaultPort)
	if port := os.Getenv(portEnvVariable); port != "" {
		portNumber, err := strconv.Atoi(port)
		if err != nil {
			log.Printf("can not parse %s as a number: %v\n", port, err)
			return defPort

		}
		return fmt.Sprintf(":%d", portNumber)
	}
	return defPort
}
