package core

import (
	"log"
	"testing"
)

type CocoTestData struct {
	viennaCode  string
	description string
	contributor string
	licenseURL  string
	categories  []string
}

var cocoTestData = CocoTestData{
	viennaCode:  "010102",
	description: "2021 인사이터 비엔나코드 010102 상표이미지 데이터 셋",
	contributor: "(주)인사이터",
	licenseURL:  "https://creativecommons.org/publicdomain/zero/1.0/",
	categories:  []string{"불완전한(미완성의) 별"},
}

type CocoImageTest struct {
	fileName string
	height   int
	width    int
	id       string
}

var cocoImageTestData = CocoImageTest{
	fileName: "60606_resize.png",
	height:   416,
	width:    416,
	id:       "60606",
}

func TestDefaultCoco(t *testing.T) {

	coco := NewCoco(cocoTestData.viennaCode, cocoTestData.categories)

	if coco.Info.Description != cocoTestData.description {
		t.Errorf("%s not eqal %s", cocoTestData.description, coco.Info.Description)
	}
	if coco.Info.Contributor != cocoTestData.contributor {
		t.Errorf("%s not eqal %s", cocoTestData.contributor, coco.Info.Contributor)
	}
	if coco.Licenses[0].URL != cocoTestData.licenseURL {
		t.Errorf("%s not eqal %s", cocoTestData.licenseURL, coco.Licenses[0].URL)
	}
	if coco.Categories[0].Name != cocoTestData.categories[0] {
		t.Errorf("%s not eqal %s", cocoTestData.categories[0], coco.Categories[0].Name)
	}
}

func TestCocoUpdateImage(t *testing.T) {

	coco := NewCoco(cocoTestData.viennaCode, cocoTestData.categories)

	coco.UpdateImage(cocoImageTestData.fileName, cocoImageTestData.id, cocoImageTestData.height, cocoImageTestData.width)

	if coco.Images[0].FileName != cocoImageTestData.fileName {
		t.Errorf("%s not eqal %s", cocoImageTestData.fileName, coco.Images[0].FileName)
	}
	if coco.Images[0].Height != cocoImageTestData.height {
		t.Errorf("%d not eqal %d", cocoImageTestData.height, coco.Images[0].Height)
	}
	if coco.Images[0].Width != cocoImageTestData.width {
		t.Errorf("%d not eqal %d", cocoImageTestData.width, coco.Images[0].Width)
	}
	if coco.Images[0].Id != cocoImageTestData.id {
		t.Errorf("%s not eqal %s", cocoImageTestData.id, coco.Images[0].Id)
	}
}

func TestCocoUpdateAnnotation(t *testing.T) {

	coco := NewCoco(cocoTestData.viennaCode, cocoTestData.categories)

	var firstLabelData LabelData
	path := "./fixture/870285/augment/870285_resize_416.bbox.json"
	err := ReadLabelData(path, &firstLabelData)
	if err != nil {
		log.Fatal(err)
	}

	var secondLabelData LabelData
	path = "./fixture/1270997/augment/1270997_resize_416.bbox.json"
	err = ReadLabelData(path, &secondLabelData)
	if err != nil {
		log.Fatal(err)
	}

	coco.UpdateAnnotation("870285", 416*416, &firstLabelData)
	coco.UpdateAnnotation("1270997", 416*416, &secondLabelData)

	firstTestBBox := []float64{
		159.97,
		87.61,
		131.07,
		227.29,
	}

	for idx, value := range coco.Annotations[0].BBox {
		if value != firstTestBBox[idx] {
			t.Errorf("%+v not eqal %+v", firstTestBBox, coco.Annotations[0].BBox)
		}
	}

	secondTestBBox := []float64{
		-14.49,
		-9.72,
		255.02,
		420.39,
	}

	for idx, value := range coco.Annotations[1].BBox {
		if value != secondTestBBox[idx] {
			t.Errorf("%+v not eqal %+v", secondTestBBox, coco.Annotations[1].BBox)
		}
	}

}

func TestGetCocoAnnotations(t *testing.T) {

	var coco Coco
	path := "./fixture/coco_annotation.json"
	ReadCocoData(path, &coco)

	annotations := coco.getCocoAnnotations("4019840004451")

	if len(annotations) != 2 {
		t.Errorf("%+v 해당 아이디의 annotation 개수가 일치 하지 않음", annotations)
	}
}

func TestInjectNewImageId(t *testing.T) {
	var coco Coco
	path := "./fixture/coco_annotation.json"
	ReadCocoData(path, &coco)

	annotations := coco.getCocoAnnotations("4019840004451")
	newAnnotations := injectNewImageId(annotations, "1000457")

	if annotations[0].ImageId == newAnnotations[0].ImageId {
		t.Errorf("%+v 값이 새로운 값 %+v 로 변하지 않음", annotations, newAnnotations)
	}
}

func TestDeleteCocoAnnotation(t *testing.T) {

	var coco Coco
	path := "./fixture/coco_annotation.json"
	ReadCocoData(path, &coco)

	coco.deleteAnnotations("4019840004451")

	if len(coco.Annotations) != 2 {
		t.Errorf("%+v에서 정상적으로 삭제되지 않음", coco.Annotations)
	}
}

func TestCocoOverrideAnnotation(t *testing.T) {
	var coco Coco

	path := "./fixture/coco_annotation.json"
	ReadCocoData(path, &coco)

	beforeOverride := len(coco.Annotations)
	coco.OverrideAnnotation("4019840004451", "1000457")
	afterOverride := len(coco.Annotations)

	if beforeOverride == afterOverride {
		t.Errorf("%+v", coco.Annotations)
	}
}
