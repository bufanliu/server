package main

import (
	"log"
	"os"

	"github.com/yokaiio/yokai_server/client"
	"github.com/yokaiio/yokai_server/entries"
)

func init() {
	// set working directory as yokai_server
	os.Chdir("../../")
}

func main() {
	// entries init
	entries.InitEntries()

	c, err := client.NewClient()
	if err != nil {
		log.Fatal("client new error:", err)
		os.Exit(1)
	}

	if err = c.Run(os.Args); err != nil {
		log.Fatal("client run error:", err)
		os.Exit(1)
	}

	c.Stop()

}
