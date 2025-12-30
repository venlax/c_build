package builder

import (
	// "fmt"
	"fmt"
	"os"

	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/docker"
	"github.com/venlax/c_build/internal/installer"
	// "github.com/venlax/c_build/internal/installer"
)

func Build() {
	// err := docker.Run([]string{"cd", config.WorkingDir}, os.Stdout)
	// if err != nil {
	// 	panic(err)
	// }

	// err := docker.Run([]string{"make", "deps"}, os.Stdout)
	// if err != nil {
	// 	panic(err)
	// }

	err := docker.Run([]string{"make", "clean"}, os.Stdout)
	if err != nil {
		panic(err)
	}

	MakeCommand := fmt.Sprintf("umask %s && make", config.Cfg.MetaData.Umask)

	err = docker.Run([]string{"sh", "-c", MakeCommand}, os.Stdout)
	if err != nil {
		panic(err)
	}

	// fmt.Println(installer.Sha256File("/ws/lua"))

	// err = docker.Run([]string{"./hello"}, os.Stdout)
	// if err != nil {
	// 	panic(err)
	// }

}

func Check()  {
	for _, artifact := range config.Cfg.Artifacts {
		sha256sum, err := installer.Sha256File(config.WorkingDir + "/" + artifact.Path)
		if err != nil {
			panic(err)
		}
		if sha256sum != artifact.Hash {
			panic(fmt.Errorf("build result [%s] hash [%s] not match the artifact hash [%s]", artifact.Path, sha256sum[:8], artifact.Hash[:8]))			
		}
		fmt.Printf("[OK]: %s=%s\n", artifact.Path, sha256sum[:8])
	}
}