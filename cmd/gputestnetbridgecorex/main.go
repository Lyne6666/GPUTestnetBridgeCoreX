// cmd/gputestnetbridgecorex/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"gputestnetbridgecorex/internal/gputestnetbridgecorex"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := gputestnetbridgecorex.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
