package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/venlax/c_build/internal/builder"
	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/docker"
	"github.com/venlax/c_build/internal/installer"
)

var	debug bool = false

func main() {
	args := os.Args
	create := false
	configPath := ""
	dstDirPath := "./build" // default gen dockerfile in the project build dir
	for _, arg := range args {
		if arg == "-c" || arg == "--create" {
			create = true
		}
		if arg == "-d" || arg == "--debug" {
			debug = true
		}
		if strings.HasPrefix(arg,"--input") {
			fmt.Sscanf(arg, "--input=%s", &configPath)	
		}
		if strings.HasPrefix(arg, "--output") {
			fmt.Sscanf(arg, "--output=%s", &dstDirPath)
		}
	}

	config.Init(configPath)

	if !debug {
		builder.RenderDockerfile(dstDirPath)
		return
	}

	docker.Init(create)

	installer.Init()
	installer.Install()
	
	builder.Build()
	builder.Check()
}

// import (
// 	"context"
// 	"fmt"

// 	"github.com/docker/docker/api/types/container"
// 	"github.com/docker/docker/client"
// )
        
// func main()  {
//     // 创建 Docker 客户端
// 	cli, err := client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		panic(err)
// 	}

// 	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All : true})
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, container := range containers {
// 		fmt.Printf("ID: %s, Image: %s, Status: %s\n", container.ID[:12], container.Image, container.Status, container.Names)
// 	}
// }