package fn

// ResourceContextProcessor interface
type ResourceContextProcessor interface {
	Process(rc *ResourceContext) (bool, error)
}

// ResourceContextProcessorFunc converts a compatible function to a ResourceContextProcessor.
// ResourceContextProcessorFunc implements a ResourceContextProcessor interface
type ResourceContextProcessorFunc func(rc *ResourceContext) (bool, error)

func (p ResourceContextProcessorFunc) Process(rc *ResourceContext) (bool, error) {
	return p(rc)
}
