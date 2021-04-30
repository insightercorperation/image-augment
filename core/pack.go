package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var augmentDirName = "augment"

type Pack struct {
	Name        string
	RootAbsPath string
	JSONPath    string
	ImagePath   string
}

func ReadPack(rootPath string, pack *Pack) error {

	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return err
	}

	pattern := filepath.Join(absPath, "*.json")
	jsonPath, err := filepath.Glob(pattern)
	if err != nil {
	}

	if len(jsonPath) != 1 {
		return errors.New("only one json file allowed")
	}

	pattern = filepath.Join(absPath, "*.jpg")
	imagePath, err := filepath.Glob(pattern)
	if err != nil {
	}
	if len(imagePath) != 1 {
		return errors.New("only one image file allowed")
	}

	baseName := filepath.Base(absPath)
	pack.Name = baseName
	pack.RootAbsPath = absPath
	pack.JSONPath = jsonPath[0]
	pack.ImagePath = imagePath[0]

	return nil
}

func (p Pack) HasJSON() bool {
	if p.JSONPath != "" {
		return true
	} else {
		return false
	}
}

func (p Pack) HasImage() bool {
	if p.ImagePath != "" {
		return true
	} else {
		return false
	}
}

func (p Pack) ResizeImagePath(size int) string {
	name := fmt.Sprintf("%s_resize_%d.%s", p.Name, size, "jpg")
	path := filepath.Join(p.RootAbsPath, augmentDirName, name)
	return path
}

func (p Pack) ResizeBBoxJSONPath(size int) string {
	name := fmt.Sprintf("%s_resize_%d.bbox.%s", p.Name, size, "json")
	path := filepath.Join(p.RootAbsPath, augmentDirName, name)
	return path
}

func (p Pack) ResizeJSONPath(size int) string {
	name := fmt.Sprintf("%s_resize_%d.%s", p.Name, size, "json")
	path := filepath.Join(p.RootAbsPath, augmentDirName, name)
	return path
}

func (p Pack) CropImagePath(size int, annoID string) string {
	name := fmt.Sprintf("%s_resize_%d_crop_%s.%s", p.Name, size, annoID, "jpg")
	path := filepath.Join(p.RootAbsPath, augmentDirName, name)
	return path
}

func (p Pack) CreateAugmentDirIfNotExist() error {
	augmentDirPath := filepath.Join(p.RootAbsPath, augmentDirName)
	err := os.MkdirAll(augmentDirPath, os.ModePerm)
	return err
}
