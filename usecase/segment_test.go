package usecase

import (
	"math"
	"testing"
)

var trainingRatio int = 8
var validationRatio int = 1
var testRatio int = 1

type SegmentTestData struct {
	trainig    []string
	validation []string
	test       []string
	sample     []string
}

var inputdata = []string{
	"1", "2", "3", "3", "3",
	"1", "2", "3", "3", "3",
	"1", "2", "3", "3", "3",
}

func MakeValueOfRatio(N int, totalRatio int, ratio int) int {
	value := (int)(math.Round(((float64)(ratio) / (float64)(totalRatio)) * (float64)(N)))

	return value
}

func TestSegmentData(t *testing.T) {
	totalRatio := trainingRatio + validationRatio + testRatio
	dataSize := len(inputdata)

	trainSize := MakeValueOfRatio(dataSize, totalRatio, trainingRatio)
	validSize := MakeValueOfRatio(dataSize, totalRatio, (validationRatio + testRatio))

	if dataSize != (trainSize + validSize) {
		t.Errorf("%d not eqal %d", dataSize, (trainSize + validSize))
	}

	if validSize%2 != 0 {
		trainSize = trainSize - 1
		validSize = validSize + 1
	}

	validSize = validSize / 2

	train := inputdata[:trainSize]
	valid := inputdata[trainSize : trainSize+validSize]
	test := inputdata[trainSize+validSize:]
	sample := append(train, valid...)

	var segmentTestData = SegmentTestData{
		trainig:    train,
		validation: valid,
		test:       test,
		sample:     sample,
	}

	if len(segmentTestData.trainig) != trainSize {
		t.Errorf("%d not eqal %d", len(segmentTestData.trainig), trainSize)
	}
	if len(segmentTestData.validation) != validSize {
		t.Errorf("%d not eqal %d", len(segmentTestData.validation), validSize)
	}
	if len(segmentTestData.test) != validSize {
		t.Errorf("%d not eqal %d", len(segmentTestData.test), validSize)
	}
	if len(segmentTestData.sample) != (trainSize + validSize) {
		t.Errorf("%d not eqal %d", len(segmentTestData.sample), trainSize+validSize)
	}
}
