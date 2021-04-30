package main

import (
	"fmt"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {

	applicationNumber := "7020190001189"
	samplePath := fmt.Sprintf("./project-25-publish/%s/%s.jpg", applicationNumber, applicationNumber)
	sampleImage, err := imaging.Open(samplePath)

	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	outputRootDir := fmt.Sprintf("./project-25-publish/%s/preprocess", applicationNumber)
	if _, err := os.Stat(outputRootDir); os.IsNotExist(err) {
		os.Mkdir(outputRootDir, 0755)
	}
	outputPath := fmt.Sprintf("./project-25-publish/%s/preprocess/%s_resize.jpg", applicationNumber, applicationNumber)
	outputImage := imaging.Resize(sampleImage, 200, 0, imaging.Lanczos)
	err = imaging.Save(outputImage, outputPath)

	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	outputPath = fmt.Sprintf("./project-25-publish/%s/preprocess/%s_invert.jpg", applicationNumber, applicationNumber)
	outputImage = imaging.Invert(sampleImage)
	err = imaging.Save(outputImage, outputPath)

	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	outputPath = fmt.Sprintf("./project-25-publish/%s/preprocess/%s_sharpen.jpg", applicationNumber, applicationNumber)
	outputImage = imaging.Sharpen(sampleImage, 2)
	err = imaging.Save(outputImage, outputPath)

	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
