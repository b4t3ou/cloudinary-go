package cloudinary

import (
	"testing"
)

func TestCloudinary_Upload(t *testing.T) {
	c := Create("937244359587683", "IX-OgfXv3c7zb5lhOsKwBeBi9cA", "dbarfmrrj")

	options := map[string]string{
		"public_id": "test_image",
	}

	resp, err := c.Upload("https://static.pexels.com/photos/36487/above-adventure-aerial-air-large.jpg", options)

	if err != nil {
		t.Error("Something went wrong with the upload")
	}

	if resp.PublicId != "test_image" {
		t.Error("Something went wrong with the upload")
	}
}

func TestCloudinary_FalseUpload(t *testing.T) {
	c := Create("937244359587683", "IX-OgfXv3c7zb5lhOsKwBeBi9cA", "dbarfmrrj")

	options := map[string]string{
		"public_id": "test_image",
	}

	resp, _ := c.Upload("https://static.pexels.com/photos/36487/above-adventure-aerial-air-large.jpgasdas", options)

	if resp.Error.Message == "" {
		t.Error("Something went wrong the upload process has to fail")
	}
}
