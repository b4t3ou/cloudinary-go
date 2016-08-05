package cloudinary

import (
	"testing"
)

func TestCloudinary_UploadInvalidOption(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic!")
		}
	}()

	c := Create("", "", "")

	c.Upload("", map[string]string{"foo": "bar"})
}
