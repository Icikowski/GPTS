package common

// MIME types
const (
	// ContentTypeJSON represents MIME media type for JSON
	ContentTypeJSON = "application/json"

	// ContentTypeYAML represents MIME media type for YAML
	ContentTypeYAML = "text/yaml"
)

// Messages
const (
	// MsgContentTypeNotAllowed represent a human-readable message for wrong media type error
	MsgContentTypeNotAllowed = "wrong media type (accepting application/json or text/yaml)"
)

// Methods
const (
	// MethodAll represents all HTTP methods at once
	MethodAll = "ALL"
)
