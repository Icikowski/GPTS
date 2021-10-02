package common

// MIME types
const (
	// ContentTypeJSON represents MIME media type for JSON
	ContentTypeJSON = "application/json"

	// ContentTypeYAML represents MIME media type for YAML
	ContentTypeYAML = "text/yaml"
)

// Headers
const (
	// HeaderContentType represents "Content-Type" header key
	HeaderContentType = "Content-Type"
)

// Messages
const (
	// MsgContentTypeNotAllowed represent a human-readable message for wrong media type error
	MsgContentTypeNotAllowed = "wrong media type (accepting application/json or text/yaml)"
)

// Build variables default value
const (
	// BuildValueUnknown represents an "unknown" word
	BuildValueUnknown = "unknown"
)
