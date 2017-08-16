package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		configPath string
		logLevel   int
	)
	flag.StringVar(&configPath, "c", "", "input config file path")
	flag.IntVar(&logLevel, "d", 0, "input log level")
	flag.Parse()
	fmt.Println("path:", configPath)
	fmt.Println("logLevel:", logLevel)
}
