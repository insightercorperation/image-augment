package core

import (
	"testing"
)

type TestPackData struct {
	packname string
	hasImage bool
	hasJSON  bool
}

var testpackdata = TestPackData{
	packname: "870285",
	hasImage: true,
	hasJSON:  true,
}

func TestSample(t *testing.T) {
	packRootPath := "./fixture/870285"

	var pack Pack
	ReadPack(packRootPath, &pack)

	if pack.Name != testpackdata.packname {
		t.Errorf("%s not equal %s", pack.Name, testpackdata.packname)
	}

	if pack.Name != testpackdata.packname {
		t.Errorf("%s not equal %s", pack.Name, testpackdata.packname)
	}

	if pack.HasJSON() != testpackdata.hasJSON {
		t.Errorf("%s not equal %t", pack.JSONPath, testpackdata.hasJSON)
	}

	if pack.HasImage() != testpackdata.hasImage {
		t.Errorf("%s not equal %t", pack.ImagePath, testpackdata.hasImage)
	}

}
