package model

// TaxonomyNode represents a node in a taxonomy structure.
type TaxonomyNode struct {
	Title    string         `json:"title"`
	URI      string         `json:"uri"`
	Type     string         `json:"type,omitempty"`
	Children []TaxonomyNode `json:"children,omitempty"`
}
