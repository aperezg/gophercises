package urlshort

import (
	"net/http"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusPermanentRedirect)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var s []shortener

	err := yaml.Unmarshal(yml, &s)
	if err != nil {
		return nil, errors.Wrap(err, "error trying to unmashal yaml")
	}
	pathMap := buildMap(s)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(s []shortener) map[string]string {
	m := make(map[string]string)
	for _, short := range s {
		m[short.Path] = short.URL
	}
	return m
}

type shortener struct {
	Path string `yml:"path"`
	URL  string `yml:"url"`
}
