package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var validBuckets = []string{
	"avatars",
	"subtitles",
	"posters",
}

type ConfMap struct {
	App struct {
		Version string `yaml:"version"`
		Debug   bool   `yaml:"debug"`
		Env     string `yaml:"env"`
	} `yaml:"app"`
	Grpc struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"grpc"`
	Http struct {
		AccessControlAllowOrigin string `yaml:"access_control_allow_origin"`
	} `yaml:"http"`
	Secrets struct {
		ObjectStorage struct {
			Endpoint           string `yaml:"endpoint"`
			Region             string `yaml:"region"`
			UseHttps           bool   `yaml:"use_https"`
			InsecureSkipVerify bool   `yaml:"insecure_skip_verify"`
			AccessKey          string `yaml:"access_key"`
			SecretKey          string `yaml:"secret_key"`
		}
		SentryDsn      string `yaml:"sentry_dsn"`
		HcaptchaSecret string `yaml:"hcaptcha_secret"`
	} `yaml:"secrets"`
}

var Map = new(ConfMap)

func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open config file: %v", err)
	}
	if err := yaml.NewDecoder(file).Decode(&Map); err != nil {
		return fmt.Errorf("could not decode config file: %v", err)
	}
	log.Printf("ConfigMap Loaded: [version: %s]", Map.App.Version)
	return nil
}

func IsValidBucketName(bucketname string) bool {
	for _, bk := range validBuckets {
		if bucketname == bk {
			return true
		}
	}
	return false
}
