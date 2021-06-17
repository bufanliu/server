package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/east-eden/server/services/mail"
	"github.com/east-eden/server/utils"
	log "github.com/rs/zerolog/log"
)

var (
	BinaryVersion string
	GoVersion     string
	GitLastLog    string
)

func version() {
	fmt.Println("BinaryVersion:", BinaryVersion)
	fmt.Println("GoVersion:", GoVersion)
	fmt.Println("GitLastLog:", GitLastLog)
	os.Exit(0)
}

func help() {
	fmt.Println("The commands are:")
	fmt.Println("version       see all versions")
	os.Exit(0)
}

func main() {
	utils.LDFlagsCheck(os.Args, version, help)

	m := mail.New()
	if err := m.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("mail run failed")
	}

	m.Stop()
}
