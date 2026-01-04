package installer

import (
	"log/slog"
	"strings"

	"github.com/venlax/c_build/internal/config"
	"github.com/venlax/c_build/internal/docker"
)

func Check(libInfo config.LibInfo) bool {
	pkg_mgr := GetPkgMgr("dpkg")

	if version := (&pkg_mgr).runGetLibVersion(libInfo); version != libInfo.Version {
		return false;
	}

	if sha256sum, err := Sha256File(libInfo.Path); err == nil {
		if sha256sum != libInfo.Sha256 {
			return false;
		}
	} else {
		panic(err)
	}

	return true;
}

func Sha256File(path string) (string, error) {

	var sb strings.Builder

	err := docker.Run([]string {"sha256sum", path}, &sb)

	if err != nil {
		slog.Error(sb.String())
		return "", err
	}

	return sb.String()[:64], nil
	

	// TODO: think more about this and split the raw str plz



	// data, err := docker.ReadFileFromContainer(path)
	// if err != nil {
	// 	return "", err
	// }
	// sum := sha256.Sum256(data)
    // return hex.EncodeToString(sum[:]), nil

	// f, err := os.Open(path)
	// if err != nil {
	// 	return "", err
	// }
	// defer f.Close()

	// h := sha256.New()
	// if _, err := io.Copy(h, f); err != nil {
	// 	return "", err
	// }

	// sum := h.Sum(nil)
	// return fmt.Sprintf("%x", sum), nil
}
