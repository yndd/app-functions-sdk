package fn

import "sigs.k8s.io/kustomize/kyaml/yaml"

// Severity indicates the severity of the Result
type Severity string

const (
	// Error indicates the result is an error.  Will cause the function to exit non-0.
	Error Severity = "error"
	// Warning indicates the result is a warning
	Warning Severity = "warning"
	// Info indicates the result is an informative message
	Info Severity = "info"
)

// Result defines a validation result
type Result struct {
	// Message is a human readable message. This field is required.
	Message string `yaml:"message,omitempty" json:"message,omitempty"`

	// Severity is the severity of this result
	Severity Severity `yaml:"severity,omitempty" json:"severity,omitempty"`

	// ResourceRef is a reference to a resource.
	// Required fields: apiVersion, kind, name.
	ResourceRef *yaml.ResourceIdentifier `yaml:"resourceRef,omitempty" json:"resourceRef,omitempty"`

	// Tags is an unstructured key value map stored with a result that may be set
	// by external tools to store and retrieve arbitrary metadata
	Tags map[string]string `yaml:"tags,omitempty" json:"tags,omitempty"`
}

type Results []*Result
