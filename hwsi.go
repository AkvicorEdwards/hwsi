package main

import (
	"fmt"
	"hwsi/config"
	"hwsi/handler"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const Version string = "1.0.0"

const Help string = `
Hwsi is a file indexer for HTTP web servers with focus on your files.

Usage:

	config config_file      use config file
	port port               port
	title title             Title
	password password       password
	work dir                work directory
	upload dir              upload directory
	version print           Hwsi version
`

const WrongArgs string = `
Unknown command
Run 'hwsi help' for usage.
`

func main() {

	wd, _ := os.Getwd()
	fmt.Println(os.Args[0], "\n", wd)

	parseArgs()

	//if err := config.Data.Get(); err != nil {
	//	log.Error(err)
	//}
	handler.Init()
	server := http.Server{
		Addr:              config.Data.Server.Addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Second,
	}
	fmt.Println(config.Data.String())
	log.Println("ListenAndServe:", config.Data.Server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func parseArgs() {
	args := make(map[string]string)
	args["config"] = ""
	args["port"] = "8021"
	args["title"] = "Hwsi"
	args["password"] = ""
	args["work"], _ = os.Getwd()
	args["upload"], _ = os.Getwd()
	args["theme"] = "ori"

	lastArg, arg := "", ""
	idx := 0

	for idx, arg = range os.Args {
		if idx == 0 {
			continue
		}
		if arg == "help" {
			fmt.Print(Help)
			os.Exit(0)
		}
		if arg == "version" {
			fmt.Println(Version)
			os.Exit(0)
		}
		if (idx&1) == 1 {
			lastArg = arg
		} else {
			if _, ok := args[lastArg]; ok {
				args[lastArg] = arg
			} else {
				fmt.Print(WrongArgs)
				os.Exit(0)
			}
		}
	}

	if (idx&1) == 1 {
		fmt.Print(WrongArgs)
		os.Exit(0)
	}

	if len(args["config"]) == 0 {
		if err := config.Data.GetByMap(args); err != nil {
			fmt.Println("read by map error")
			os.Exit(0)
		}
	}else {
		if err := config.Data.GetByFile(args["config"]); err != nil {
			fmt.Println("read by file error")
			os.Exit(0)
		}
	}

}
