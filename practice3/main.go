package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	VERSION = "VERSION"
)

func main() {
	//log output
	out := zerolog.MultiLevelWriter(os.Stdout)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: out, TimeFormat: time.RFC3339})

	//config
	configFile, err := getConfigFile()
	if err != nil {
		log.Error().Err(err).Msg("get config file failed")
		return
	}
	conf, err := loadConfig(configFile)
	if err != nil {
		log.Error().Err(err).Msg("load config failed")
		return
	}
	http.HandleFunc("/svc", WebHandler)
	http.HandleFunc("/healthz", HealthHandler)
	log.Info().Msg("http server is running")
	err = http.ListenAndServe(conf.Server.Port, nil)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	log.Fatal().Err(err).Msg("http server run over")
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	rHeader := r.Header
	for k, v := range rHeader {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	w.Header().Add(VERSION, os.Getenv(VERSION))
	w.WriteHeader(http.StatusOK)
	fmt.Printf("client host:%s, response code:%d", r.Host, http.StatusOK)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

type Config struct {
	yamlConfig
	Log      string
	LogLevel zerolog.Level
}
type yamlConfig struct {
	Server serverConfig
}
type serverConfig struct {
	Port string
}

func getConfigFile() (string, error) {
	defaultConfigFile, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	var configFileName string
	pflag.StringVar(&configFileName,
		"config",
		filepath.Join(defaultConfigFile, "httpserver.yaml"),
		"httpserver config file")
	pflag.Parse()
	return configFileName, nil
}

func loadConfig(fileName string) (*Config, error) {
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn().Err(err).Msg("viper read in config failed, use default")
		return nil, err
	}
	config := &Config{Log: "info"}
	err = viper.Unmarshal(&config.yamlConfig)
	if err != nil {
		return nil, err
	}
	config.LogLevel, err = zerolog.ParseLevel(config.Log)
	if err != nil {
		return nil, err
	}
	return config, nil
}
