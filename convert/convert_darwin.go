//go:build darwin
// +build darwin

package converter

import (
	"bytes"
	"os/exec"
)

func init() {
	var stdout bytes.Buffer
	cmd := exec.Command("brew -v")
	cmd.Stdout = &stdout // 校验本地是否有 brew的版本

	cmd.Run()

	if len(stdout.Bytes()) == 0 {
		sh := "./script/install.sh"
		chmod := exec.Command("cd ../script", "&&", "chmod", "-R", "755", "install.sh")
		chmod.Run()
		cmdInstBrew := exec.Command("/bin/bash", "-c", sh)
		//fmt.Println(cmdInstBrew)
		cmdInstBrew.Run()
	}
}

func Convert2Gltf(inpath, outpath string) error {
	cmd := exec.Command("assimp", "export", inpath, outpath)
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Wait()
	return nil
}
