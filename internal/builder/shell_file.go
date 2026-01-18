package builder

import (
	"bytes"
	"os"
	"text/template"

	"github.com/venlax/c_build/internal/config"
)

//TODO

const shellFileTmpl string = `#!/usr/bin/env bash
set -euo pipefail

IMAGE_NAME="{{ .Image }}"
CONTAINER_NAME="{{ .ContainerName }}"

WORKDIR="{{ .WorkDir }}"

PROJ_ROOT=""

REPROBUILD_DIR="{{ .ReprobuildDir }}"

REPROBUILD_PATH=""

usage() {
  cat <<EOF
Usage: $0 --proj_root=PATH --reprobuild_path=PATH [options]

Options:
  --proj_root=PATH        Project root directory (required)
  --reprobuild_path=PATH  Reprobuild full path
  -h, --help              Show this help message
EOF
}

for arg in "$@"; do
  case "$arg" in
    --proj_root=*)
      PROJ_ROOT="${arg#*=}"
      ;;
    --reprobuild_path=*)
      REPROBUILD_PATH="${arg#*=}"
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "Unknown option: $arg"
      usage
      exit 1
      ;;
  esac
done

# ---------- validation ----------
if [[ -z "$PROJ_ROOT" ]]; then
  echo "ERROR: --proj_root is required"
  usage
  exit 1
fi

if [[ ! -d "$PROJ_ROOT" ]]; then
  echo "ERROR: proj_root does not exist: $PROJ_ROOT"
  exit 1
fi

if [[ -z "$REPROBUILD_PATH" ]]; then
  echo "ERROR: --reprobuild_path is required"
  usage
  exit 1
fi

if [[ ! -d "$REPROBUILD_PATH" ]]; then
  echo "ERROR: reprobuild_path does not exist: $REPROBUILD_PATH"
  exit 1
fi


# ---------- sanity check ----------
if [ ! -f "./Dockerfile" ]; then
  echo "ERROR: Dockerfile not found at ${DOCKERFILE_PATH}" >&2
  exit 1
fi

# ---------- build ----------
echo "==> Building image ${IMAGE_NAME}"
docker build \
  --network host \
  --pull \
  -t reprobuild \
  .

# ---------- cleanup ----------
if docker ps -a --filter "name=^${CONTAINER_NAME}$" --format '{{"{{"}}.ID{{"}}"}}' | grep -q .; then
  docker rm -f "${CONTAINER_NAME}"
fi


# ---------- run ----------
echo "==> Running build container ${CONTAINER_NAME}"
docker run --rm \
  --name "${CONTAINER_NAME}" \
  --network host \
  -v "${PROJ_ROOT}:${WORKDIR}" \
  -v "${REPROBUILD_PATH}:${REPROBUILD_DIR}"\
  reprobuild

echo "==> Build finished."

`

type ShellfileTmplData struct {
	Image string
	ContainerName string
	WorkDir string
  ReprobuildDir string
}


func RenderShellfile(dstDir string, digest string) {
	tmpl, err := template.New("").Parse(shellFileTmpl)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(dstDir + "/build.sh")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	
	err = tmpl.Execute(&buf, genShellfileData(digest))

	if err != nil {
		panic(err)
	}

	_, err = f.Write(buf.Bytes())

	if err != nil {
		panic(err)
	}
}

func genShellfileData(digest string) ShellfileTmplData {
	var data ShellfileTmplData
	data.Image = digest
	data.WorkDir = config.WorkingDir
	data.ContainerName = config.ContainerName
  data.ReprobuildDir = config.ReprobuildDir
	return data
}