package fn

const (
	// internalPrefix is the prefix given to internal annotations that are used
	// internally by the orchestrator
	//internalPrefix string = "internal.config.kubernetes.io/"

	// IndexAnnotation records the index of a specific resource in a file or input stream.
	//IndexAnnotation string = internalPrefix + "index"

	// PathAnnotation records the path to the file the Resource was read from
	//PathAnnotation string = internalPrefix + "path"

	// SeqIndentAnnotation records the sequence nodes indentation of the input resource
	//SeqIndentAnnotation string = internalPrefix + "seqindent"

	// IdAnnotation records the id of the resource to map inputs to outputs
	//IdAnnotation string = internalPrefix + "id"

	// InternalAnnotationsMigrationResourceIDAnnotation is used to uniquely identify
	// resources during round trip to and from a function execution. We will use it
	// to track the internal annotations and reconcile them if needed.
	//InternalAnnotationsMigrationResourceIDAnnotation = internalPrefix + "annotations-migration-resource-id"

	// ConfigPrefix is the prefix given to the custom kubernetes annotations.
	ConfigPrefix string = "config.kubernetes.io/"

	// KptLocalConfig marks a KRM resource to be skipped from deploying to the cluster via `kpt live apply`.
	KptLocalConfig = ConfigPrefix + "local-config"
)