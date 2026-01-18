package builder

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/installer"
)


const dockerfileTmpl string = `FROM {{.Image}}

{{- if .Env }}
ENV {{ join .Env " \\\n    " }}
{{- end }}

WORKDIR {{.WorkDir}}

{{- if .InstallCmds }}
{{- range .InstallCmds }}
RUN {{.}}
{{- end }}
{{- end }}

CMD ["/bin/sh", "-c", "{{ .BuildCmd }}"]
`

type DockerfileTmplData struct {
	Image string
	Env []string
	WorkDir string
	InstallCmds []string
	BuildCmd string
}

func RenderDockerfile(dstDir string, digest string) {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse(dockerfileTmpl)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(dstDir + "/Dockerfile")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, genDockerfileData(digest))

	if err != nil {
		panic(err)
	}

	_, err = f.Write(buf.Bytes())

	if err != nil {
		panic(err)
	}
}

func genDockerfileData(digest string) DockerfileTmplData {
	var data DockerfileTmplData
	data.Image = digest
	data.WorkDir = config.WorkingDir
	data.Env = config.Env
	data.InstallCmds = installer.InstallStrs()
	var ld_path string
	if config.HasCustom {
		ld_path = fmt.Sprintf("env LD_PRELOAD=%s/libreprobuild_interceptor.so LD_LIBRARY_PATH=\"%s/deps:$LD_LIBRARY_PATH\"", config.ReprobuildDir, config.WorkingDir)
	} else {
		ld_path = fmt.Sprintf("env LD_PRELOAD=%s/libreprobuild_interceptor.so", config.ReprobuildDir)
	}
	BuildCommand := fmt.Sprintf("umask %s && %s", config.Cfg.MetaData.Umask, config.BuildCmd)
	data.BuildCmd = "make clean && " + strings.ReplaceAll(BuildCommand, "&&", "&& " + ld_path)
	return data
}