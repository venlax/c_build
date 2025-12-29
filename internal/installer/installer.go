package installer

import (
	"fmt"

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

	for _, libInfo := range config.Libs {
		if !Check(libInfo) {
			panic(fmt.Errorf("dependency <%s> version:<%s> path:<%s> not match the required.", libInfo.Name, libInfo.Version, libInfo.Path))
		}
	}
}