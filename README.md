## Usage

### Basic Usage

```bash
# pre install docker in the build host
# when new create or recreate pre-run
docker stop $containerName
docker rm $containerName
# main
go run cmd/c_build/*.go [-/--args[=value]]
```  
  
|Args     | HasValue($value) | Explain|  
|---------|-------   |-------|
|-c/--create| false| 创建该镜像下对应容器名(config.ContainerName)的容器,用于首次创建容器|
|-d/--debug| false | 调试模式, 所有命令直接通过Go程序与容器交互, 不生成Dockerfile, 仅用于调试|
|--input| true| $value为yaml配置文件路径|
|--output| true| $value为Dockerfile/build.sh目录路径,默认为`./build`|
|--log_level|true|$value为日志级别, 可选值为debug, info, error,默认为info|
|--ld_preload|true|$value为reprobuild在宿主机的路径, 目前只需要保证该目录下有c_build所需的动态链接库即可|
### Result
- Dockerfile
- build.sh

### Producing Env (Reproduce Build with the Dockerfile/build.sh)
```shell
# pre install docker in the build host

sudo chmod +x build.sh

./build.sh --proj_root=XXX --reprobuild_path=XXXX
```

## Note   
首次运行c_build请携带参数-c/--create, 容器名在[internal/config/docker.go](internal/config/docker.go)设置, 在携带-d/--debug情况下,所有命令直接通过Go程序与容器交互, 不生成Dockerfile,用于验证最终一致性以及yaml文件配置在当前环境的可行性, 并且生成[build/digest.yaml](build/digest.yaml)用于非debug模式下的校验和版本锁定(不要自主修改digest.yaml), 当需要投入生产环境大量在容器重复构建时, 在debug验证了可行性情况下, 运行c_build(without -d/--debug), 将会在build目录下生成Dockerfile/build.sh,携带这两个文件在生产环境按照`Producing Env`构建即可