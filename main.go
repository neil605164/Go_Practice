package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("http request.")
	fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func pinger(port string) error {
	resp, err := http.Get("http://localhost:" + port)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("server return not 200 status code")
	}

	return nil
}

func main() {
	var port string
	var ping bool

	flag.StringVar(&port, "port1", "8080", "server port")
	flag.StringVar(&port, "port2", "8081", "server port")
	flag.BoolVar(&ping, "ping", false, "check_docker_server")
	flag.Parse()

	// check http service
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	// check docker service
	if ping {
		if err := pinger(port); err != nil {
			log.Printf("Ping docker server error: %v \n", err)
		}

		return
	}

	http.HandleFunc("/", handler)
	log.Println("http server run on ", port, " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
