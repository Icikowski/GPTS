package common

// MIME types
const (
	// ContentTypeJSON represents MIME media type for JSON
	ContentTypeJSON string = "application/json"

	// ContentTypeYAML represents MIME media type for YAML
	ContentTypeYAML string = "text/yaml"
)

// Headers
const (
	// HeaderContentType represents "Content-Type" header key
	HeaderContentType string = "Content-Type"
)

// Messages
const (
	// MsgContentTypeNotAllowed represent a human-readable message for wrong media type error
	MsgContentTypeNotAllowed string = "wrong media type (accepting application/json or text/yaml)"
)

// Build variables default value
const (
	// BuildValueUnknown represents an "unknown" word
	BuildValueUnknown string = "unknown"
)

// Component names
const (
	// ComponentConfig represents a name of configuration manager component
	ComponentConfig string = "configuration"

	// ComponentHealth represents a name of health component
	ComponentHealth string = "health"

	// ComponentService represents a name of service component
	ComponentService string = "service"

	// ComponentCLI represents a name of CLI component
	ComponentCLI string = "cli"
)
