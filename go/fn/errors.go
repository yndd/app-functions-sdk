package fn

import (
	"fmt"
	"strings"
)

// ErrMissingFnConfig raises error if a required functionConfig is missing.
type ErrMissingFnConfig struct{}

func (ErrMissingFnConfig) Error() string {
	return "unable to find the functionConfig in the resourceList"
}

// errKubeObjectFields raises if the KubeObject operation panics.
type errKubeObjectFields struct {
	obj    *KubeObject
	fields []string
}

func (e *errKubeObjectFields) Error() string {
	return fmt.Sprintf("Resource(apiVersion=%v, kind=%v, Name=%v) has unmatched field type: `%v",
		e.obj.GetAPIVersion(), e.obj.GetKind(), e.obj.GetName(), strings.Join(e.fields, "/"))
}

// errSubObjectFields raises if the SubObject operation panics.
type errSubObjectFields struct {
	fields []string
}

func (e *errSubObjectFields) Error() string {
	return fmt.Sprintf("SubObject has unmatched field type: `%v", strings.Join(e.fields, "/"))
}

type errResultEnd struct {
	obj     *KubeObject
	message string
}

func (e *errResultEnd) Error() string {
	if e.obj != nil {
		return fmt.Sprintf("function is terminated by %v: %v", e.obj.ShortString(), e.message)
	}
	return fmt.Sprintf("function is terminated: %v", e.message)
}