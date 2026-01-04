package config

import (
	"fmt"
	"log/slog"
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
	slog.Info(fmt.Sprintf("Image: <%s>\n", Image))

	slog.Info(fmt.Sprintf("PkgMgr: <%s>\n", PkgMgrName))

	HostBuildRootDir = Cfg.MetaData.BuildPath

	slog.Info(fmt.Sprintf("Build root dir: <%s>\n", HostBuildRootDir))

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
	slog.Info(fmt.Sprintf("set SOURCE_DATE_EPOCH=%d\n", t.Unix()))
	Env = append(Env, fmt.Sprintf("SOURCE_DATE_EPOCH=%d", t.Unix()))

	CFLAGS := fmt.Sprintf("CFLAGS=-ffile-prefix-map=%s=.",WorkingDir)
	CXXFLAGS := fmt.Sprintf("CXXFLAGS=-ffile-prefix-map=%s=.", WorkingDir)

	slog.Info(fmt.Sprintf("set %s\n", CFLAGS))
	slog.Info(fmt.Sprintf("set %s\n", CXXFLAGS))

	Env = append(Env, CFLAGS, CXXFLAGS)

	locales := strings.Split(Cfg.MetaData.Locale, ";")

	Env = append(Env, locales...)

	slog.Info(fmt.Sprintf("set locale: %s\n", Cfg.MetaData.Locale))

	slog.Info(fmt.Sprintf("Container Env: %v\n", Env))

}