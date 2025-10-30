package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	appVersion = "trunc"
)

const (
	defaultPort     = 8080
	portEnvVariable = "PORT"
)

func main() {
	port := selectPort()
	http.HandleFunc("/", showIP)
	http.HandleFunc("/version", showVersion)
	log.Printf("starting server on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func showVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("app version: %s", appVersion))
}

func showIP(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, "https://ifconfig.me", nil)
	if err != nil {
		log.Printf("can not create new HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// we set "curl" as the user agent, as the response will just be the
	// plain IP then. Otherwise it will be a full HTML page.
	req.Header.Set("User-Agent", "curl")
	resp, err := http.DefaultClient.Do(req)
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
