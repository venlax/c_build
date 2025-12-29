package docker

import (
	"context"
	"io"
	"os"
)

func ReadFileFromContainer(abSrcPath string) ([]byte, error) {
	reader, _, err := Cli.CopyFromContainer(context.Background(), containerID, abSrcPath)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	content, err := io.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func CopyFileFromContainer(abSrcPath string, hostDestPath string) error {
	reader, _, err := Cli.CopyFromContainer(context.Background(), containerID, abSrcPath)
    if err != nil {
        return err
    }
    defer reader.Close()

    f, err := os.Create(hostDestPath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    _, err = io.Copy(f, reader)
    if err != nil {
        return err
    }

	return nil
}