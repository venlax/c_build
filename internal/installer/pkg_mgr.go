package installer

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/docker"
)

type commandArgsPassWay int

const (
	Tail commandArgsPassWay = iota
	Xargs 
	NoArgs
)

type command struct {
	strs []string
	passWay commandArgsPassWay
}

type pkgMgr struct{
	name string
	updateCommand command
	installCommand command
	getLibVersionCommand command
	
}

var pkgMgrs map[string]pkgMgr = map[string]pkgMgr {
	"apt" : {
		"apt", 
		command {
			[]string{"apt", "update"},
			NoArgs,
		}, 
		command {
			[]string{"apt", "install", "-y", "--allow-downgrades"},
			Tail,
		},  
		command {
			[]string{},
			NoArgs,
		},
	},
	"dpkg" : {
		"dpkg", 
		command {
			[]string{}, 
			NoArgs,
		}, 
		command {
			[]string{}, 
			NoArgs,
		}, 
		command {
			[]string{"dpkg", "-s", "|", "grep", "^Version"},
			Xargs,
		},
	},
}

func GetPkgMgr(name string) pkgMgr {
	return pkgMgrs[name]
}

func (p *pkgMgr) runUpdate() {
	runCommand(p.updateCommand, []string{}, os.Stdout)
}

func (p *pkgMgr) runInstall(libInfo config.LibInfo) {
	var arg string
	if libInfo.Version == "" {
		arg = libInfo.Name
	} else {
		arg = libInfo.Name + "=" + libInfo.Version
	}
	runCommand(p.installCommand, []string{arg}, os.Stdout)
	// argvs := make([]string, 0, len(p.installCommand.strs) + 1)
	// argvs = append(argvs, p.installCommand.strs...)
	// argvs = append(argvs, libInfo.Name + "=" + libInfo.Version)
	// err := docker.Run(argvs)
	// if (err != nil) {
	// 	panic(err)
	// }

	// go to checker please
} 

func (p *pkgMgr) runInstallAll() {
	tmp := make([]string, len(config.Libs))
	for i, libInfo := range config.Libs {
		var arg string
		if libInfo.Version == "" {
			arg = libInfo.Name
		} else {
			arg = libInfo.Name + "=" + libInfo.Version
		}
		tmp[i] = arg
	} 
	runCommand(p.installCommand, tmp, os.Stdout)
} 

func (p *pkgMgr) runGetLibVersion(libInfo config.LibInfo) string {
	var sb strings.Builder 
	runCommand(p.getLibVersionCommand, []string{libInfo.Name}, &sb)
	versionRawStr := sb.String()
	return strings.TrimSpace(strings.TrimPrefix(versionRawStr, "Version:"))
}

func runCommand(c command, args []string, writer io.Writer) {
	cmd := make([]string, 0)
	switch c.passWay {
	case NoArgs:
		cmd = append(cmd, c.strs...)
	case Tail:
		cmd = append(cmd, c.strs...)
		cmd = append(cmd, args...)
	case Xargs:
		cmd = append(cmd, "sh", "-c")
		realCmdStr := fmt.Sprintf(`echo %s | xargs %s`, strings.Join(args, ", "), strings.Join(c.strs, " "))
		cmd = append(cmd, realCmdStr)		
		// cmd = append(cmd, "xargs")
		// cmd = append(cmd, args...)
		// cmd = append(cmd)
	default:
		panic("Something wrong.")
	}

	err := docker.Run(cmd, writer)
	if (err != nil) {
		panic(err)
	}
}

func commandStr(c command, args []string) string {
	cmd := make([]string, 0)
	switch c.passWay {
	case NoArgs:
		cmd = append(cmd, c.strs...)
	case Tail:
		cmd = append(cmd, c.strs...)
		cmd = append(cmd, args...)
	case Xargs:
		cmd = append(cmd, "sh", "-c")
		realCmdStr := fmt.Sprintf(`echo %s | xargs %s`, strings.Join(args, ", "), strings.Join(c.strs, " "))
		cmd = append(cmd, realCmdStr)		
		// cmd = append(cmd, "xargs")
		// cmd = append(cmd, args...)
		// cmd = append(cmd)
	default:
		panic("Something wrong.")
	}
	return strings.Join(cmd, " ")
}