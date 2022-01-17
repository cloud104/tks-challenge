package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var Start int
var Stop int

func hostname(w http.ResponseWriter, _ *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		w.WriteHeader(500)
		_, _ = fmt.Fprintf(w, "ERROR getting hostname")
		return
	}
	_, _ = fmt.Fprintf(w, "%s\n", hostname)
}

func randomError() {
	for {
		rangeLower := 0
		rangeUpper := 59
		randomMin := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
		Start = randomMin
		Stop = randomMin + 5
		if Stop > 59 {
			Stop = Stop - 59
		}
		time.Sleep(5 * time.Minute)
	}

}

func health(w http.ResponseWriter, _ *http.Request) {
	_, minutes, _ := time.Now().Clock()
	if (minutes >= Start) && (minutes <= Stop) {
		w.WriteHeader(500)
		log.Printf("bad status code triggered: Starting at minute: %d until minute: %d\n", Start, Stop)
		_, _ = fmt.Fprintf(w, "something went wrong\n")
	} else {
		_, _ = fmt.Fprintf(w, "OK\n")
	}
}

func init() {
	log.SetPrefix("TKS-Challenge: ")
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	go randomError()
	http.HandleFunc("/", hostname)
	http.HandleFunc("/health", health)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
