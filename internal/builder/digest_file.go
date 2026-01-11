package builder

import (
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/docker"
	"gopkg.in/yaml.v3"
)

type DigestFileData struct {
	Digest string `yaml:"digest"`
	ConfigHash string `yaml:"config_hash"`
	ImmutabilityHash string `yaml:"immu_hash"`
}

func RenderDigestFile(dstDir string, configPath string) {
	var data DigestFileData
	image := docker.GetImageInspect(config.Cfg.MetaData.Distribution)
	data.Digest = image.RepoDigests[0]
	tmp, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(tmp)
	data.ConfigHash = hex.EncodeToString(sum[:])
	sum = sha256.Sum256([]byte(data.Digest + data.ConfigHash))
	data.ImmutabilityHash = hex.EncodeToString(sum[:])
	out, err := yaml.Marshal(data)
	if err != nil {
		panic(err)
	}
	os.WriteFile(dstDir + "/digest.yaml", out, 0644)
}

func GetDigestWithCheck(dstDir string, configPath string) string {
	file_data, err := os.ReadFile(dstDir + "/digest.yaml")
	if err != nil {
		panic(err)
	}
	var data DigestFileData
	if err := yaml.Unmarshal(file_data, &data); err != nil {
		panic(err)
	}
	tmp, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(tmp)
	if data.ConfigHash != hex.EncodeToString(sum[:]) {
		panic("Run c_build with --debug after change build_record.yaml and do not change digest.yaml.")
	}
	sum = sha256.Sum256([]byte(data.Digest + data.ConfigHash))
	if data.ImmutabilityHash != hex.EncodeToString(sum[:]) {
		panic("Run c_build with --debug after change build_record.yaml and do not change digest.yaml.")
	}
	return data.Digest
}