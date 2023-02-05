package main

import (
	"fmt"
	"os"

	"github.com/soducool/auditlog/config"
	"github.com/soducool/auditlog/logger"
)

func main() {
	// init config
	if err := config.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "init config error: %v\n", err)
		os.Exit(1)
	}

	// init log
	if err := logger.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "init logger error: %v\n", err)
		os.Exit(1)
	}

	// init server

	// graceful shutdown

}
