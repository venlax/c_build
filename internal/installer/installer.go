package installer

import (
	"github.com/venlax/c_build/internal/config"
)


func Init() {
	pkgMgr := GetPkgMgr(config.PkgMgrName)
	(&pkgMgr).runUpdate() 
}


func Install() {
	pkgMgr := GetPkgMgr(config.PkgMgrName)

	for _, libInfo := range config.Libs {
		(&pkgMgr).runInstall(libInfo)
		// tmp := GetPkgMgr("dpkg")
		// (&tmp).runGetLibVersion(libInfo)
	}
}