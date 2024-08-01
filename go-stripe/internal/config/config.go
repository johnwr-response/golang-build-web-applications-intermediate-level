package config

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"strings"
)

//go:embed config.yml
var defaultConfiguration []byte

type Web struct {
	HostInterface string `json:"host_interface,omitempty"`
	Port          int    `json:"port,omitempty"`
	Dsn           string `json:"dsn,omitempty"`
}
type Api struct {
	HostInterface string `json:"host_interface,omitempty"`
	Port          int    `json:"port,omitempty"`
	Dsn           string `json:"dsn,omitempty"`
}
type Urls struct {
	Frontend string `json:"frontend,omitempty"`
	Api      string `json:"api,omitempty"`
	Invoice  string `json:"invoice,omitempty"`
}
type Smtp struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
type Stripe struct {
	Key    string `json:"key,omitempty"`
	Secret string `json:"secret,omitempty"`
}
type Payment struct {
	Stripe Stripe `json:"stripe,omitempty"`
}

type Config struct {
	Env       string  `json:"env,omitempty"`
	Dsn       string  `json:"dsn,omitempty"`
	SecretKey string  `json:"secret_key,omitempty"`
	Web       Web     `json:"web,omitempty"`
	Api       Api     `json:"api,omitempty"`
	Urls      Urls    `json:"urls,omitempty"`
	Smtp      Smtp    `json:"smtp,omitempty"`
	Payment   Payment `json:"payment,omitempty"`
}

func Read() (*Config, error) {
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Configuration file
	viper.SetConfigType("yml")
	// Read configuration
	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		return nil, err
	}
	// Unmarshal the configuration
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
