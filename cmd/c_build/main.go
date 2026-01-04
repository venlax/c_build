package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/venlax/c_build/internal/builder"
	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/docker"
	"github.com/venlax/c_build/internal/installer"
)

var debug bool = false

func main() {
	args := os.Args

	create := false
	configPath := ""
	dstDirPath := "./build" // default gen dockerfile in the project build dir
	logLevel := slog.LevelInfo
	for _, arg := range args {
		switch {
		case arg == "-c" || arg == "--create":
			create = true

		case arg == "-d" || arg == "--debug":
			debug = true

		case strings.HasPrefix(arg, "--input"):
			fmt.Sscanf(arg, "--input=%s", &configPath)

		case strings.HasPrefix(arg, "--output"):
			fmt.Sscanf(arg, "--output=%s", &dstDirPath)
		case strings.HasPrefix(arg, "--log_level"):
			var tmp string
			fmt.Sscanf(arg, "--log_level", &tmp)
			switch tmp {
			case "debug":
				logLevel = slog.LevelDebug
			case "info":
				break
			case "error":
				logLevel = slog.LevelError
			default:
				panic(fmt.Errorf("Unknown value [%s] for log_level.",tmp))
			}
		}
	}


	InitLogger(logLevel)

	slog.Debug("parsed args",
		"create", create,
		"debug", debug,
		"config", configPath,
		"output", dstDirPath,
	)

	// ---- load config ----
	slog.Info("init config", "path", configPath)
	config.Init(configPath)

	if !debug {
		slog.Info("render only mode", "output", dstDirPath)

		builder.RenderDockerfile(dstDirPath)
		builder.RenderShellfile(dstDirPath)

		slog.Info("render finished")
		return
	}

	// ---- full build pipeline ----
	slog.Info("init docker", "create", create)
	docker.Init(create)

	slog.Info("init installer")
	installer.Init()

	slog.Info("install dependencies")
	installer.Install()

	slog.Info("start build")
	builder.Build()

	slog.Info("check build result")
	builder.Check()

	slog.Info("build finished successfully")
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