package main

import (
	"os"

	"github.com/venlax/c_build/internal/builder"
	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/docker"
	"github.com/venlax/c_build/internal/installer"
)

func main() {
	args := os.Args
	create := false
	for _, arg := range args {
		if arg == "-c" || arg == "--create" {
			create = true
		}
	}

	config.Init()
	docker.Init(create)
	installer.Init()
	// docker.Run([]string{"apt", "list", "-a", "nginx"})
	installer.Install()
	builder.Build()
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