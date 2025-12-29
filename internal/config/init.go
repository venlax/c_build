package config

import (
	"fmt"
	"strings"
)

func Init(configPath string) {
	PkgMgrName = "apt"
	// Image = "ubuntu:22.04"
	// Image = "gcc:13"
	// Libs = append(Libs, LibInfo{
	// 	// Name : "tmux",
	// 	// Version: "3.2a-4ubuntu0.2",
	// 	Name : "build-essential",
	// 	Version: "12.9ubuntu3",
	// })

	Parse(configPath)

	Image = strings.ReplaceAll(strings.ToLower(Cfg.MetaData.Distribution), " ", ":") // Not stable
	fmt.Printf("Image: <%s>\n", Image)

	fmt.Printf("PkgMgr: <%s>\n", PkgMgrName)

	HostBuildRootDir = Cfg.MetaData.BuildPath

	fmt.Printf("Build root dir: <%s>\n", HostBuildRootDir)

	// fmt.Println("Dependencies:")
	for _, dep := range Cfg.Dependencies {
		lib := LibInfo{
			Name: dep.Name,
			Path: dep.Path,
			Version: dep.Version,
			Sha256: dep.Hash,
		}
		Libs = append(Libs, lib)
		// fmt.Printf("%+v\n", lib)
	}
}