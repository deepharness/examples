// Code generated by goa v3.11.2, DO NOT EDIT.
//
// HTTP request path constructors for the updown service.
//
// Command:
// $ goa gen goa.design/examples/upload_download/design -o upload_download

package client

import (
	"fmt"
)

// UploadUpdownPath returns the URL path to the updown service upload HTTP endpoint.
func UploadUpdownPath(dir string) string {
	return fmt.Sprintf("/upload/%v", dir)
}

// DownloadUpdownPath returns the URL path to the updown service download HTTP endpoint.
func DownloadUpdownPath(filename string) string {
	return fmt.Sprintf("/download/%v", filename)
}
