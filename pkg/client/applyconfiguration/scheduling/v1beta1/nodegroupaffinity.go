/*
Copyright The Volcano Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// NodeGroupAffinityApplyConfiguration represents a declarative configuration of the NodeGroupAffinity type for use
// with apply.
type NodeGroupAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []string `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []string `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NodeGroupAffinityApplyConfiguration constructs a declarative configuration of the NodeGroupAffinity type for use with
// apply.
func NodeGroupAffinity() *NodeGroupAffinityApplyConfiguration {
	return &NodeGroupAffinityApplyConfiguration{}
}

// WithRequiredDuringSchedulingIgnoredDuringExecution adds the given value to the RequiredDuringSchedulingIgnoredDuringExecution field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (b *NodeGroupAffinityApplyConfiguration) WithRequiredDuringSchedulingIgnoredDuringExecution(values ...string) *NodeGroupAffinityApplyConfiguration {
	for i := range values {
		b.RequiredDuringSchedulingIgnoredDuringExecution = append(b.RequiredDuringSchedulingIgnoredDuringExecution, values[i])
	}
	return b
}

// WithPreferredDuringSchedulingIgnoredDuringExecution adds the given value to the PreferredDuringSchedulingIgnoredDuringExecution field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PreferredDuringSchedulingIgnoredDuringExecution field.
func (b *NodeGroupAffinityApplyConfiguration) WithPreferredDuringSchedulingIgnoredDuringExecution(values ...string) *NodeGroupAffinityApplyConfiguration {
	for i := range values {
		b.PreferredDuringSchedulingIgnoredDuringExecution = append(b.PreferredDuringSchedulingIgnoredDuringExecution, values[i])
	}
	return b
}
