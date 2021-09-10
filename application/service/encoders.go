package service

import (
	"encoding/json"
	"io"

	"gopkg.in/yaml.v2"
	"icikowski.pl/gpts/common"
)

func getEncoder(contentType string, writer io.Writer) func(v interface{}) error {
	switch contentType {
	case common.ContentTypeJSON:
		return json.NewEncoder(writer).Encode
	case common.ContentTypeYAML:
		return yaml.NewEncoder(writer).Encode
	}
	return nil
}

func getDecoder(contentType string, reader io.Reader) func(v interface{}) error {
	switch contentType {
	case common.ContentTypeJSON:
		return json.NewDecoder(reader).Decode
	case common.ContentTypeYAML:
		return yaml.NewDecoder(reader).Decode
	}
	return nil
}
