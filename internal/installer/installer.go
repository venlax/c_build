package installer

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/venlax/c_build/internal/config"
)


func Init() {
	pkgMgr := GetPkgMgr(config.PkgMgrName)
	(&pkgMgr).runUpdate() 
}


func Install() {
	pkgMgr := GetPkgMgr(config.PkgMgrName)

	// for _, libInfo := range config.Libs {
	// 	(&pkgMgr).runInstall(libInfo)
	// 	// tmp := GetPkgMgr("dpkg")
	// 	// (&tmp).runGetLibVersion(libInfo)
	// }

	(pkgMgr).runInstallAll()

	for _, libInfo := range config.Libs {
		if !Check(libInfo) {
			panic(fmt.Errorf("dependency <%s> version:<%s> path:<%s> not match the required.", libInfo.Name, libInfo.Version, libInfo.Path))
		}
	}
}

func InstallStrs() []string {
	var res []string
	pkgMgr := GetPkgMgr(config.PkgMgrName)
	res = append(res, commandStr((&pkgMgr).updateCommand, []string{}))
	tmp := make([]string, len(config.Libs))
	tpl, err := template.New("lib_full_name").Parse(pkgMgr.versionTmpl)
	if err != nil {
		panic(err)
	}
	for i, libInfo := range config.Libs {
		if libInfo.Origin == "custom" {
			continue
		}
		var arg string
		if libInfo.Version == "" {
			arg = libInfo.Name
		} else {
			var buf bytes.Buffer
			err := tpl.Execute(&buf, libInfo)
			if err != nil {
				panic(err)
			}
			arg = buf.String()
		}
		tmp[i] = arg
	} 
	res = append(res, commandStr((&pkgMgr).installCommand, tmp))
	return res
}