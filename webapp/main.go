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
	//outputEncoded := "ICAgICAgICAgICAgLy8gICAgICAgICAgIC8vCiAgICAgICAgICAgLy8vICAgICAgICAgIC8vLwogICAgICAgICAgLy8vLyAgICAgICAgIC8vLy8KICAgICAgICAgIHwvLy8vICAgICAgIC8vLy8vCiAgICAgICAgICB8KSkvLzsgICAgIC8pKSkvLzsKICAgICAgICAgLykpKSkpLzsgICAvKSkpKSkvOwogICAgIC4tLS1gLCkpKSkvOyAgLykpKSkpKSkvOwogX18tLVwvNi0gIFxgKSkvOyB8KSkpKSkpKS87CigtLS0tLyAgICBcXFxgYDsgIHwpKSkpKSkvOwogICB+Ly1cICBcXFxcXGBgICAgXCkpKSkpKS87CiAgICAgICBcXFxcXFxcXGAgICAgfCkpKSkpLzsKICAgICAgIHxcXFxcXFxcXF9fXy8pKSkpKSkvO19fLS0tLS0tLS4KICAgICAgIC8vLy8vL3wgICUlXy8pKSkpKSkvOyAgICAgICAgICAgXF9fXywKICAgICAgfHx8fHx8fFwgICBcJSUlJVZMSzs6ICAgICAgICAgICAgICBcXy4gXAogICAgICB8XFxcXFxcXFxcICAgICAgICAgICAgICAgICAgICAgICAgfCAgfCB8CiAgICAgICBcXFxcXFxcICAgICAgICAgICAgICAgICAgICAgICAgICB8ICB8IHwKICAgICAgICB8XFxcXCAgICAgICAgICAgICAgIF9ffCAgICAgICAgLyAgIC8gLwogICAgICAgIHwgXFxfX1wgICAgIFxfX18tLS0tICB8ICAgICAgIHwgICAvIC8KICAgICAgICB8ICAgIC8gfCAgICAgPiAgICAgXCAgIFwgICAgICBcICAvIC8KICAgICAgICB8ICAgLyAgfCAgICAvICAgICAgIFwgICBcICAgICAgPi8gLyAgLCwKICAgICAgICB8ICAgfCAgfCAgIHwgICAgICAgICB8ICAgfCAgICAvLyAvICAvLywKICAgICAgICB8ICAgfCAgfCAgIHwgICAgICAgICB8ICAgfCAgIC98IHwgICB8XFwsCiAgICAgXy0tJyAgIF8tLScgICB8ICAgICBfLS0tXy0tLScgIHwgIFwgXF9fL1x8LwogICAgKC0oLT09PSgtKC0oPT09LyAgICAoLSgtPSgtKC0oPT0vICAgXF9fX18vCiAgICAgICAgICAgICAgICAgICAtVmFsa3lyaWUtCgoK"
	//outputText, err := base64.StdEncoding.DecodeString(string(sc.Website))
	count := sc.incrementCounter()
	countString := []byte("\nVisit Counter: " + strconv.Itoa(count) + "\n")
	outputText := append(sc.Website[:], countString[:]...)

	fmt.Fprintf(w, string(outputText))
}

func (sc *serviceConfig) incrementCounter() int {
	conn := sc.RedisConn.Get()
	count, _ := redis.Int(conn.Do("INCR", "webapp-demo"))

	return count
}

var port = flag.Int("port", 80, "Port to run webserver on")

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

	//copy(sc.Website, art)
	sc.Website = art
	fmt.Println("Art: ", string(art))

	http.HandleFunc("/", sc.homeHandler)

	Info.Println("starting server...")

	Error.Println(http.ListenAndServe(fmt.Sprintf(":%d", sc.AppPort), nil))
}
