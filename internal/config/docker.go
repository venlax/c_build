package config

var Image string = ""

var ContainerName string = "unspecified" // container already exists just reuse it not remove it

// var ContainerName string = "unspecified-gcc" // container already exists just reuse it not remove it

var WorkingDir = "/ws"

var Env []string = []string {
	"http_proxy=http://127.0.0.1:7890",
	"https_proxy=http://127.0.0.1:7890",
}