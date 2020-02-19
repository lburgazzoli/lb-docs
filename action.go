/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"net/url"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Action describe a processing capability available on the platform
// and captures both the logical (meta data, public api) the physical
// details (type of implementation).
type Action struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionSpec   `json:"spec,omitempty"`
	Status ActionStatus `json:"status,omitempty"`
}

// ********************************************************************************
//
// Spec
//
// ********************************************************************************

// ActionSpec describe the Action.
type ActionSpec struct {
	// DataShape defines the consumed/produced data formats for
	// the action
	DataShape struct {
		In       *Schema           `json:"int,omitempty"`
		Out      *Schema           `json:"our,omitempty"`
		Metadata map[string]string `json:"metadata,omitempty"`
	}

	// Parameters defines parameters required by the action to run.
	// They can be authentication information
	Parameters map[string]Parameter

	ImplementationSpec `json:",inline"`
	Meta               *ImplementationSpec `json:"meta,omitempty"`
}

// ImplementationSpec describes the implementation details of
// the Action. Note that
type ImplementationSpec struct {
	// Endpoint defines an endpoint that need to be invoked to
	// execute the action.
	Endpoint *struct {
		// The URL fo the service to be invoked
		URL url.URL `json:"url,omitempty"`

		//TODO: add fields to configure how to bind parameters
		//      to headers
	}

	// Container defines a container that implement the action,
	// how to run it application specific as example it can be
	// used to create a pod and expose it as as service ot it
	// can run a sidecar.
	Container *struct {
		// The application container that you want to run,
		corev1.Container `json:"container"`

		//TODO: add fields to configure how to bind parameters
		//      to headers, volumes, etc.
	}

	// Dependency defines a set of dependencies that have to be
	// taken into account by the runtime to execute the action.
	Dependency *struct {
		Dependencies []Dependency `json:"dependencies,omitempty"`

		//TODO: add fields to configure how to bind parameters
		//      to headers, volumes, etc.
	}
}

// Dependency ---
type Dependency struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Parameter ---
type Parameter struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Required    string `json:"required,omitempty"`

	// The type of the parameter, can be a simple type such as
	// string, boolean, integer or a complex one defined by a
	// schema.
	Type struct {
		Name   string
		Schema *Schema `json:"schema,omitempty"`
	}
}

// Schema describe the shape of data consumed or produced by the
// Action
type Schema struct {
	// Type specify the type of the schema stored in one of the
	// options below
	Type string `json:"type,omitempty"`

	// RegistryRef references a schema defined in a schema registry
	// like Apicurio Schema Registry (https://github.com/Apicurio/apicurio-registry)
	RegistryRef string `json:"registryKeyRef,omitempty"`

	// ConfigMapKeyRef references a key from a ConfigMap where the scheme
	// is defined. The type of the content is defined by the field Type
	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
}

// ********************************************************************************
//
// Status
//
// ********************************************************************************

// ActionStatus represent the status of an Action at a given time.
type ActionStatus struct {
}
