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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "volcano.sh/apis/pkg/apis/bus/v1alpha1"
	busv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/bus/v1alpha1"
)

// FakeCommands implements CommandInterface
type FakeCommands struct {
	Fake *FakeBusV1alpha1
	ns   string
}

var commandsResource = v1alpha1.SchemeGroupVersion.WithResource("commands")

var commandsKind = v1alpha1.SchemeGroupVersion.WithKind("Command")

// Get takes name of the command, and returns the corresponding command object, and an error if there is any.
func (c *FakeCommands) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Command, err error) {
	emptyResult := &v1alpha1.Command{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(commandsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Command), err
}

// List takes label and field selectors, and returns the list of Commands that match those selectors.
func (c *FakeCommands) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CommandList, err error) {
	emptyResult := &v1alpha1.CommandList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(commandsResource, commandsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CommandList{ListMeta: obj.(*v1alpha1.CommandList).ListMeta}
	for _, item := range obj.(*v1alpha1.CommandList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested commands.
func (c *FakeCommands) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(commandsResource, c.ns, opts))

}

// Create takes the representation of a command and creates it.  Returns the server's representation of the command, and an error, if there is any.
func (c *FakeCommands) Create(ctx context.Context, command *v1alpha1.Command, opts v1.CreateOptions) (result *v1alpha1.Command, err error) {
	emptyResult := &v1alpha1.Command{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(commandsResource, c.ns, command, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Command), err
}

// Update takes the representation of a command and updates it. Returns the server's representation of the command, and an error, if there is any.
func (c *FakeCommands) Update(ctx context.Context, command *v1alpha1.Command, opts v1.UpdateOptions) (result *v1alpha1.Command, err error) {
	emptyResult := &v1alpha1.Command{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(commandsResource, c.ns, command, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Command), err
}

// Delete takes name of the command and deletes it. Returns an error if one occurs.
func (c *FakeCommands) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(commandsResource, c.ns, name, opts), &v1alpha1.Command{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCommands) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(commandsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CommandList{})
	return err
}

// Patch applies the patch and returns the patched command.
func (c *FakeCommands) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Command, err error) {
	emptyResult := &v1alpha1.Command{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(commandsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Command), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied command.
func (c *FakeCommands) Apply(ctx context.Context, command *busv1alpha1.CommandApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Command, err error) {
	if command == nil {
		return nil, fmt.Errorf("command provided to Apply must not be nil")
	}
	data, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}
	name := command.Name
	if name == nil {
		return nil, fmt.Errorf("command.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.Command{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(commandsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.Command), err
}
