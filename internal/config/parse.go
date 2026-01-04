package config

import (
	// "fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MetaData MetaData `yaml:"metadata"`
	Dependencies []Dependency `yaml:"dependencies"`
	Artifacts []Artifact `yaml:"artifacts"`
}

type MetaData struct {
	Architecture string `yaml:"architecture"`
	Distribution string `yaml:"distribution"`
	BuildPath string `yaml:"build_path"`
	BuildTimeStamp string `yaml:"build_timestamp"`
	HostName string `yaml:"hostname"`
	Locale string	`yaml:"locale"`
	Umask string	`yaml:"umask"`
	RandomSeed string `yaml:"random_seed"` // TODO
}

type Dependency struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Version string `yaml:"version"`
	Hash string `yaml:"hash"`	
}


type Artifact struct {
	Path string `yaml:"path"`
	Hash string	`yaml:"hash"`
	Type string	`yaml:"type"`
}

var Cfg Config

func Parse(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		panic(err)
	}

	// fmt.Println(Cfg)
}