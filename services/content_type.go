package services

import (
	"errors"
	"strings"
)

// parseParams takes a raw string of parameters passed to a content
// type, and returns them in map[string]string form.
func parseParams(rawParams []string) (map[string]string, error) {
	params := make(map[string]string)
	for _, param := range rawParams {
		nameAndValue := strings.SplitN(param, "=", 2)
		if len(nameAndValue) != 2 {
			return nil, errors.New("Received parameter " + param + " with no equal sign")
		}
		name := strings.TrimSpace(nameAndValue[0])
		value := strings.TrimSpace(nameAndValue[1])

		params[name] = value
	}
	return params, nil
}

// ContentType represents a single content type, complete with
// parameters, such as that passed in an HTTP Accept or Content-Type
// header.
type ContentType struct {
	MimeType   string
	Parameters map[string]string
}

// ParseContentType takes a content-type string and parses it into a
// mimetype and parameters, returning the ContentType representing the
// string.
func ParseContentType(rawType string) (*ContentType, error) {
	rawType = strings.TrimSpace(strings.ToLower(rawType))
	if len(rawType) == 0 {
		return nil, nil
	}
	contentType := new(ContentType)
	mimeAndParams := strings.Split(rawType, ";")
	contentType.MimeType = strings.TrimSpace(mimeAndParams[0])
	params, err := parseParams(mimeAndParams[1:])
	if err != nil {
		return nil, err
	}
	contentType.Parameters = params
	return contentType, nil
}
