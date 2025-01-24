package service

import (
	"encoding/json"
	"io"

	"git.sr.ht/~icikowski/gpts/common"
	"gopkg.in/yaml.v2"
)

func getEncoder(contentType string, writer io.Writer) func(v any) error {
	switch contentType {
	case common.ContentTypeJSON:
		return json.NewEncoder(writer).Encode
	case common.ContentTypeYAML:
		return yaml.NewEncoder(writer).Encode
	}
	return nil
}

func getDecoder(contentType string, reader io.Reader) func(v any) error {
	switch contentType {
	case common.ContentTypeJSON:
		return json.NewDecoder(reader).Decode
	case common.ContentTypeYAML:
		return yaml.NewDecoder(reader).Decode
	}
	return nil
}
