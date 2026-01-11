package config

import (
	// "fmt"
	"os"
	"regexp"
	"strings"

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

var versionRe = regexp.MustCompile(`\b\d+(?:\.\d+)*\b`)

func Parse(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		panic(err)
	}
	
	lower := strings.ToLower(Cfg.MetaData.Distribution)

	for _, d := range distros {
		if strings.HasPrefix(lower, d) {
			ver := versionRe.FindString(lower)
			if ver == "" {
				panic("the distribution version tag is empty.")
			}
			if d == "ubuntu" {
				parts := strings.Split(ver, ".")
				if len(parts) >= 2 {
					ver = parts[0] + "." + parts[1]
				}
			}
			Cfg.MetaData.Distribution = d + ":" + ver
			break
		}
	}


	// fmt.Println(Cfg)
}