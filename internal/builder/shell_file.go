package builder

//TODO

var shellFileTmpl string = `#!/usr/bin/env bash
set -euo pipefail

IMAGE_NAME="{{ .ImageName }}"
CONTAINER_NAME="{{ .ContainerName }}"

HOST_BUILD_ROOT="{{ .HostBuildRootDir }}"
WORKDIR="{{ .WorkingDir }}"


PULL="{{ .Pull }}"

# ---------- sanity check ----------
if [ ! -f "./Dockerfile" ]; then
  echo "ERROR: Dockerfile not found at ${DOCKERFILE_PATH}" >&2
  exit 1
fi

# ---------- build ----------
echo "==> Building image ${IMAGE_NAME}"
docker build \
  {{- if .Pull }} --pull {{- end }} \
  -t "${IMAGE_NAME}" 

# ---------- cleanup ----------
if docker ps -a --filter "name=^${CONTAINER_NAME}$" --format '{{.ID}}' | grep -q .; then
  docker rm -f "${CONTAINER_NAME}"
fi


# ---------- run ----------
echo "==> Running build container ${CONTAINER_NAME}"
docker run --rm \
  --name "${CONTAINER_NAME}" \
  --network host \
  -v "${HOST_BUILD_ROOT}:${WORKDIR}" \
  {{- range .ExtraArgs }}
  {{ . }} \
  {{- end }}
  "${IMAGE_NAME}"

echo "==> Build finished."

`

type ShellfileTmplData struct {

}