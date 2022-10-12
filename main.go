package main

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"github.com/kaimu/speedtest/providers/netflix"
)

func main() {
	user, err := speedtest.FetchUserInfo()

	serverList, err := speedtest.FetchServers(user)
	targets, err := serverList.FindServer([]int{})

	if err != nil{
		fmt.Println(err)
	}

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(true)
		s.UploadTest(true)

		fmt.Printf("Speed (SpeedTest): Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
	}

	res,err := Netflix()
	fmt.Print("Speed (Fast): ")
	if(err != nil){
		fmt.Print(err)
	}
	fmt.Print(res)
}

func Netflix() (up float64, err error) {
	return netflix.Fetch()
}