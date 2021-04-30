package usecase

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/insightercorperation/image-augment/core"
)

func Coco(parentDir string, size int, category string, targetOnPolygon bool) {

	pattern := filepath.Join(parentDir, "*")
	packRootPaths, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println(err.Error())
	}

	outputDirRoot := filepath.Join(parentDir, "..")
	outputDirName := fmt.Sprintf("%s_coco", path.Base(parentDir))
	outputDirPath := filepath.Join(outputDirRoot, outputDirName)

	err = os.MkdirAll(outputDirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	viennaCode := path.Base(parentDir)
	categories := []string{category}
	coco := core.NewCoco(viennaCode, categories)

	count := len(packRootPaths)
	bar := pb.StartNew(count)

	for _, packRootPath := range packRootPaths {
		copyImageAndCreateCoco(bar, &coco, packRootPath, size, outputDirPath, targetOnPolygon)
	}

	bar.Finish()

	cocoFileName := "annotation.json"
	cocoFilePath := filepath.Join(outputDirPath, cocoFileName)
	core.SaveCoco(cocoFilePath, coco)
}

func copyImageAndCreateCoco(bar *pb.ProgressBar, coco *core.Coco, packRootPath string, size int, outputDirPath string, targetOnPolygon bool) {
	var pack core.Pack
	core.ReadPack(packRootPath, &pack)

	resizedImagePath := pack.ResizeImagePath(size)
	originalImage, err := os.Open(resizedImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer originalImage.Close()

	newFileName := path.Base(resizedImagePath)
	targetImagePath := filepath.Join(outputDirPath, newFileName)
	targetImage, err := os.Create(targetImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer targetImage.Close()

	_, err = io.Copy(targetImage, originalImage)
	if err != nil {
		log.Fatal(err)
	}

	if !targetOnPolygon {
		bboxJsonPath := pack.ResizeBBoxJSONPath(size)

		var labelData core.LabelData
		err := core.ReadLabelData(bboxJsonPath, &labelData)
		if err != nil {
			log.Fatal(err)
		}

		imageFileName := path.Base(resizedImagePath)
		imageID := strings.Split(imageFileName, "_")[0]

		coco.UpdateImage(imageFileName, imageID, size, size)
		coco.UpdateAnnotation(imageID, float64(size*size), &labelData)

	} else {
		fmt.Println("BBox 에 대해서만 coco 생성이 가능합니다.")
	}

	bar.Increment()
}
