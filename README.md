## Usage

### Basic Usage

```bash
# when new create or recreate pre-run
docker stop $containerName
docker rm $containerName
# main
go run cmd/c_build/main [-/--args[=value]]
```  
  
|Args     | HasValue($value) | Explain|  
|---------|-------   |-------|
|-c/--create| false| 创建该镜像下对应容器名(config.ContainerName)的容器,用于首次创建容器|
|-d/--debug| false | 调试模式, 所有命令直接通过Go程序与容器交互, 不生成Dockerfile, 仅用于调试|
|--input| true| $value为yaml配置文件路径|

## TODO
暂不支持生成Dockerfile,直接运行程序构建, 首次构建请携带参数-c/--create, 容器名在[internal/config/docker.go](internal/config/docker.go)设置