package shortyaml

import (
	"github.com/aperezg/gophercises/2_url_shortener/shortener"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// ShortYAML representation of shortener from YAML
type ShortYAML struct {
	data []byte
}

// New initialize of ShortYAML
func New(data []byte) shortener.Shortener {
	return &ShortYAML{data}
}

// Parse parse YAML data into a ShortMap
func (y *ShortYAML) Parse() (shortener.ShortMap, error) {
	var s shortener.ShortMap
	err := yaml.Unmarshal(y.data, &s)
	if err != nil {
		return nil, errors.Wrap(err, "error trying to unmashal yaml")
	}

	return s, nil
}
