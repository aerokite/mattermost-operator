/*
Copyright The Kubernetes Authors.

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
	v1beta1 "github.com/minio/minio-operator/pkg/apis/miniocontroller/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMirrors implements MirrorInterface
type FakeMirrors struct {
	Fake *FakeMinIOV1beta1
	ns   string
}

var mirrorsResource = schema.GroupVersionResource{Group: "min.io.io", Version: "v1beta1", Resource: "mirrors"}

var mirrorsKind = schema.GroupVersionKind{Group: "min.io.io", Version: "v1beta1", Kind: "Mirror"}

// Get takes name of the mirror, and returns the corresponding mirror object, and an error if there is any.
func (c *FakeMirrors) Get(name string, options v1.GetOptions) (result *v1beta1.Mirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mirrorsResource, c.ns, name), &v1beta1.Mirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Mirror), err
}

// List takes label and field selectors, and returns the list of Mirrors that match those selectors.
func (c *FakeMirrors) List(opts v1.ListOptions) (result *v1beta1.MirrorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mirrorsResource, mirrorsKind, c.ns, opts), &v1beta1.MirrorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MirrorList{ListMeta: obj.(*v1beta1.MirrorList).ListMeta}
	for _, item := range obj.(*v1beta1.MirrorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mirrors.
func (c *FakeMirrors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mirrorsResource, c.ns, opts))

}

// Create takes the representation of a mirror and creates it.  Returns the server's representation of the mirror, and an error, if there is any.
func (c *FakeMirrors) Create(mirror *v1beta1.Mirror) (result *v1beta1.Mirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mirrorsResource, c.ns, mirror), &v1beta1.Mirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Mirror), err
}

// Update takes the representation of a mirror and updates it. Returns the server's representation of the mirror, and an error, if there is any.
func (c *FakeMirrors) Update(mirror *v1beta1.Mirror) (result *v1beta1.Mirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mirrorsResource, c.ns, mirror), &v1beta1.Mirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Mirror), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMirrors) UpdateStatus(mirror *v1beta1.Mirror) (*v1beta1.Mirror, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mirrorsResource, "status", c.ns, mirror), &v1beta1.Mirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Mirror), err
}

// Delete takes name of the mirror and deletes it. Returns an error if one occurs.
func (c *FakeMirrors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mirrorsResource, c.ns, name), &v1beta1.Mirror{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMirrors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mirrorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.MirrorList{})
	return err
}

// Patch applies the patch and returns the patched mirror.
func (c *FakeMirrors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Mirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mirrorsResource, c.ns, name, data, subresources...), &v1beta1.Mirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Mirror), err
}
