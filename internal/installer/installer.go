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
		fmt.Println(Sha256File(libInfo.Path))
		(&pkgMgr).runInstall(libInfo)
	}
}