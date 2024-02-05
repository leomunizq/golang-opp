package main // package main is the entry point of the program

import (
	"github.com/leomunizq/golang-opp/config"
	"github.com/leomunizq/golang-opp/router"
)

var (
	logger *config.Logger
)

func main() {
 logger = config.GetLogger("main")
	// Call the Init function from the config package
	err:= config.Init() 
	if err != nil {

logger.Errf("Error initializing the application: %v", err)
return
	}

	// Call the Initialize function from the router package
  router.Initialize()
}