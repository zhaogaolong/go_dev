package main

import (
	"flag"
	"fmt"
)

func main() {
	var configPath string
	var logLevel int
	var debug bool

	flag.StringVar(&configPath, "c", "", "please input config file path")
	flag.IntVar(&logLevel, "d", 10, "please input log level")
	flag.BoolVar(&debug, "l", false, "if you on debug")
	flag.Parse()

	fmt.Println("configFile path:", configPath)
	fmt.Println("log level:", logLevel)
	fmt.Println("debug status", debug)
}
