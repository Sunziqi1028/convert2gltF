//go:build windows
// +build windows

package converter

import (
	"os"
	"os/exec"
)

func Convert2Gltf(inpath, outpath string) error {
	var abPath string
	abPath, _ = os.Getwd()
	command := abPath + "\\bin\\assimp.exe"
	cmd := exec.Command(command, "export", inpath, outpath)

	err := cmd.Start()
	if err != nil {
		return err
	}

	cmd.Wait()
	return nil
}
