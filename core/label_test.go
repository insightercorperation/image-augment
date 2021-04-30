package core

import (
	"log"
	"reflect"
	"testing"
)

type TestData struct {
	link         string
	originalName string
	labels       map[string][]Label
}

var testdata = TestData{
	link:         "https://markcloud-annotaion-images.s3.ap-northeast-2.amazonaws.com/611525.jpg",
	originalName: "611525",
	labels: map[string][]Label{
		"0xsl93ixs": []Label{
			Label{
				Name: "불완전한(미완성의) 별",
				Anno: Annotation{
					ID:   "udf4vu9r2xfupce3zwdp1",
					Type: "bbox",
					Points: []Point{
						Point{
							Lat: -89.45263237915378,
							Lng: 618.7274796756423,
						},
						Point{
							Lat: 1241.5167630373503,
							Lng: 1779.605820087373,
						},
					},
				},
			},
		},
		"ik4yg7xlq": []Label{
			Label{
				Name: "불완전한(미완성의) 별",
				Anno: Annotation{
					ID:   "1f60fijo55xtjfuwknp4r",
					Type: "polygon",
					Points: []Point{
						Point{
							Lat: 310.4687456302902,
							Lng: 690.8878430608253,
						},
						Point{
							Lat: 460.99504630239477,
							Lng: 664.504244415104,
						},
						Point{
							Lat: 458.3542340099017,
							Lng: 470.5847943690536,
						},
						Point{
							Lat: 278.7789981203734,
							Lng: 465.3080746399094,
						},
						Point{
							Lat: 132.21391588700845,
							Lng: 615.69458692052,
						},
						Point{
							Lat: 128.25269744826886,
							Lng: 839.9551754091497,
						},
						Point{
							Lat: 211.4382846618003,
							Lng: 846.55107507058,
						},
					},
				},
			},
		},
		"73em6szar": []Label{},
	},
}

func TestLabelData(t *testing.T) {
	var labelData LabelData
	path := "./fixture/611525.json"
	err := ReadLabelData(path, &labelData)
	if err != nil {
		log.Fatal(err)
	}

	if testdata.link != labelData.Link {
		t.Errorf("%s not eqal %s", testdata.link, labelData.Link)
	}

	if testdata.originalName != labelData.OriginalName {
		t.Errorf("%s not eqal %s", testdata.originalName, labelData.OriginalName)
	}

	for labelID, labels := range testdata.labels {

		testAnno := labels
		fixtureAnno := labelData.Labels[labelID]

		if !reflect.DeepEqual(testAnno, fixtureAnno) {
			t.Errorf("%+v not equal %+v", testAnno, fixtureAnno)
		}
	}
}

type TestPointsData struct {
	polygon []Point
	bbox    []Point
}

var testPointsData = TestPointsData{
	polygon: []Point{
		Point{
			Lat: 310.4687456302902,
			Lng: 690.8878430608253,
		},
		Point{
			Lat: 460.99504630239477,
			Lng: 664.504244415104,
		},
		Point{
			Lat: 458.3542340099017,
			Lng: 470.5847943690536,
		},
		Point{
			Lat: 278.7789981203734,
			Lng: 465.3080746399094,
		},
		Point{
			Lat: 132.21391588700845,
			Lng: 615.69458692052,
		},
		Point{
			Lat: 128.25269744826886,
			Lng: 839.9551754091497,
		},
		Point{
			Lat: 211.4382846618003,
			Lng: 846.55107507058,
		},
	},
	bbox: []Point{
		Point{
			Lat: 128.25269744826886,
			Lng: 465.3080746399094,
		},
		Point{
			Lat: 460.99504630239477,
			Lng: 846.55107507058,
		},
	},
}

func TestPolygonToBBox(t *testing.T) {

	convertedBboxPoints := polygonToBBox(testPointsData.polygon)
	if !reflect.DeepEqual(testPointsData.bbox, convertedBboxPoints) {
		t.Errorf("%+v not eqal %+v", testPointsData.bbox, convertedBboxPoints)
	}

}
