package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type serviceConfig struct {
	RedisHost string
	RedisPort string
	AppPort   int
	Website   []byte
	RedisConn *redis.Pool
}

func (sc *serviceConfig) homeHandler(w http.ResponseWriter, r *http.Request) {
	count := sc.incrementCounter()
	countString := []byte("\nVisit Counter: " + strconv.Itoa(count) + "\n")
	outputText := append(sc.Website[:], countString[:]...)

	fmt.Fprintf(w, string(outputText))

	content, err := ioutil.ReadFile("art.txt")
	if err != nil {
		fmt.Println("Cound not read art.txt")
		return
	}
	fmt.Fprintf(w, string(content))
}

func (sc *serviceConfig) incrementCounter() int {
	conn := sc.RedisConn.Get()
	count, _ := redis.Int(conn.Do("INCR", "webapp-demo"))

	return count
}

var port = flag.Int("port", 80, "Listening port for web server")

func main() {
	logInit()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	var sc serviceConfig
	sc.RedisHost = os.Getenv("RedisEndpoint")
	sc.RedisPort = os.Getenv("RedisPort")
	sc.AppPort = *port
	sc.dbInit()
	defer sc.RedisConn.Close()

	if sc.RedisHost == "" {
		Error.Println("No Redis host provided...exiting...")
		os.Exit(1)
	}

	if sc.RedisPort == "" {
		Error.Println("No Redis port provided...exiting...")
		os.Exit(1)
	}

	art, err := ioutil.ReadFile("art.txt")
	if err != nil {
		Error.Println("Could not open art.txt: ", err)
	}

	sc.Website = art
	fmt.Println("\n", string(art))

	http.HandleFunc("/", sc.homeHandler)

	Info.Printf("starting server on port %v...", *port)

	Error.Println(http.ListenAndServe(fmt.Sprintf(":%d", sc.AppPort), nil))
}
