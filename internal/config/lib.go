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
	Origin string 
}

var Libs []LibInfo

var HasCustom bool = false



