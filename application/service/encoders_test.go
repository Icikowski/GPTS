package service

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"icikowski.pl/gpts/common"
)

var testSubject = map[string]interface{}{
	"hello": "world",
	"test":  true,
}

func TestGetEncoder(t *testing.T) {
	tests := map[string]struct {
		targetEncoding         string
		encoderShouldBePresent bool
		expectedOutput         string
	}{
		"to JSON": {
			targetEncoding:         common.ContentTypeJSON,
			encoderShouldBePresent: true,
			expectedOutput:         `{"hello":"world","test":true}`,
		},
		"to YAML": {
			targetEncoding:         common.ContentTypeYAML,
			encoderShouldBePresent: true,
			expectedOutput:         "hello: world\ntest: true",
		},
		"to unknown type": {
			targetEncoding:         "text/plain",
			encoderShouldBePresent: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			buffer := new(strings.Builder)
			function := getEncoder(tc.targetEncoding, buffer)

			if !tc.encoderShouldBePresent {
				require.Nil(t, function, "encoder should not be found")
				return
			}

			_ = function(testSubject)

			require.Equal(t, tc.expectedOutput, strings.TrimSpace(buffer.String()), "output different than expected")
		})
	}
}

func TestGetDecoder(t *testing.T) {
	tests := map[string]struct {
		sourceEncoding         string
		decoderShouldBePresent bool
		givenSource            string
	}{
		"from JSON": {
			sourceEncoding:         common.ContentTypeJSON,
			decoderShouldBePresent: true,
			givenSource:            `{"hello":"world","test":true}`,
		},
		"from YAML": {
			sourceEncoding:         common.ContentTypeYAML,
			decoderShouldBePresent: true,
			givenSource:            "hello: world\ntest: true",
		},
		"from unknown type": {
			sourceEncoding:         "text/plain",
			decoderShouldBePresent: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			buffer := bytes.NewBufferString(tc.givenSource)
			function := getDecoder(tc.sourceEncoding, buffer)

			if !tc.decoderShouldBePresent {
				require.Nil(t, function, "decoder should not be found")
				return
			}

			var actualOutput map[string]interface{}
			function(&actualOutput)

			require.Equal(t, testSubject, actualOutput, "output different than expected")
		})
	}
}
