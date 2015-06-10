package main

import (
	"flag"
	"fmt"
	"github.com/xsar/healthz/cron"
	"github.com/xsar/healthz/g"
	"github.com/xsar/healthz/http"
	"log"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	if err := g.ParseConfig(*cfg); err != nil {
		log.Fatalln(err)
	}

	go http.Start()
	go cron.HealthCheck()

	select {}
}
