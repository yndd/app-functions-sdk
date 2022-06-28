package fn

import "github.com/yndd/ndd-runtime/pkg/resource"

type ManagedResourceProcessor interface {
	Process(mr resource.Managed) (bool, error)
}

// ManagedResourceProcessorFunc converts a compatible function to a ManagedResourceProcessor.
type ManagedResourceProcessorFunc func(mr resource.Managed) (bool, error)

func (p ManagedResourceProcessorFunc) Process(mr resource.Managed) (bool, error) {
	return p(mr)
}