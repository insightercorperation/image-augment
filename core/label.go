package core

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
)

type LabelData struct {
	Link         string             `json:"link"`
	OriginalName string             `json:"originalName"`
	Labels       map[string][]Label `json:"labels"`
}

type Label struct {
	Name string     `json:"name"`
	Anno Annotation `json:"annotations"`
}

type Annotation struct {
	ID     string  `json:"id"`
	Type   string  `json:"type"`
	Points []Point `json:"points"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func ReadLabelData(jsonPath string, labelData *LabelData) error {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &labelData)
	jsonFile.Close()

	return nil
}

func SaveLabelData(jsonPath string, labelData LabelData) {
	prettyJSON, err := json.MarshalIndent(labelData, "", "    ")
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

func (labelData *LabelData) ConvertToBBox() {
	for _, labels := range labelData.Labels {
		for key, label := range labels {

			if label.Anno.Type == "polygon" {

				label.Anno.Type = "bbox"
				label.Anno.Points = polygonToBBox(label.Anno.Points)

				labels[key] = label
			}
		}

	}
}

func (labelData *LabelData) ScalePoints(ratioh float64, ratiov float64) {
	for _, labels := range labelData.Labels {
		for key, label := range labels {

			numOfPoints := len(label.Anno.Points)

			rescaledPoints := make([]Point, numOfPoints)

			for idx, point := range label.Anno.Points {

				newPoint := Point{
					Lat: point.Lat * ratiov,
					Lng: point.Lng * ratioh,
				}
				rescaledPoints[idx] = newPoint
			}

			labels[key].Anno.Points = rescaledPoints

		}

	}
}

func (anno *Annotation) Rect() image.Rectangle {
	if anno.Type == "polygon" {
		return image.Rectangle{}
	}

	return image.Rectangle{
		Min: image.Point{
			X: int(anno.Points[0].Lng),
			Y: int(anno.Points[0].Lat),
		},
		Max: image.Point{
			X: int(anno.Points[1].Lng),
			Y: int(anno.Points[1].Lat),
		},
	}

}

func polygonToBBox(points []Point) []Point {

	var minX float64
	var minY float64
	var maxX float64
	var maxY float64

	minX = points[0].Lng
	minY = points[0].Lat

	maxY = points[0].Lng
	maxY = points[0].Lat

	for _, point := range points {
		if point.Lng <= minX {
			minX = point.Lng
		}

		if point.Lat <= minY {
			minY = point.Lat
		}

		if point.Lng >= maxX {
			maxX = point.Lng
		}

		if point.Lat >= maxY {
			maxY = point.Lat
		}
	}

	return []Point{
		Point{
			Lat: minY,
			Lng: minX,
		},
		Point{
			Lat: maxY,
			Lng: maxX,
		},
	}
}
