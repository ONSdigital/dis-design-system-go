package model

// Metadata represents the metadata fields for a page.
type Metadata struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ServiceName string   `json:"serviceName"`
	Keywords    []string `json:"keywords"`
}
