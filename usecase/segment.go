package usecase

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"

	"github.com/insightercorperation/image-augment/util"
)

func Segment(parentDir string, outputDir string, trainingRatio int, validationRatio int, testRatio int, hasSample bool) {

	pattern := filepath.Join(parentDir, "*")
	packRootPaths, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println(err.Error())
	}

	totalRatio := trainingRatio + validationRatio + testRatio
	dataSize := len(packRootPaths)

	trainSize := makeValueOfRatio(dataSize, totalRatio, trainingRatio)
	validationSize := makeValueOfRatio(dataSize, totalRatio, (validationRatio + testRatio))

	if validationSize%2 != 0 {
		trainSize = trainSize - 1
		validationSize = validationSize + 1
	}

	validationSize = validationSize / 2

	trainData := packRootPaths[:trainSize]
	validationData := packRootPaths[trainSize : trainSize+validationSize]
	testData := packRootPaths[trainSize+validationSize:]

	viennaCode := path.Base(parentDir)
	copyData(outputDir, "train", viennaCode, trainData)
	copyData(outputDir, "validation", viennaCode, validationData)
	copyData(outputDir, "test", viennaCode, testData)

	if hasSample {
		sampleData := append(trainData, validationData...)[:3]
		copyData(outputDir, "sample", viennaCode, sampleData)
	}
}

func copyData(outputDir string, dataSegment string, viennaCode string, targetDataPaths []string) {

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	segmentDirPath := filepath.Join(outputDir, dataSegment)
	err = os.MkdirAll(segmentDirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	packParentDirPath := filepath.Join(segmentDirPath, viennaCode)
	err = os.MkdirAll(packParentDirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	for idx, segmentDataPath := range targetDataPaths {
		copyAndCreateDir(packParentDirPath, segmentDataPath, idx)
	}

}

func copyAndCreateDir(segmentDirPath string, segmentDataPath string, idx int) {

	dirName := path.Base(segmentDataPath)
	newSegmentDataPath := filepath.Join(segmentDirPath, dirName)
	err := os.MkdirAll(newSegmentDataPath, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = util.CopyDirectory(segmentDataPath, newSegmentDataPath)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func makeValueOfRatio(N int, totalRatio int, ratio int) int {
	value := (int)(math.Round(((float64)(ratio) / (float64)(totalRatio)) * (float64)(N)))
	return value
}
