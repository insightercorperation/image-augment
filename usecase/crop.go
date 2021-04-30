package usecase

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/cheggaaa/pb/v3"
	"github.com/insightercorperation/image-augment/core"
)

func Crop(parentDir string, size int) {

	pattern := filepath.Join(parentDir, "*")
	packRootPaths, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println(err.Error())
	}

	count := len(packRootPaths)
	bar := pb.StartNew(count)

	wg := new(sync.WaitGroup)
	wg.Add(len(packRootPaths))

	for _, packRootPath := range packRootPaths {
		go cropImage(wg, bar, packRootPath, size)
	}

	wg.Wait()
	bar.Finish()
}

func cropImage(wg *sync.WaitGroup, bar *pb.ProgressBar, packRootPath string, size int) {

	defer wg.Done()

	var pack core.Pack
	core.ReadPack(packRootPath, &pack)

	err := pack.CreateAugmentDirIfNotExist()
	if err != nil {
		fmt.Println(err)
		return
	}

	var labelData core.LabelData
	core.ReadLabelData(pack.ResizeBBoxJSONPath(size), &labelData)

	for _, labels := range labelData.Labels {
		for _, label := range labels {
			if label.Name != "제거" {
				rect := label.Anno.Rect()
				cropedImage := core.Crop(pack.ResizeImagePath(size), rect)
				cropImagePath := pack.CropImagePath(size, label.Anno.ID)
				core.SaveImage(cropImagePath, cropedImage)
			}
		}
	}

	bar.Increment()
}
