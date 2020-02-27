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

package cib

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerImage describe a Container Image
type ContainerImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerImageSpec   `json:"spec,omitempty"`
	Status ContainerImageStatus `json:"status,omitempty"`
}

// ExecutionType ---
type ExecutionType string

// AssemblerType ---
type AssemblerType string

// ContainerImagePhase --
type ContainerImagePhase string

const (
	// ExecutionTypeRoutine ---
	ExecutionTypeRoutine ExecutionType = "routine"

	// ExecutionTypePod ---
	ExecutionTypePod     ExecutionType = "pod"

	// AssemblerTypeBuildah ---
	AssemblerTypeBuildah AssemblerType = "buildah"

	// AssemblerTypeTekton ---
	AssemblerTypeTekton  AssemblerType = "tekton"
)

// ContainerImageCondition ---
type ContainerImageCondition struct {
	// Type of integration condition.
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

// ContainerImageSpec ---
type ContainerImageSpec struct {
	Dependencies []string `json:"dependencies,omitempty"`
	From         string   `json:"from,omitempty"`
	Steps        []string `json:"steps,omitempty"`

	// Registry ---
	Registry struct {
		Insecure     bool   `json:"insecure,omitempty"`
		Address      string `json:"address,omitempty"`
		Secret       string `json:"secret,omitempty"`
		CA           string `json:"ca,omitempty"`
		Organization string `json:"organization,omitempty"`
	}

	Strategy struct {
		// Type define the strategy used to create the container
		// image.
		Type AssemblerType `json:"type,omitempty"`

		// Metadata add additional context to the assembler that
		// can use it as example to build a JVM or native image.
		Metadata map[string]string `json:"metadata,omitempty"`

		// Finalizer defines a container that assemble the resolved
		// dependencies to the structure required by the requestor.
		Finalizer *corev1.Container `json:"finalizer,omitempty"`

		// Execution describe how to execute the assemble phase.
		Execution ExecutionType `json:"execution,omitempty"`
	}
}

// ContainerImageStatus ---
type ContainerImageStatus struct {
	// Image is the final image name
	Image string `json:"image,omitempty"`

	// The phase in which the container image is
	Phase ContainerImagePhase `json:"phase,omitempty"`

	// Conditions detail the current conditions of this container image process.
	Conditions []ContainerImageCondition `json:"conditions,omitempty"`
}
