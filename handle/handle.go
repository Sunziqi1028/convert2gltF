package handle

import (
	converter "convert2gltF/convert"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
	InFilepath string `json:"in_filepath"`
	OutDirPath string `json:"out_dir_path"`
}

func Convert(c *gin.Context) {
	cPath := new(Path)
	if err := c.BindJSON(cPath); err != nil {
		c.JSON(500, "post body 解析失败!")
		return
	}
	fmt.Println(cPath.InFilepath, "===", cPath.OutDirPath)
	err := cPath.Convert2glft()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, fmt.Sprintf("模型转换失败,err:%v", fmt.Sprintln(err)))
		return
	}
	c.JSON(200, "模型转换成功！")
	return
}

func (p Path) Convert2glft() error {
	outFilepath := handleFilepath(p.InFilepath, p.OutDirPath)
	if len(outFilepath) == 0 {
		return errors.New("目录格式不正确！")
	}
	err := converter.Convert2Gltf(p.InFilepath, outFilepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func handleFilepath(InFilepath, OutDirPath string) string {
	ext := filepath.Ext(InFilepath)                             // .FBX .obj
	basePath := filepath.Base(InFilepath)                       // xxx.FBX JS1902.obj
	tempFilepath := strings.TrimRight(basePath, ext)            // xxx     JS1902
	outFilename := tempFilepath + ".gltf"                       // xxx.gltf
	savePath := filepath.Join(OutDirPath, tempFilepath, "gltf") //  /xxx/xx/x/gltf
	if pathExists(savePath) == false {
		err := os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return ""
		}
	}
	outFilepath := filepath.Join(savePath, outFilename) // /xxx/xx/x/gltf/xxx.gltf
	return outFilepath
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
