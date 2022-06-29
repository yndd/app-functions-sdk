package fn

import (
	"fmt"
	"sort"

	"github.com/yndd/app-functions-sdk/go/fn/internal"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	ResourceContextKind       = "ResourceContext"
	ResourceContextAPIVersion = "app.yndd.io/v1"
)

type ResourceContext struct {
	Input   *ResourceContextInputs `yaml:"input" json:"input"`                         // the input CR(s)
	Outputs KubeObjects            `yaml:"outputs,omitempty" json:"outputs,omitempty"` // the rendered output CR
	Results Results                `yaml:"results,omitempty" json:"results,omitempty"` // result context
}

type ResourceContextInputs struct {
	Origin *KubeObject `yaml:"origin" json:"origin"`                     // the origin CR in the intent/app
	Target *KubeObject `yaml:"target,omitempty" json:"target,omitempty"` // could be node or target
	Items  KubeObjects `yaml:"items,omitempty" json:"items,omitempty"`   // additional input items like OC
}

// ParseResourceContext parses a ResourceContext from the input byte array. This function can be used to parse either KRM fn input
// or KRM fn output
func ParseResourceContext(in []byte) (*ResourceContext, error) {
	rctx := &ResourceContext{}
	rctxObj, err := ParseKubeObject(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input bytes: %w", err)
	}
	if rctxObj.GetKind() != ResourceContextKind {
		return nil, fmt.Errorf("input was of unexpected kind %q; expected %s", rctxObj.GetKind(), ResourceContextKind)
	}
	if rctxObj.GetAPIVersion() != ResourceContextAPIVersion {
		return nil, fmt.Errorf("input was of unexpected apiversion %q; expected %s", rctxObj.GetAPIVersion(), ResourceContextAPIVersion)
	}
	// Parse input. Input cannot be empty, e.g. an input ResourceContext always need to origin CR.
	input, found, err := rctxObj.obj.GetNestedMap("input")
	if err != nil {
		return nil, fmt.Errorf("failed when tried to get input: %w", err)
	}
	if !found {
		return nil, fmt.Errorf("input was of expected but not found in %s", ResourceContextKind)
	}
	rctx.Input = &ResourceContextInputs{}
	// Parse origin, Origin cannot be empty, e.g. an input ResourceContext always need to origin CR.
	origin, found, err := input.GetNestedMap("origin")
	if err != nil {
		return nil, fmt.Errorf("failed when tried to get origin: %w", err)
	}
	if !found {
		return nil, fmt.Errorf("origin was of expected but not found in %s", ResourceContextKind)
	}
	rctx.Input.Origin = asKubeObject(origin)
	// Parse target, Target can be empty
	target, found, err := input.GetNestedMap("target")
	if err != nil {
		return nil, fmt.Errorf("failed when tried to get target: %w", err)
	}
	if found {
		rctx.Input.Target = asKubeObject(target)
	}
	// Parse items, Items can be empty, serve as additional input context
	items, found, err := input.GetNestedSlice("items")
	if err != nil {
		return nil, fmt.Errorf("failed when tried to get items: %w", err)
	}
	if found {
		objectItems, err := items.Elements()
		if err != nil {
			return nil, fmt.Errorf("failed to extract objects from items: %w", err)
		}
		for i := range objectItems {
			rctx.Input.Items = append(rctx.Input.Items, asKubeObject(objectItems[i]))
		}
	}
	// Parse outputs, Outputs can be empty, will be added when processed
	outputs, found, err := rctxObj.obj.GetNestedSlice("outputs")
	if err != nil {
		return nil, fmt.Errorf("failed when tried to get outputs: %w", err)
	}
	if found {
		objectOutputs, err := outputs.Elements()
		if err != nil {
			return nil, fmt.Errorf("failed to extract objects from ouputs: %w", err)
		}
		for i := range objectOutputs {
			rctx.Outputs = append(rctx.Input.Items, asKubeObject(objectOutputs[i]))
		}
	}
	// Parse Results. Results can be empty.
	res, found, err := rctxObj.obj.GetNestedSlice("results")
	if err != nil {
		return nil, fmt.Errorf("failed when tried to get results: %w", err)
	}
	if found {
		var results Results
		err = res.Node().Decode(&results)
		if err != nil {
			return nil, fmt.Errorf("failed to decode results: %w", err)
		}
		rctx.Results = results
	}

	return rctx, nil
}

// toYNode converts the ResourceContext to the yaml.Node representation.
func (rctx *ResourceContext) toYNode() (*yaml.Node, error) {
	reMap := internal.NewMap(nil)
	reObj := &KubeObject{SubObject{reMap}}
	reObj.SetAPIVersion(ResourceContextAPIVersion)
	reObj.SetKind(ResourceContextKind)

	if rctx.Input != nil {
		if rctx.Input.Origin != nil {
			if err := reMap.SetNestedMap(rctx.Input.Origin.node(), "origin"); err != nil {
				return nil, err
			}
		}
		if rctx.Input.Target != nil {
			if err := reMap.SetNestedMap(rctx.Input.Origin.node(), "target"); err != nil {
				return nil, err
			}
		}
		if rctx.Input.Items != nil && len(rctx.Input.Items) > 0 {
			itemsSlice := internal.NewSliceVariant()
			for i := range rctx.Input.Items {
				itemsSlice.Add(rctx.Input.Items[i].node())
			}
			if err := reMap.SetNestedSlice(itemsSlice, "items"); err != nil {
				return nil, err
			}
		}
	}

	if rctx.Outputs != nil && len(rctx.Outputs) > 0 {
		outputsSlice := internal.NewSliceVariant()
		for i := range rctx.Outputs {
			outputsSliceSlice.Add(rctx.Outputs[i].node())
		}
		if err := reMap.SetNestedSlice(outputsSliceSlice, "outputs"); err != nil {
			return nil, err
		}
	}

	if rctx.Results != nil && len(rctx.Results) > 0 {
		resultsSlice := internal.NewSliceVariant()
		for _, result := range rctx.Results {
			mv, err := internal.TypedObjectToMapVariant(result)
			if err != nil {
				return nil, err
			}
			resultsSlice.Add(mv)
		}
		if err := reMap.SetNestedSlice(resultsSlice, "results"); err != nil {
			return nil, err
		}
	}

	return reMap.Node(), nil
}

// ToYAML converts the ResourceList to yaml.
func (rctx *ResourceContext) ToYAML() ([]byte, error) {
	// Sort the resources input.Items and outputs first.
	rctx.Sort()
	ynode, err := rctx.toYNode()
	if err != nil {
		return nil, err
	}
	doc := internal.NewDoc([]*yaml.Node{ynode}...)
	return doc.ToYAML()
}

// Sort sorts the ResourceContext.input by apiVersion, kind, namespace and name.
// Sort sorts the ResourceContext.output by apiVersion, kind, namespace and name.
func (rctx *ResourceContext) Sort() {
	sort.Sort(rctx.Input.Items)
	sort.Sort(rctx.Outputs)
}

func (rctx *ResourceContext) AddOuput(output *KubeObject) error {
	if rctx.Outputs == nil {
		rctx.Outputs = KubeObjects{}
	}
	rctx.Outputs = append(rctx.Outputs, output)
	return nil
}
