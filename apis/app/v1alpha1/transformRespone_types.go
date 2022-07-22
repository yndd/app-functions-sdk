/*
Copyright 2021 NDD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//targetv1alpha1pb "github.com/yndd/topology/gen/go/apis/topo/v1alpha1"
)

// TransformRequestSpec struct
type TransformResponseSpec struct {
	Output runtime.RawExtension        `json:"output,omitempty"`
	Result TransformResponseSpecResult `json:"result,omitempty"`
}

type TransformResponseSpecResult struct {
	Message  string `json:"message,omitempty"`
	Severity string `json:"severity,omitempty"`
}

// +kubebuilder:object:root=true

// TransformRequest is the Schema for the TransformRequest API
// +kubebuilder:resource:categories={yndd,app}
type TransformResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TransformResponseSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// TransformRequestList contains a list of TransformRequests
type TransformResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransformRequest `json:"items"`
}
