package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kaimu/speedtest/providers/netflix"
	"github.com/rs/cors"
	"github.com/showwin/speedtest-go/speedtest"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Load .env file to use the environment variable
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := mux.NewRouter()
	p := os.Getenv("PORT")

	r.HandleFunc("/",SpeedTester).Methods("GET")

	c := cors.New(cors.Options{

		// AllowCredentials: true,
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedOrigins:     []string{"*"},
		AllowedHeaders:     []string{"Content-Type", "content-type", "Origin", "Accept", "Access-Control-Allow-Origin"},
		OptionsPassthrough: false,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})
	handler := c.Handler(r)
	fmt.Println("Server is running on port: ", strings.Split(":"+p, ":")[1])
	log.Fatal(http.ListenAndServe(":"+p, handler))
}

func Netflix() (up float64, err error) {
	return netflix.Fetch()
}

func SpeedTester(w http.ResponseWriter, r *http.Request) {
	user, err := speedtest.FetchUserInfo()

	serverList, err := speedtest.FetchServers(user)
	targets, err := serverList.FindServer([]int{})

	if err != nil {
		fmt.Println(err)
	}

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(true)
		s.UploadTest(true)

		fmt.Printf("Speed (SpeedTest): Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
	}

	res, err := Netflix()
	fmt.Print("Speed (Fast): ")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(res)
}