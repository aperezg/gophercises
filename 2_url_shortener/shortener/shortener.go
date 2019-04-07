package shortener

// Shortener especification of all shortener types
type Shortener interface {
	Parse() (ShortMap, error)
}

// Link representation of redirect from path
type Link struct {
	Path string `yml:"path"`
	URL  string `yml:"url"`
}

// ShortMap is a collection of shorteners
type ShortMap []Link

// ToMap transform ShortMap into a map[string]string
func (s ShortMap) ToMap() map[string]string {
	m := make(map[string]string)
	for _, short := range s {
		m[short.Path] = short.URL
	}
	return m
}
