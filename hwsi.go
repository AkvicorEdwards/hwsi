package main

import (
	"fmt"
	"github.com/AkvicorEdwards/argsparser"
	"hwsi/config"
	"hwsi/handler"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	setDir()
	parseArgs()

	handler.Init()
	server := http.Server{
		Addr:              config.Addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Minute,
	}
	fmt.Printf("WORK   DIR: [%s]\n", config.WorkDir)
	fmt.Printf("UPLOAD DIR: [%s]\n", config.UploadDir)
	fmt.Printf("PASSWORD:   [%s]\n", config.Password)
	ips := GetIntranetIp()
	for k, v := range ips {
		fmt.Printf("[%d] http://%s%s\n", k, v, config.Addr)
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func setDir() {
	config.WorkDir, _ = os.Getwd()
	config.WorkDir += "/"
	config.UploadDir, _ = os.Getwd()
	config.UploadDir += "/"
}

func parseArgs() {
	argsparser.Version = "2.0"
	argsparser.Help = `
Usage: hwsi [option...] [arguments...]

Example: 

	hwsi port 7020

The commands are:

	title  [title]      Set server title
	port   [port]       Set http listening port
	work   [path]       Set work dir
	upload [path]       Set upload dir
	password [password] Set upload password

Default:

	title   Akvicor's file System
	port    8021
	work    ./
	upload  ./
`
	argsparser.AddBasicArg()
	argsparser.Add("title", 1, func(str []string) {
		config.Title = str[1]
	})
	argsparser.Add("port", 1, func(str []string) {
		config.Addr = fmt.Sprintf(":%s", str[1])
	})
	argsparser.Add("work", 1, func(str []string) {
		config.Addr = str[1]
	})
	argsparser.Add("upload", 1, func(str []string) {
		config.Addr = str[1]
	})
	argsparser.Add("password", 1, func(str []string) {
		config.Password = str[1]
	})

	argsparser.Parse()
}

func GetIntranetIp() (ip []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = append(ip, ipnet.IP.String())
			}
		}
	}
	return
}