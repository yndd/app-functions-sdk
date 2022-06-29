package fn

import (
	"fmt"
	"strings"

	"sigs.k8s.io/kustomize/kyaml/yaml"
)

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

func (i Result) Error() string {
	return (i).String()
}

// String provides a human-readable message for the result item
func (i Result) String() string {
	identifier := i.ResourceRef
	var idStringList []string
	if identifier != nil {
		if identifier.APIVersion != "" {
			idStringList = append(idStringList, identifier.APIVersion)
		}
		if identifier.Kind != "" {
			idStringList = append(idStringList, identifier.Kind)
		}
		if identifier.Namespace != "" {
			idStringList = append(idStringList, identifier.Namespace)
		}
		if identifier.Name != "" {
			idStringList = append(idStringList, identifier.Name)
		}
	}
	formatString := "[%s]"
	severity := i.Severity
	// We default Severity to Info when converting a result to a message.
	if i.Severity == "" {
		severity = Info
	}
	list := []interface{}{severity}
	if len(idStringList) > 0 {
		formatString += " %s"
		list = append(list, strings.Join(idStringList, "/"))
	}
	formatString += ": %s"
	list = append(list, i.Message)
	return fmt.Sprintf(formatString, list...)
}

type Results []*Result

// Error enables Results to be returned as an error
func (e Results) Error() string {
	var msgs []string
	for _, i := range e {
		msgs = append(msgs, i.String())
	}
	return strings.Join(msgs, "\n\n")
}

func ErrorResult(err error) *Result {
	return GeneralResult(err.Error(), Error)
}

func GeneralResult(msg string, severity Severity) *Result {
	return &Result{
		Message:  msg,
		Severity: severity,
	}
}
