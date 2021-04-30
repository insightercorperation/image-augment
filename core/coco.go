package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

type Coco struct {
	Info        Information      `json:"info"`
	Licenses    []License        `json:"licenses"`
	Images      []CocoImage      `json:"images"`
	Annotations []CocoAnnotation `json:"annotations"`
	Categories  []CocoCategory   `json:"categories"`
}

type Information struct {
	Description string `json:"description"`
	URL         string `json:"url"`
	Version     string `json:"version"`
	Year        int    `json:"year"`
	Contributor string `json:"contributor"`
	DateCreated string `json:"dateCreated"`
}

type License struct {
	URL  string `json:"url"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CocoImage struct {
	FileName string `json:"fileName"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	Id       string `json:"id"`
}

type CocoAnnotation struct {
	Area       float64   `json:"area"`
	IsCrowd    bool      `json:"iscrowd"`
	ImageId    string    `json:"image_id"`
	BBox       []float64 `json:"bbox"`
	CategoryId int       `json:"category_id"`
	ID         string    `json:"id"`
}

type CocoCategory struct {
	ID            int    `json:"id"`
	SuperCategory string `json:"supercategory"`
	Name          string `json:"name"`
}

func NewCoco(viennaCode string, categories []string) Coco {
	coco := Coco{}

	year := 2021
	coco.Info = Information{
		Description: fmt.Sprintf("%d 인사이터 비엔나코드 %s 상표이미지 데이터 셋", year, viennaCode),
		URL:         "",
		Version:     "1.0",
		Year:        year,
		Contributor: "(주)인사이터",
		DateCreated: "2021/01/22",
	}

	coco.Licenses = []License{
		{
			ID:   1,
			URL:  "https://creativecommons.org/publicdomain/zero/1.0/",
			Name: "CC0 1.0 Universal (CC0 1.0) Public Domain Dedication License",
		},
	}

	cocoCateroies := make([]CocoCategory, len(categories))
	for idx, category := range categories {
		cocoCateroies[idx] = CocoCategory{
			ID:            idx + 1,
			SuperCategory: category,
			Name:          category,
		}
	}
	coco.Categories = cocoCateroies

	return coco
}

func (coco *Coco) UpdateImage(fileName string, id string, height int, width int) {

	coco.Images = append(coco.Images, CocoImage{
		FileName: fileName,
		Height:   height,
		Width:    width,
		Id:       id,
	})
}

func round2(value float64) float64 {
	return math.Round(value*100) / 100
}

func extractBBox(label Label) []float64 {

	if label.Anno.Type == "bbox" {
		return []float64{
			round2(label.Anno.Points[0].Lng),
			round2(label.Anno.Points[0].Lat),
			round2(label.Anno.Points[1].Lng - label.Anno.Points[0].Lng),
			round2(label.Anno.Points[1].Lat - label.Anno.Points[0].Lat),
		}
	} else {
		return make([]float64, 0)
	}

}

func matchCategoryId(labelName string, categories []CocoCategory) int {
	return 1
}

func (coco *Coco) UpdateAnnotation(imageId string, area float64, labelData *LabelData) {

	for _, labels := range labelData.Labels {
		for _, label := range labels {
			coco.Annotations = append(coco.Annotations, CocoAnnotation{
				Area:       area,
				IsCrowd:    false,
				ImageId:    imageId,
				BBox:       extractBBox(label),
				CategoryId: matchCategoryId(label.Name, coco.Categories),
				ID:         label.Anno.ID,
			})
		}
	}
}

func remove(s []CocoAnnotation, i int) []CocoAnnotation {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func filter(vs []CocoAnnotation, f func(CocoAnnotation) bool) []CocoAnnotation {
	vsf := make([]CocoAnnotation, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func (coco *Coco) deleteAnnotations(imageId string) {

	newAnnotations := filter(coco.Annotations, func(anno CocoAnnotation) bool {
		return anno.ImageId != imageId
	})

	coco.Annotations = newAnnotations
}

func (coco *Coco) OverrideAnnotation(representId string, targetId string) {

	representAnnotations := coco.getCocoAnnotations(representId)

	if len(representAnnotations) == 0 {
		return
	}

	newAnnotations := injectNewImageId(representAnnotations, targetId)
	coco.deleteAnnotations(targetId)

	for _, anno := range newAnnotations {
		coco.Annotations = append(coco.Annotations, anno)
	}
}

func (coco *Coco) getCocoAnnotations(imageId string) []CocoAnnotation {

	annotations := make([]CocoAnnotation, 0)
	for _, anno := range coco.Annotations {
		if anno.ImageId == imageId {
			annotations = append(annotations, anno)
		}
	}
	return annotations
}

func createAnnotationId() string {
	uui := uuid.Must(uuid.NewV4()).String()
	newUUID := strings.Replace(uui, "-", "", -1)[0:22]
	return newUUID
}

func injectNewImageId(annotations []CocoAnnotation, imageId string) []CocoAnnotation {

	newAnnotations := make([]CocoAnnotation, 0)

	for _, anno := range annotations {
		anno.ImageId = imageId
		anno.ID = createAnnotationId()
		newAnnotations = append(newAnnotations, anno)
	}
	return newAnnotations
}

func SaveCoco(jsonPath string, coco Coco) {
	prettyJSON, err := json.MarshalIndent(coco, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(jsonPath, prettyJSON, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadCocoData(jsonPath string, coco *Coco) error {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &coco)
	jsonFile.Close()

	return nil
}
