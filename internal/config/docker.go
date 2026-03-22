package config

var Image string = ""

var ImageTag string = "" // if minor version not needed, equals to Image

var ContainerName string = "unspecified" // container already exists just reuse it not remove it

// var ContainerName string = "unspecified-gcc" // container already exists just reuse it not remove it

var WorkingDir = "/ws"

var ReprobuildDir = "/opt/reprobuild"

var GraphOutputPath = ""

var Env []string = []string {
	"http_proxy=http://127.0.0.1:7890",
	"https_proxy=http://127.0.0.1:7890",
	"CC=/usr/bin/x86_64-linux-gnu-gcc",
	"CXX=/usr/bin/x86_64-linux-gnu-g++",
	// "CFLAGS=-ffile-prefix-map=/ws=.",
	// "CXXFLAGS=-ffile-prefix-map=/ws=.",
}

var BuildCmd string = "make"