package cloudinary

import (
	"encoding/json"
	"net/url"
)

// Available upload fields
// Usage http://cloudinary.com/documentation/image_upload_api_reference#api_example_1
var uploadOptions []string = []string{
	"public_id",
	"use_filename",
	"unique_filename",
	"folder",
	"overwrite",
	"resource_type",
	"type",
	"tags",
	"context",
	"transformation",
	"format",
	"allowed_formats",
	"eager",
	"async",
	"eager_async",
	"proxy",
	"headers",
	"callback",
	"notification_url",
	"eager_notification_url",
	"faces",
	"image_metadata",
	"exif",
	"colors",
	"phash",
	"face_coordinates",
	"custom_coordinates",
	"backup",
	"return_delete_token",
	"invalidate",
	"discard_original_filename",
	"moderation",
	"upload_preset",
	"raw_convert",
	"categorization",
	"auto_tagging",
	"background_removal",
	"detection",
	"timestamp",
}

// Upload image success response struct
type Upload struct {
	PublicId         string    `json:"public_id"`
	Version          int       `json:"version"`
	Signature        string    `json:"signature"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	Format           string    `json:"format"`
	ResourceType     string    `json:"resource_type"`
	CreatedAt        string    `json:"created_at"`
	Tags             []string  `json:"tags,omitempty"`
	Bytes            int       `json:"bytes"`
	Type             string    `json:"type"`
	Etag             string    `json:"etag"`
	Url              string    `json:"url"`
	SecureUrl        string    `json:"secure_url"`
	OriginalFilename string    `json:"original_filename"`
	Error            ErrorResp `json:"error,omitempty"`
}

// Upload is uploading an image
func (c *Cloudinary) Upload(file string, options Option) (*Upload, error) {
	c.checkOptionsAreValid(options, uploadOptions)
	options = c.sortParamsByKey(options)

	form := url.Values{}
	for paramName, value := range options {
		form.Add(paramName, value)
	}
	form.Add("file", file)

	body := c.send("upload", form, options)

	upload := &Upload{}
	err := json.Unmarshal(body, upload)

	if err != nil {
		return nil, err
	}

	return upload, nil
}
