package config

type LibType int
const (
	STATIC LibType = 1 << iota
	SHARED
) 

type LibInfo struct {
	Name string
	Version string
	Sha256 string
	Path string
	Type LibType
}

var Libs []LibInfo



