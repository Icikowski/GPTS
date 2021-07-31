package service

type defaultResponse struct {
	Host    string              `json:"host" yaml:"host"`
	Path    string              `json:"path" yaml:"path"`
	Method  string              `json:"method" yaml:"method"`
	Headers map[string][]string `json:"headers" yaml:"headers"`
}
