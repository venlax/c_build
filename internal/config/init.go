package config

import (
	"fmt"
	"strings"
	"time"
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

    t, _ := time.Parse(time.RFC3339, Cfg.MetaData.BuildTimeStamp)
	fmt.Printf("set SOURCE_DATE_EPOCH=%d\n", t.Unix())
	Env = append(Env, fmt.Sprintf("SOURCE_DATE_EPOCH=%d", t.Unix()))

	CFLAGS := fmt.Sprintf("CFLAGS=-ffile-prefix-map=%s=.",WorkingDir)
	CXXFLAGS := fmt.Sprintf("CXXFLAGS=-ffile-prefix-map=%s=.", WorkingDir)

	fmt.Printf("set %s\n", CFLAGS)
	fmt.Printf("set %s\n", CXXFLAGS)

	Env = append(Env, CFLAGS, CXXFLAGS)

	fmt.Printf("Container Env: %v\n", Env)

}