package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"flag"

	"gopkg.in/yaml.v3"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"net/http"
)

type Config struct {
	Port *int `json:"port"`
	Data *string `json:"data"`
}

func readConfig(filename string) (*Config,error) {
	config := &Config{
		Port: new(int),
		Data: new(string),
	}

	// Default values
	*config.Port = 3000
	*config.Data = "public"

	log.Info().Str("file",filename).Msg("Reading configuration")

	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buffer, config) 
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	filename := flag.String("config","config.yaml","configuration file localtion")
	flag.Parse()

	config , err := readConfig(*filename)
	if err != nil {
		log.Error().Err(err).Msg("Can't read configuration")
		os.Exit(1)
	}

	log.Info().Interface("config", *config).Msg("Config read")

	http.Handle("/", http.FileServer(http.Dir(*config.Data)))

	log.Info().Msg("Serving...")

	bindAddr := fmt.Sprintf(":%d",*config.Port)

	log.Info().Str("addr", bindAddr).Msg("Serving...")

	err = http.ListenAndServe(bindAddr, nil)
	if err != nil {
		log.Error().Err(err).Msg("can't start server")
		os.Exit(1)
	}

	os.Exit(0)
}