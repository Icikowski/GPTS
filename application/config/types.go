package config

import (
	"net/http"
	"sync"

	"github.com/rs/zerolog"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/utils"
)

// Response represents response details
type Response struct {
	Status      *int               `json:"status,omitempty" yaml:"status,omitempty"`
	ContentType *string            `json:"content_type,omitempty" yaml:"content_type,omitempty"`
	Content     *string            `json:"content,omitempty" yaml:"content,omitempty"`
	Headers     *map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// Route represents single route details
type Route struct {
	AllowSubpaths bool      `json:"allow_subpaths" yaml:"allow_subpaths"`
	GET           *Response `json:"get,omitempty" yaml:"get,omitempty"`
	POST          *Response `json:"post,omitempty" yaml:"post,omitempty"`
	PUT           *Response `json:"put,omitempty" yaml:"put,omitempty"`
	PATCH         *Response `json:"patch,omitempty" yaml:"patch,omitempty"`
	DELETE        *Response `json:"delete,omitempty" yaml:"delete,omitempty"`
	Default       *Response `json:"default,omitempty" yaml:"default,omitempty"`
}

type configuration struct {
	routes map[string]Route
	mutex  sync.Mutex
}

// GetConfiguration returns a map of paths and their configuration entries
func (c *configuration) GetConfiguration() map[string]Route {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	routes := map[string]Route{}
	for path, route := range c.routes {
		path, route := path, route
		routes[path] = route
	}
	return routes
}

// SetConfiguration cleans current configuration and sets new map of paths and their configuration entries
func (c *configuration) SetConfiguration(routes map[string]Route) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.routes = map[string]Route{}
	for path, route := range routes {
		path, route := path, route
		c.routes[path] = route
	}
}

// SetDefaultConfiguration sets up startup configuration with one prepared route
func (c *configuration) SetDefaultConfiguration() {
	c.mutex.Lock()
	c.routes = map[string]Route{
		"/hello": {
			AllowSubpaths: true,
			Default: &Response{
				Status:      utils.IntToPointer(http.StatusOK),
				ContentType: utils.StringToPointer(common.ContentTypeJSON),
				Content:     utils.StringToPointer(`{"message":"Hello World!"}`),
				Headers: &map[string]string{
					"X-SentBy": "GPTS - General Purpose Test Service",
				},
			},
		},
	}
	c.mutex.Unlock()
	log.Info().Msg("loaded default config as current")
}

var _ zerolog.LogObjectMarshaler = Response{}

// MarshalZerologObject implements zerolog.LogObjectMarshaler interface
func (r Response) MarshalZerologObject(e *zerolog.Event) {
	e.Interface("status", map[bool]interface{}{
		true:  *r.Status,
		false: "nil",
	}[r.Status != nil])
	e.Interface("contentType", map[bool]interface{}{
		true:  *r.ContentType,
		false: "nil",
	}[r.ContentType != nil])
	e.Interface("headers", map[bool]interface{}{
		true:  *r.Headers,
		false: "nil",
	}[r.Headers != nil])
}

var _ zerolog.LogObjectMarshaler = Route{}

// MarshalZerologObject implements zerolog.LogObjectMarshaler interface
func (r Route) MarshalZerologObject(e *zerolog.Event) {
	e.Bool("allowSubpaths", r.AllowSubpaths)
	e.Dict("methods", zerolog.Dict().
		Bool("default", r.Default != nil).
		Bool("get", r.GET != nil).
		Bool("post", r.POST != nil).
		Bool("put", r.PUT != nil).
		Bool("patch", r.PATCH != nil).
		Bool("delete", r.DELETE != nil),
	)
}
