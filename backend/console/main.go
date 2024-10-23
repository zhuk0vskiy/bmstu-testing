package main

import (
	"fmt"
	"github.com/Mx1q/ppo_console/app"
	"github.com/Mx1q/ppo_console/config"
	"github.com/Mx1q/ppo_services/logger"
	"log"
	"os"
)

func main() {
	//ctx := context.Background()
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	logFile, err := os.OpenFile(cfg.Logger.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(logFile)
	logger := logger.NewLogger(cfg.Logger.Level, logFile)

	//db, err := app.NewConn(ctx, cfg)
	//if err != nil {
	//	logger.Fatalf("Unable to connect to database: %v", err)
	//}
	fmt.Println("running")
	app.Run(cfg.CurrentDb.Db, cfg, logger)
}
