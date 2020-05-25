package utils

import (
	"os"
	"testing"
)

func TestZipx_Zip(t *testing.T) {

	err := Zipx.Zip("~/Code/data/mm_wiki/attachment/1/4/", "/Users/dombine/Downloads/test.zip")
	if err != nil {
		t.Error(err)
	}
}

func TestZipx_Compress(t *testing.T) {

	osFiles := []*os.File{}
	f3, err := os.Open("/Users/dombine/Code/data/mm_wiki/images/1/4/redis2.jpeg")
	if err != nil {
		t.Error(err)
	}
	defer f3.Close()
	osFiles = append(osFiles, f3)

	err = Zipx.Compress(osFiles, "/Users/dombine/Downloads/demo.zip")
	if err != nil {
		t.Error(err)
	}
}
