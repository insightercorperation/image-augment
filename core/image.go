package core

import (
	"image"

	"github.com/disintegration/imaging"
)

func LoadImage(imagePath string) image.Image {
	img, err := imaging.Open(imagePath)
	if err != nil {
		return nil
	}
	return img
}

func GetImageSize(image image.Image) (image.Point, error) {
	return image.Bounds().Size(), nil
}

func ImageResize(image image.Image, size int) *image.NRGBA {
	resizedImage := imaging.Resize(image, size, size, imaging.Lanczos)
	return resizedImage
}

func SaveImage(imagePath string, image *image.NRGBA) {
	imaging.Save(image, imagePath)
}

func Crop(imagePath string, rect image.Rectangle) *image.NRGBA {
	inputImage, err := imaging.Open(imagePath)
	if err != nil {
		return nil
	}
	return imaging.Crop(inputImage, rect)
}
