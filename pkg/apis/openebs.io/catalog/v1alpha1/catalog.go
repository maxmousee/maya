/*
Copyright 2019 The OpenEBS Authors

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
)

func init() {
	// Register adds catalog & cataloglist objects to
	// SchemeBuilder so they can be added to a Scheme
	SchemeBuilder.Register(&Catalog{}, &CatalogList{})
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=catalog

// Catalog contains the desired specification of one or
// more resources
type Catalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CatalogSpec   `json:"spec"`
	Status CatalogStatus `json:"status"`
}

// CatalogSpec provides the specifications of a Catalog
type CatalogSpec struct {
	Items []CatalogResource `json:"items"`
}

type CatalogResource struct {
	//Group     string `json:"group"`
	//Kind      string `json:"kind"`
	//Version   string `json:"version"`
	Namespace string `json:"namespace"`
	Template  string `json:"template"`
}

// CatalogStatus represents the current state of Catalog
type CatalogStatus struct {
	Phase string `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=catalogs

// CatalogList is a list of Catalogs
type CatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Catalog `json:"items"`
}
