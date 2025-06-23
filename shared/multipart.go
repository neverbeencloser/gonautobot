package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	reflect "reflect"
	"strings"
)

// NewMultipartBody creates a new multipart body for the given input struct.
//
// This should only be used on API endpoints that support
// multipart form data and image uploads, such as /dcim/device-types/.
//
// The input struct should have a `form` tag on the fields specifying
// the form field name. The `,omitempty` tag option can be used to skip
// empty fields. The `,upload` tag option can be used to indicate that
// the field is a file upload.
func NewMultipartBody(inputStruct any) (*bytes.Buffer, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	v := reflect.ValueOf(inputStruct)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("form")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		formName := parts[0]
		var omitempty, upload bool
		for _, part := range parts[1:] {
			if part == "omitempty" {
				omitempty = true
			} else if part == "upload" {
				upload = true
			}
		}

		if omitempty && value.IsZero() {
			continue
		}

		if upload {
			if value.Kind() != reflect.String {
				return nil, "", fmt.Errorf("field %s is marked as upload but is not a string", formName)
			}
			filePath := value.String()
			if filePath == "" {
				// For empty file paths, set the field to an empty string and continue.
				// For existing resources on Update(), this will clear a file image.
				if err := w.WriteField(formName, fmt.Sprintf("%v", value.Interface())); err != nil {
					return nil, "", fmt.Errorf("failed to write field %s: %w", formName, err)
				}
				continue
			}

			file, err := os.Open(filePath) // nolint: gosec
			if err != nil {
				return nil, "", fmt.Errorf("failed to open file %s: %w", filePath, err)
			}
			defer file.Close()

			part, err := w.CreateFormFile(formName, filepath.Base(filePath))
			if err != nil {
				return nil, "", fmt.Errorf("failed to create form file for %s: %w", formName, err)
			}
			if _, err = io.Copy(part, file); err != nil {
				return nil, "", fmt.Errorf("failed to copy file content for %s: %w", formName, err)
			}
		} else {
			switch value.Kind() { // nolint: exhaustive
			case reflect.Slice:
				for j := 0; j < value.Len(); j++ {
					if err := w.WriteField(formName, fmt.Sprintf("%v", value.Index(j).Interface())); err != nil {
						return nil, "", fmt.Errorf("failed to write slice field %s: %w", formName, err)
					}
				}
			case reflect.Map:
				jsonBytes, err := json.Marshal(value.Interface())
				if err != nil {
					return nil, "", fmt.Errorf("failed to marshal map field %s: %w", formName, err)
				}
				if err := w.WriteField(formName, string(jsonBytes)); err != nil {
					return nil, "", fmt.Errorf("failed to write map field %s: %w", formName, err)
				}
			default:
				if err := w.WriteField(formName, fmt.Sprintf("%v", value.Interface())); err != nil {
					return nil, "", fmt.Errorf("failed to write field %s: %w", formName, err)
				}
			}
		}
	}

	if err := w.Close(); err != nil {
		return nil, "", fmt.Errorf("failed to close multipart writer: %w", err)
	}

	return &b, w.FormDataContentType(), nil
}
