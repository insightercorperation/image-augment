package usecase

import (
	"fmt"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
	"github.com/insightercorperation/image-augment/core"
)

func Resize(parentDir string, target string, size int, remainPolygon bool) {

	if target != "all" {
		return
	}

	pattern := filepath.Join(parentDir, "*")
	packRootPaths, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println(err.Error())
	}

	count := len(packRootPaths)
	bar := pb.StartNew(count)

	for _, packRootPath := range packRootPaths {
		resizeData(bar, packRootPath, size, remainPolygon)
	}

	bar.Finish()
}

func resizeData(bar *pb.ProgressBar, packRootPath string, size int, remainPolygon bool) {

	var pack core.Pack
	core.ReadPack(packRootPath, &pack)

	err := pack.CreateAugmentDirIfNotExist()
	if err != nil {
		fmt.Println(err)
		return
	}

	originalImage := core.LoadImage(pack.ImagePath)
	if originalImage == nil {
		fmt.Printf("No OriginalIamge %s", pack.ImagePath)
		return
	}

	imageSize, err := core.GetImageSize(originalImage)
	if err != nil {
		fmt.Println(err)
		return
	}

	resizedImage := core.ImageResize(originalImage, size)
	if resizedImage != nil {
		core.SaveImage(pack.ResizeImagePath(size), resizedImage)
	}

	var labelData core.LabelData
	core.ReadLabelData(pack.JSONPath, &labelData)

	ratioHorizontal := float64(size) / float64(imageSize.X)
	ratioVertical := float64(size) / float64(imageSize.Y)

	if !remainPolygon {
		labelData.ConvertToBBox()
		labelData.ScalePoints(ratioHorizontal, ratioVertical)
		core.SaveLabelData(pack.ResizeBBoxJSONPath(size), labelData)
	} else {
		labelData.ScalePoints(ratioHorizontal, ratioVertical)
		core.SaveLabelData(pack.ResizeJSONPath(size), labelData)
	}

	bar.Increment()
}
