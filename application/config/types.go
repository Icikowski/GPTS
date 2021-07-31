package config

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"icikowski.pl/gpts/common"
	"icikowski.pl/gpts/utils"
)

// RouteDefinition represents single route details
type RouteDefinition struct {
	ID          string             `json:"id" yaml:"id"`
	Path        string             `json:"path" yaml:"path"`
	Method      string             `json:"method" yaml:"method"`
	Status      *int               `json:"status,omitempty" yaml:"status,omitempty"`
	ContentType *string            `json:"content_type,omitempty" yaml:"content_type,omitempty"`
	Content     *string            `json:"content,omitempty" yaml:"content,omitempty"`
	Headers     *map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// String returns a name of RouteDefinition basing on its Path & Method
func (rd RouteDefinition) String() string {
	return fmt.Sprintf("[%s %s]", rd.Method, rd.Path)
}

type configuration struct {
	routes []RouteDefinition
	mutex  sync.Mutex
}

// GetConfiguration returns a map of paths and their configuration entries
func (c *configuration) GetConfiguration() []RouteDefinition {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	routes := make([]RouteDefinition, len(c.routes))
	copy(routes, c.routes)
	return routes
}

// SetConfiguration cleans current configuration and sets new map of paths and their configuration entries
func (c *configuration) SetConfiguration(routes []RouteDefinition) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	validated, warns := merge(c.routes, routes)
	if warns == nil {
		c.routes = *validated
	}
}

func (c *configuration) ValidateConfiguration(routes []RouteDefinition) (bool, *[]string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, warns := merge(c.routes, routes)
	if warns == nil {
		return true, nil
	}
	return false, warns
}

// SetDefaultConfiguration sets up startup configuration with one prepared route
func (c *configuration) SetDefaultConfiguration() {
	c.mutex.Lock()
	c.routes = []RouteDefinition{
		{
			ID:          uuid.NewString(),
			Path:        "/hello",
			Method:      http.MethodGet,
			Status:      utils.IntToPointer(http.StatusOK),
			ContentType: utils.StringToPointer(common.ContentTypeJSON),
			Content:     utils.StringToPointer(`{"message":"Hello World!"}`),
			Headers: &map[string]string{
				"X-SentBy": "GPTS - General Purpose Test Service",
			},
		},
	}
	c.mutex.Unlock()
	log.Info().Msg("loaded default config as current")
}
