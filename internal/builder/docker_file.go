package builder

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/installer"
)


var dockerfileTmpl string = `FROM {{.Image}}

{{- if .Env }}
ENV {{ join .Env " \\\n    " }}
{{- end }}

WORKDIR {{.WorkDir}}

{{- if .InstallCmds }}
{{- range .InstallCmds }}
RUN {{.}}
{{- end }}
{{- end }}

CMD {{.BuildCmd}}
`

type DockerfileTmplData struct {
	Image string
	Env []string
	WorkDir string
	InstallCmds []string
	BuildCmd []string
}

func RenderDockerfile(dstDir string) {
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
	
	err = tmpl.Execute(&buf, genDockerfileData())

	_, err = f.Write(buf.Bytes())

	if err != nil {
		panic(err)
	}
} 

func genDockerfileData() DockerfileTmplData {
	var data DockerfileTmplData
	data.Image = config.Image
	data.WorkDir = config.WorkingDir
	data.Env = config.Env
	data.InstallCmds = installer.InstallStrs()
	data.BuildCmd = append(data.BuildCmd, "make")
	return data
}