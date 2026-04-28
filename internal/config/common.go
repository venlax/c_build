package config

import (
	"strings"
	"time"
)


var PkgMgrName = ""

var HostBuildRootDir = ""

var HostReprobuildDir = ""

type DistributionInfo struct {
	Name string
	PkgMgrName string
	Versions []VersionInfo
}

type VersionInfo struct {
	Version string
	CodeName string
}

var Distros = []DistributionInfo {
	{
		Name:       "ubuntu",
		PkgMgrName: "apt",
		Versions: []VersionInfo {
			{"10.04", "lucid"},
			{"12.04", "precise"},
			{"12.10", "quantal"},
			{"13.04", "raring"},
			{"13.10", "saucy"},
			{"14.04", "trusty"},
			{"14.10", "utopic"},
			{"15.04", "vivid"},
			{"15.10", "wily"},
			{"16.04", "xenial"},
			{"16.10", "yakkety"},
			{"17.04", "zesty"},
			{"17.10", "artful"},
			{"18.04", "bionic"},
			{"18.10", "cosmic"},
			{"19.04", "disco"},
			{"19.10", "eoan"},
			{"20.04", "focal"},
			{"20.10", "groovy"},
			{"21.04", "hirsute"},
			{"21.10", "impish"},
			{"22.04", "jammy"},
			{"22.10", "kinetic"},
			{"23.04", "lunar"},
			{"23.10", "mantic"},
			{"24.04", "noble"},
			{"24.10", "oracular"},
			{"25.04", "plucky"},
			{"25.10", "questing"},
			{"26.04", "resolute"},
		},
	},
	{
		Name:       "debian",
		PkgMgrName: "apt",
		Versions: []VersionInfo {
			{"6.0", "squeeze"},
			{"7", "wheezy"},
			{"8", "jessie"},
			{"9", "stretch"},
			{"10", "buster"},
			{"11", "bullseye"},
			{"12", "bookworm"},
			{"13", "trixie"},
			{"14", "forky"},
		},
	},
	{
		Name:       "alpine",
		PkgMgrName: "apk",
		Versions: []VersionInfo {
			{"3.18", "n/a"},
			{"3.17", "n/a"},
			{"3.16", "n/a"},
		},
	},
	{
		Name:       "fedora",
		PkgMgrName: "dnf",
		Versions: []VersionInfo {
			{"38", "n/a"},
			{"37", "n/a"},
			{"36", "n/a"},
		},
	},
	{
		Name:       "centos",
		PkgMgrName: "yum",
		Versions: []VersionInfo {
			{"8", "n/a"},
			{"7", "n/a"},
		},
	},
	{
		Name:       "rocky",
		PkgMgrName: "dnf",
		Versions: []VersionInfo {
			{"9", "n/a"},
			{"8", "n/a"},
		},
	},
	{
		Name:       "arch",
		PkgMgrName: "pacman",
		Versions: []VersionInfo {},
	},
}

var EnableTimer bool
var BuildTime time.Duration

func GetDistroInfo(distroName string) *DistributionInfo {
	for idx := range Distros {
		if Distros[idx].Name == distroName {
			return &Distros[idx]
		}
	}
	return nil
}

func GetCodeName(distroName, distroVersion string) string {
	distro := GetDistroInfo(distroName)
	switch distroName {
		case "ubuntu":
			for _, ve := range distro.Versions {
				if strings.HasPrefix(distroVersion, ve.Version) {
					return ve.CodeName;
				}
			}
		case "debian":
			for _, ve := range distro.Versions {
				if distroVersion == ve.Version {
					return ve.CodeName
				}
			}
	}
	return "n/a"
}