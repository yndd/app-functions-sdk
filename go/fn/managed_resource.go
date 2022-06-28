package fn

import (
	"fmt"
)

// ParseManagedResource parses a Managed Resource from the input byte array. This function can be used to parse either KRM fn input
// or KRM fn output
func ParseManagedResource(in []byte) (*KubeObject, error) {
	//func ParseManagedResource(in []byte) (resource.Managed, error) {
	//rl := &ResourceList{}
	mrObj, err := ParseKubeObject(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input bytes: %w", err)
	}
	/*
		if mrObj.GetKind() != kio.ResourceListKind {
			return nil, fmt.Errorf("input was of unexpected kind %q; expected ResourceList", mrObj.GetKind())
		}
			// Parse FunctionConfig. FunctionConfig can be empty, e.g. `kubeval` fn does not require a FunctionConfig.
			fc, found, err := mrObj.obj.GetNestedMap("functionConfig")
			if err != nil {
				return nil, fmt.Errorf("failed when tried to get functionConfig: %w", err)
			}
			if found {
				rl.FunctionConfig = asKubeObject(fc)
			} else {
				rl.FunctionConfig = NewEmptyKubeObject()
			}

			// Parse Items. Items can be empty, e.g. an input ResourceList for a generator function may not have items.
			items, found, err := rlObj.obj.GetNestedSlice("items")
			if err != nil {
				return nil, fmt.Errorf("failed when tried to get items: %w", err)
			}
			if found {
				objectItems, err := items.Elements()
				if err != nil {
					return nil, fmt.Errorf("failed to extract objects from items: %w", err)
				}
				for i := range objectItems {
					rl.Items = append(rl.Items, asKubeObject(objectItems[i]))
				}
			}

			// Parse Results. Results can be empty.
			res, found, err := rlObj.obj.GetNestedSlice("results")
			if err != nil {
				return nil, fmt.Errorf("failed when tried to get results: %w", err)
			}
			if found {
				var results Results
				err = res.Node().Decode(&results)
				if err != nil {
					return nil, fmt.Errorf("failed to decode results: %w", err)
				}
				rl.Results = results
			}
	*/
	return mrObj, nil
}
