package config


var PkgMgrName = ""

var HostBuildRootDir = ""

var HostReprobuildDir = ""

type DistributionInfo struct {
	Name string
	PkgMgrName string
}

var distros = []DistributionInfo{
	{"ubuntu", "apt"},
	{"debian", "apt"},
	{"alpine", "apk"},
	{"fedora", "dnf"},
	{"centos", "yum"},
	{"rocky",  "dnf"},
	{"arch",   "pacman"},
}