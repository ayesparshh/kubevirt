/*
Copyright 2023 The KubeVirt Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubevirt.io/api/instancetype/v1beta1"
)

// FakeVirtualMachineClusterPreferences implements VirtualMachineClusterPreferenceInterface
type FakeVirtualMachineClusterPreferences struct {
	Fake *FakeInstancetypeV1beta1
}

var virtualmachineclusterpreferencesResource = schema.GroupVersionResource{Group: "instancetype.kubevirt.io", Version: "v1beta1", Resource: "virtualmachineclusterpreferences"}

var virtualmachineclusterpreferencesKind = schema.GroupVersionKind{Group: "instancetype.kubevirt.io", Version: "v1beta1", Kind: "VirtualMachineClusterPreference"}

// Get takes name of the virtualMachineClusterPreference, and returns the corresponding virtualMachineClusterPreference object, and an error if there is any.
func (c *FakeVirtualMachineClusterPreferences) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineClusterPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(virtualmachineclusterpreferencesResource, name), &v1beta1.VirtualMachineClusterPreference{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineClusterPreference), err
}

// List takes label and field selectors, and returns the list of VirtualMachineClusterPreferences that match those selectors.
func (c *FakeVirtualMachineClusterPreferences) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineClusterPreferenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(virtualmachineclusterpreferencesResource, virtualmachineclusterpreferencesKind, opts), &v1beta1.VirtualMachineClusterPreferenceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineClusterPreferenceList{ListMeta: obj.(*v1beta1.VirtualMachineClusterPreferenceList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineClusterPreferenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineClusterPreferences.
func (c *FakeVirtualMachineClusterPreferences) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(virtualmachineclusterpreferencesResource, opts))
}

// Create takes the representation of a virtualMachineClusterPreference and creates it.  Returns the server's representation of the virtualMachineClusterPreference, and an error, if there is any.
func (c *FakeVirtualMachineClusterPreferences) Create(ctx context.Context, virtualMachineClusterPreference *v1beta1.VirtualMachineClusterPreference, opts v1.CreateOptions) (result *v1beta1.VirtualMachineClusterPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(virtualmachineclusterpreferencesResource, virtualMachineClusterPreference), &v1beta1.VirtualMachineClusterPreference{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineClusterPreference), err
}

// Update takes the representation of a virtualMachineClusterPreference and updates it. Returns the server's representation of the virtualMachineClusterPreference, and an error, if there is any.
func (c *FakeVirtualMachineClusterPreferences) Update(ctx context.Context, virtualMachineClusterPreference *v1beta1.VirtualMachineClusterPreference, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineClusterPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(virtualmachineclusterpreferencesResource, virtualMachineClusterPreference), &v1beta1.VirtualMachineClusterPreference{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineClusterPreference), err
}

// Delete takes name of the virtualMachineClusterPreference and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineClusterPreferences) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(virtualmachineclusterpreferencesResource, name), &v1beta1.VirtualMachineClusterPreference{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineClusterPreferences) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(virtualmachineclusterpreferencesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineClusterPreferenceList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineClusterPreference.
func (c *FakeVirtualMachineClusterPreferences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineClusterPreference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(virtualmachineclusterpreferencesResource, name, pt, data, subresources...), &v1beta1.VirtualMachineClusterPreference{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineClusterPreference), err
}