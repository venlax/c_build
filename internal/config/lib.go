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
	FullName string // the lib's fullname in the env, e.g. openssl -> libopenssl.xx.so. if not provided, then = lib{$Name}.a/.so
	Type LibType
}

var Libs []LibInfo



