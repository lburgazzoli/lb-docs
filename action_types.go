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
	//Defines the consumed/produced messages
	Produces *Message `json:"produces,omitempty"`
	Consumes *Message `json:"consumes,omitempty"`

	// Parameters defines parameters required by the action to be
	// executed.
	Parameters map[string]Parameter

	// Binding define the implementation detail of
	// the Action
	Binding *BindingSpec `json:"binding,omitempty"`

	// Meta define the implementation detail of
	// the Action metadata service
	Meta *BindingSpec `json:"meta,omitempty"`
}

// BindingSpec describes the implementation details of
// the Action. Note that
type BindingSpec struct {
	// HTTP defines an http endpoint that need to be invoked to
	// execute the action.
	HTTP *struct {
		// The URL fo the service to be invoked which support mustache
		// template engine for easy binding.
		//
		// The special scheme `container` can be used to to reference
		// a container defined by the `Container` struct (useful when
		// the container should be executed as sidecar)
		//
		// ImplementationSpec {
		//     Endpoint: Endpoint {
		//	       URL: 'container://my-container/search?q={{parameter-name}}'
		//     }
		//     Container: Container {
		//         Name: my-container
		//	       Image: 'my-imag'
		//     }
		// }
		//
		URL string `json:"url,omitempty"`

		// Defines the binding of Action's parameters to the
		// target endpoint.
		Parameters []struct {
			ParameterBinding `json:",inline"`

			// The target, like query, header
			Target string `json:"target,omitempty"`
		}

		// Container defines a container that implement the action,
		// how to run it application specific as example it can be
		// used to create a pod and expose it as as service ot it
		// can run a sidecar.
		Container *struct {
			// The application container that you want to run,
			corev1.Container `json:"container"`

			// Defines the binding of Action's parameters to the
			// target endpoint.
			Parameters []struct {
				ParameterBinding `json:",inline"`

				// The target environment variable
				Target corev1.EnvVar `json:"target,omitempty"`
			}

			// Metadata ---
			Metadata Metadata `json:"metadata,omitempty"`
		}
	}

	// Dependency defines a set of dependencies that have to be
	// taken into account by the runtime to execute the action.
	//
	// How parameters are bound to the action is implementation
	// specific.
	Dependency *struct {
		Dependencies []Dependency `json:"dependencies,omitempty"`

		// Metadata ---
		Metadata Metadata `json:"metadata,omitempty"`
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
	Deprecated  string `json:"deprecated,omitempty"`
	Schema      Schema `json:"schema,omitempty"`
}

// ParameterBinding ---
type ParameterBinding struct {
	// The name of the parameter
	Name string `json:"name,omitempty"`

	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	SecretKeyRef    *corev1.SecretKeySelector    `json:"secretKeyRef,omitempty"`
}

type EndpointTarget string

const (
	Header EndpointTarget = "header"
	Query  EndpointTarget = "query"
)

type Metadata Metadata

// Message describes
type Message struct {
	// ContentType ---
	ContentType `json:"contentType,omitempty"`

	// Schema ---
	Schema Schema `json:"schema,omitempty"`

	// Metadata ---
	Metadata Metadata `json:"metadata,omitempty"`
}

// Schema describes the shape of message consumed or produced by the
// Action
type Schema struct {
	// Type specifies the type of the data it refers to, this field should be
	// used for primitive types
	Type string `json:"type,omitempty"`

	// Format specifies the format of the data it refers to, for referenced values
	// it can define format fo the schema.
	//
	// The example below, references a json-schema stored in a schema-registry:
	//
	// Schema {
	//     Format:      'json-schema'
	//     RegistryRef: 'http://registry/...'
	// }
	Format string `json:"format,omitempty"`

	// RegistryRef references a schema defined in a schema registry
	// like Apicurio Schema Registry (https://github.com/Apicurio/apicurio-registry)
	//
	// TODO: this could be struct with schema-registry specific bits
	RegistryRef string `json:"registryKeyRef,omitempty"`

	// ConfigMapKeyRef references a key from a ConfigMap where the scheme
	// is defined. The type of the content is defined by the field Type
	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`

	// Metadata ---
	Metadata Metadata `json:"metadata,omitempty"`
}

// ********************************************************************************
//
// Status
//
// ********************************************************************************

// ActionStatus represent the status of an Action at a given time.
type ActionStatus struct {
}
