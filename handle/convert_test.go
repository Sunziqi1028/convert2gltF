package handle

import (
	converter "convert2gltF/convert"
	"testing"
)

// .FBX .3ds .glb .obj to gltf
func TestConvert(t *testing.T) {
	//InPath := "/Users/Sun/Downloads/Hei/Hei.FBX"
	//OutPath := "/Users/Sun/Downloads"

	//InPath := "/Users/Sun/Downloads/model/JS1901.obj"
	//OutPath := "/Users/Sun/Downloads"

	//InPath := "/Users/Sun/Downloads/model/JS1901.glb"
	//OutPath := "/Users/Sun/Downloads"

	InPath := "/Users/Sun/Downloads/model/JS1901.3ds"
	OutPath := "/Users/Sun/Downloads"
	var cPath = Path{
		InFilepath: InPath,
		OutDirPath: OutPath,
	}
	err := cPath.Convert2glft()
	if err != nil {
		t.Error(err)
	}
}

func TestInit(t *testing.T) {
	converter.Init()

}
